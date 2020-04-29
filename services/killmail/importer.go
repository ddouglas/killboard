package killmail

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/eveisesi/neo"
	"github.com/go-redis/redis"
	"github.com/korovkin/limiter"
	"github.com/sirupsen/logrus"
)

// now := time.Now()
// location, _ := time.LoadLocation("UTC")

// // then := time.Until
// return nil

func (s *service) loopManager() {

	for {

		status, err := s.redis.Get(neo.REDIS_ESI_TRACKING_STATUS).Int64()
		if err != nil && err.Error() != neo.ErrRedisNil.Error() {
			break
		}

		if status == neo.COUNT_STATUS_DOWNTIME ||
			status == neo.COUNT_STATUS_RED {

			s.logger.WithField("status", status).Info("loop manager blocking process")
			time.Sleep(time.Second)
			continue
		} else if status == neo.COUNT_STATUS_YELLOW {
			s.logger.WithField("status", status).Warning("slowing down due to status")
			time.Sleep(time.Millisecond * 250)
			break
		} else if status == neo.COUNT_STATUS_GREEN {
			break
		}

	}

}

func (s *service) Importer(channel string, gLimit, gSleep int64) error {

	limit := limiter.NewConcurrencyLimiter(int(gLimit))

	for {
		count, err := s.redis.ZCount(channel, "-inf", "+inf").Result()
		if err != nil {
			s.logger.WithError(err).Fatal("unable to determine count of message queue")
		}

		if count == 0 {
			s.logger.Info("message queue is empty")
			time.Sleep(time.Second * 2)
			continue
		}

		results, err := s.redis.ZPopMin(channel, gLimit).Result()
		if err != nil {
			s.logger.WithError(err).Fatal("unable to retrieve hashes from queue")
		}

		for _, result := range results {
			s.loopManager()
			message := result.Member.(string)
			limit.ExecuteWithTicket(func(workerID int) {
				s.processMessage([]byte(message), workerID, gSleep)
			})
		}
	}
}

func (s *service) processMessage(message []byte, workerID int, sleep int64) {

	var ctx = context.Background()

	var payload Message
	err := json.Unmarshal(message, &payload)
	if err != nil {
		s.logger.WithField("message", string(message)).Fatal("failed to unmarhal message into message struct")
	}

	killmailLoggerFields := logrus.Fields{
		"id":     payload.ID,
		"hash":   payload.Hash,
		"worker": workerID,
	}

	killmailID, err := strconv.ParseUint(payload.ID, 10, 64)
	if err != nil {
		s.logger.WithFields(killmailLoggerFields).Error("unable to parse killmail id to uint")
		return
	}

	exists, err := s.KillmailExists(ctx, killmailID, payload.Hash)
	if err != nil {
		s.logger.WithError(err).
			WithFields(killmailLoggerFields).Error("error encountered checking if killmail exists")
	}

	if exists {
		s.logger.WithFields(killmailLoggerFields).Info("skipping existing killmail")
		return
	}

	killmail, m := s.esi.GetKillmailsKillmailIDKillmailHash(payload.ID, payload.Hash)
	if m.IsError() {
		s.logger.WithError(m.Msg).WithFields(logrus.Fields{
			"code":  m.Code,
			"path":  m.Path,
			"query": m.Query,
		}).Error("failed to fetch killmail from esi")
		s.redis.ZAdd(channel, redis.Z{Score: 0, Member: message})
		return
	}

	if m.Code != 200 {
		s.logger.WithFields(killmailLoggerFields).WithFields(logrus.Fields{
			"code":  m.Code,
			"path":  m.Path,
			"query": m.Query,
		}).WithError(err).Error("unexpected response code from esi")
		s.redis.ZAdd(channel, redis.Z{Score: 0, Member: message})
		return
	}

	killmail.Hash = payload.Hash

	_, err = s.universe.SolarSystem(ctx, killmail.SolarSystemID)
	if err != nil {
		s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to prime solar system")
	}

	victim := killmail.Victim

	if victim.AllianceID.Valid {
		_, err := s.alliance.Alliance(ctx, victim.AllianceID.Uint64)
		if err != nil {
			s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to prime victim alliance")
		}
	}

	if victim.CorporationID.Valid {
		_, err := s.corporation.Corporation(ctx, victim.CorporationID.Uint64)
		if err != nil {
			s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to prime victim character")
		}
	}

	if victim.CharacterID.Valid {
		_, err := s.character.Character(ctx, victim.CharacterID.Uint64)
		if err != nil {
			s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to prime victim character")
		}
	}

	_, err = s.universe.Type(ctx, victim.ShipTypeID)
	if err != nil {
		s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to prime victim ship type")
	}

	for _, attacker := range killmail.Attackers {
		if attacker.AllianceID.Valid {
			_, err := s.alliance.Alliance(ctx, attacker.AllianceID.Uint64)
			if err != nil {
				s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to prime attacker alliance")
			}
		}

		if attacker.CorporationID.Valid {
			_, err := s.corporation.Corporation(ctx, attacker.CorporationID.Uint64)
			if err != nil {
				s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to prime attacker character")
			}
		}

		if attacker.CharacterID.Valid {
			_, err := s.character.Character(ctx, attacker.CharacterID.Uint64)
			if err != nil {
				s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to prime attacker character")
			}
		}

		if attacker.ShipTypeID.Valid {
			_, err = s.universe.Type(ctx, attacker.ShipTypeID.Uint64)
			if err != nil {
				s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to prime attacker ship type")
			}
		}

		if attacker.WeaponTypeID.Valid {
			_, err = s.universe.Type(ctx, attacker.WeaponTypeID.Uint64)
			if err != nil {
				s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to prime attacker ship type")
			}
		}
	}

	s.handleVictimItems(killmail.Victim.Items)

	txn, err := s.txn.Begin()
	if err != nil {
		s.logger.WithError(err).Error("failed to start transaction")
		return
	}

	_, err = s.KillmailRespository.CreateKillmailTxn(ctx, txn, killmail)
	if err != nil {
		s.logger.WithFields(killmailLoggerFields).WithError(err).Error("error encountered inserting killmail into db")
		return
	}

	killmail.Victim.KillmailID = killmail.ID

	if killmail.Victim.Position != nil {
		killmail.Victim.PosX.SetValid(killmail.Victim.Position.X.Float64)
		killmail.Victim.PosY.SetValid(killmail.Victim.Position.Y.Float64)
		killmail.Victim.PosZ.SetValid(killmail.Victim.Position.Z.Float64)
	}

	date := killmail.KillmailTime

	var totalValue = make([]float64, 0)
	shipValue := s.market.FetchTypePrice(killmail.Victim.ShipTypeID, date)
	killmail.Victim.ShipValue = shipValue

	_, err = s.KillmailRespository.CreateKillmailVictimTxn(ctx, txn, killmail.Victim)
	if err != nil {
		s.logger.WithFields(killmailLoggerFields).WithError(err).Error("error encountered inserting killmail victim into db")
	}

	for _, attacker := range killmail.Attackers {
		attacker.KillmailID = killmailID
	}

	_, err = s.KillmailRespository.CreateKillmailAttackersTxn(ctx, txn, killmail.Attackers)
	if err != nil {
		s.logger.WithFields(killmailLoggerFields).WithError(err).Error("error encountered inserting killmail attackers into db")
	}

	for _, item := range killmail.Victim.Items {
		item.KillmailID = killmail.ID
		itemValue := float64(0)
		if item.Singleton != 2 {
			itemValue = s.market.FetchTypePrice(item.ItemTypeID, date)
		} else {
			itemValue = 0.01
		}

		quantity := item.QuantityDestroyed.Uint64 + item.QuantityDropped.Uint64
		totalValue = append(totalValue, itemValue*float64(quantity))

		item.ItemValue = itemValue
	}

	_, err = s.KillmailRespository.CreateKillmailItemsTxn(ctx, txn, killmail.Victim.Items)
	if err != nil {
		s.logger.WithError(err).Error("failed to insert items into db")
	}

	destroyedValue := float64(0)
	droppedValue := float64(0)

	for _, item := range killmail.Victim.Items {
		if len(item.Items) > 0 {
			for _, subItem := range item.Items {
				subItem.KillmailID = killmailID
				subItem.ParentID.SetValid(item.ID)
				subItemValue := float64(0)
				if item.Singleton != 2 {
					subItemValue = 0.01
				} else {
					subItemValue = s.market.FetchTypePrice(item.ItemTypeID, date)
				}
				itemTotal := subItemValue * float64(subItem.QuantityDestroyed.Uint64+subItem.QuantityDropped.Uint64)
				totalValue = append(totalValue, itemTotal)
				subItem.ItemValue = subItemValue

				if subItem.QuantityDestroyed.Valid {
					destroyedValue += itemTotal
				} else if subItem.QuantityDropped.Valid {
					droppedValue += itemTotal
				}
			}

			_, err = s.KillmailRespository.CreateKillmailItemsTxn(ctx, txn, item.Items)
			if err != nil {
				s.logger.WithFields(killmailLoggerFields).WithError(err).Error("error encountered insert sub items")
			}
		}
		if item.QuantityDestroyed.Valid {
			destroyedValue += item.ItemValue * item.QuantityDestroyed.Uint64
		} else if item.QuantityDropped.Valid {
			droppedValue += item.ItemValue * item.QuantityDropped.Uint64
		}
	}

	fittedValue := s.calculatedFittedValue(killmail.Victim.Items)
	shipValue += fittedValue
	totalValue = append(totalValue, shipValue)

	sum := float64(0)
	for _, v := range totalValue {
		sum += v
	}

	err = txn.Commit()
	if err != nil {
		s.logger.WithFields(killmailLoggerFields).WithError(err).Error("failed to commit transaction")
		return
	}

	killmail.DestroyedValue = destroyedValue
	killmail.DroppedValue = droppedValue
	killmail.FittedValue = fittedValue
	killmail.TotalValue = sum

	err = s.KillmailRespository.UpdateKillmail(ctx, killmail.ID, killmail.Hash, killmail)
	if err != nil {
		s.logger.WithFields(killmailLoggerFields).WithError(err).Error("error encountered inserting killmail victim into db")
	}

	s.logger.WithFields(killmailLoggerFields).Info("killmail successfully imported")
	time.Sleep(time.Millisecond * time.Duration(sleep))
}

func (s *service) handleVictimItems(items []*neo.KillmailItem) {
	for _, item := range items {
		_, err := s.universe.Type(context.Background(), item.ItemTypeID)
		if err != nil {
			s.logger.WithError(err).WithFields(logrus.Fields{
				"item_type_id": item.ItemTypeID,
			}).Error("encountered error")
		}
		if len(item.Items) > 0 {
			s.handleVictimItems(item.Items)
		}
	}
}

var fittedFlags = map[uint64]bool{11: true, 12: true, 13: true, 87: true, 89: true, 93: true, 158: true, 159: true, 172: true, 2663: true, 3772: true}

func (s *service) calculatedFittedValue(items []*neo.KillmailItem) float64 {

	total := float64(0)
	for _, item := range items {
		if _, ok := fittedFlags[uint64(item.Flag)]; !ok {
			continue
		}
		quantity := item.QuantityDestroyed.Uint64 + item.QuantityDropped.Uint64
		total += item.ItemValue * float64(quantity)
	}

	return total
}