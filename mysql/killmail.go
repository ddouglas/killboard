package mysql

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/eveisesi/neo/mysql/boiler"

	"github.com/eveisesi/neo"
	"github.com/jmoiron/sqlx"
)

type killmailRepository struct {
	db *sqlx.DB
}

func NewKillmailRepository(db *sqlx.DB) neo.KillmailRespository {
	return &killmailRepository{
		db,
	}
}

func (r *killmailRepository) Killmail(ctx context.Context, id uint64, hash string) (*neo.Killmail, error) {

	var killmail = neo.Killmail{}
	err := boiler.Killmails(
		boiler.KillmailWhere.ID.EQ(id),
		boiler.KillmailWhere.Hash.EQ(hash),
	).Bind(ctx, r.db, &killmail)

	return &killmail, err

}

func (r *killmailRepository) KillmailGTID(ctx context.Context, id null.Uint64) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)
	mods := make([]qm.QueryMod, 0)
	if id.Valid {
		mods = append(mods, boiler.KillmailWhere.ID.GT(id.Uint64))
	}

	mods = append(mods, qm.Limit(10000), qm.OrderBy(boiler.KillmailColumns.ID+" ASC"))

	err := boiler.Killmails(
		mods...,
	).Bind(ctx, r.db, &killmails)

	return killmails, err

}

func (r *killmailRepository) CreateKillmail(ctx context.Context, killmail *neo.Killmail) (*neo.Killmail, error) {

	var bKillmail = new(boiler.Killmail)
	err := copier.Copy(bKillmail, killmail)
	if err != nil {
		return nil, errors.Wrap(err, "unable to copy killmail to orm")
	}

	err = bKillmail.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "unable to insert killmail into db")
	}

	err = copier.Copy(killmail, bKillmail)

	return killmail, errors.Wrap(err, "unable to copy orm to killmail")

}

func (r *killmailRepository) CreateKillmailTxn(ctx context.Context, txn neo.Transactioner, killmail *neo.Killmail) (*neo.Killmail, error) {
	var t = txn.(*transaction)
	var bKillmail = new(boiler.Killmail)
	err := copier.Copy(bKillmail, killmail)
	if err != nil {
		return nil, errors.Wrap(err, "unable to copy killmail to orm")
	}

	err = bKillmail.Insert(ctx, t, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "unable to insert killmail into db")
	}

	err = copier.Copy(killmail, bKillmail)

	return killmail, errors.Wrap(err, "unable to copy orm to killmail")

}

func (r *killmailRepository) UpdateKillmail(ctx context.Context, id uint64, hash string, killmail *neo.Killmail) error {

	var bKillmail = new(boiler.Killmail)
	err := copier.Copy(bKillmail, killmail)
	if err != nil {
		return errors.Wrap(err, "unable to copy killmail to orm")
	}

	bKillmail.ID = id
	bKillmail.Hash = hash

	_, err = bKillmail.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return errors.Wrap(err, "failed to update killmail in db")
	}

	err = copier.Copy(killmail, bKillmail)

	return errors.Wrap(err, "unable to copy orm to killmail")

}

func (r *killmailRepository) KillmailExists(ctx context.Context, id uint64, hash string) (bool, error) {
	return boiler.KillmailExists(ctx, r.db, id, hash)
}

func (r *killmailRepository) KillmailTop(ctx context.Context, age uint64, limit uint64) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)
	err := boiler.Killmails(
		qm.Where("killmail_time >= CURDATE() - INTERVAL ? DAY", age),
		qm.Limit(int(limit)),
		qm.OrderBy(boiler.KillmailColumns.TotalValue+" DESC"),
	).Bind(ctx, r.db, &killmails)

	return killmails, err

}

func (r *killmailRepository) KillmailRecent(ctx context.Context, page null.Int) ([]*neo.Killmail, error) {

	limit := 50
	offset := 0

	if page.Valid && page.Int >= 0 {
		limit = page.Int * limit
		offset = limit - 50
	}

	var killmails = make([]*neo.Killmail, 0)
	err := boiler.Killmails(
		qm.Limit(limit),
		qm.Offset(offset),
		qm.OrderBy(
			fmt.Sprintf(
				"%s DESC",
				boiler.KillmailColumns.ID,
			),
		),
	).Bind(ctx, r.db, &killmails)

	return killmails, err

}

func (r *killmailRepository) KillmailAttackersByKillmailIDs(ctx context.Context, ids []uint64) ([]*neo.KillmailAttacker, error) {

	var attackers = make([]*neo.KillmailAttacker, 0)
	err := boiler.KillmailAttackers(
		qm.WhereIn(
			fmt.Sprintf(
				"%s IN ?",
				boiler.KillmailAttackerColumns.KillmailID,
			),
			convertSliceUint64ToSliceInterface(ids)...,
		),
	).Bind(ctx, r.db, &attackers)

	return attackers, err

}

func (r *killmailRepository) CreateKillmailAttacker(ctx context.Context, attacker *neo.KillmailAttacker) (*neo.KillmailAttacker, error) {

	var bAttacker = new(boiler.KillmailAttacker)
	err := copier.Copy(bAttacker, attacker)
	if err != nil {
		return nil, errors.Wrap(err, "failed to copy attacker to orm")
	}

	err = bAttacker.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert attacker into db")
	}

	err = copier.Copy(attacker, bAttacker)

	return attacker, errors.Wrap(err, "failed to copy orm to attacker")

}

func (r *killmailRepository) CreateKillmailAttackerTxn(ctx context.Context, txn neo.Transactioner, attacker *neo.KillmailAttacker) (*neo.KillmailAttacker, error) {

	var t = txn.(*transaction)
	var bAttacker = new(boiler.KillmailAttacker)
	err := copier.Copy(bAttacker, attacker)
	if err != nil {
		return nil, errors.Wrap(err, "failed to copy attacker to orm")
	}

	err = bAttacker.Insert(ctx, t, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert attacker into db")
	}

	err = copier.Copy(attacker, bAttacker)

	return attacker, errors.Wrap(err, "failed to copy orm to attacker")

}

func (r *killmailRepository) CreateKillmailAttackers(ctx context.Context, attackers []*neo.KillmailAttacker) ([]*neo.KillmailAttacker, error) {

	for _, attacker := range attackers {
		var bAttacker = new(boiler.KillmailAttacker)
		err := copier.Copy(bAttacker, attacker)
		if err != nil {
			return nil, errors.Wrap(err, "failed to copy attacker to orm")
		}

		err = bAttacker.Insert(ctx, r.db, boil.Infer())
		if err != nil {
			return nil, errors.Wrap(err, "failed to insert attacker into db")
		}

		err = copier.Copy(attacker, bAttacker)
		if err != nil {
			return nil, errors.Wrap(err, "failed to copy orm to attacker")
		}

	}

	return attackers, nil

}

func (r *killmailRepository) CreateKillmailAttackersTxn(ctx context.Context, txn neo.Transactioner, attackers []*neo.KillmailAttacker) ([]*neo.KillmailAttacker, error) {

	var t = txn.(*transaction)
	for _, attacker := range attackers {
		var bAttacker = new(boiler.KillmailAttacker)
		err := copier.Copy(bAttacker, attacker)
		if err != nil {
			return nil, errors.Wrap(err, "failed to copy attacker to orm")
		}

		err = bAttacker.Insert(ctx, t, boil.Infer())
		if err != nil {
			return nil, errors.Wrap(err, "failed to insert attacker into db")
		}

		err = copier.Copy(attacker, bAttacker)
		if err != nil {
			return nil, errors.Wrap(err, "failed to copy orm to attacker")
		}

	}

	return attackers, nil

}

func (r *killmailRepository) KillmailItemsByKillmailID(ctx context.Context, id uint64) ([]*neo.KillmailItem, error) {

	var items = make([]*neo.KillmailItem, 0)
	err := boiler.KillmailItems(
		boiler.KillmailItemWhere.KillmailID.EQ(id),
	).Bind(ctx, r.db, &items)

	return items, err

}

func (r *killmailRepository) KillmailItemsByKillmailIDs(ctx context.Context, ids []uint64) ([]*neo.KillmailItem, error) {

	var items = make([]*neo.KillmailItem, 0)
	err := boiler.KillmailItems(
		qm.WhereIn(
			fmt.Sprintf(
				"%s IN ?",
				boiler.KillmailItemColumns.KillmailID,
			),
			convertSliceUint64ToSliceInterface(ids)...,
		),
	).Bind(ctx, r.db, &items)

	return items, err

}

func (r *killmailRepository) CreateKillmailItem(ctx context.Context, item *neo.KillmailItem) (*neo.KillmailItem, error) {

	var bItem = new(boiler.KillmailItem)
	err := copier.Copy(bItem, item)
	if err != nil {
		return nil, errors.Wrap(err, "failed to copy item to orm")
	}

	err = bItem.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert item into db")
	}

	err = copier.Copy(item, bItem)

	return item, errors.Wrap(err, "failed to copy orm to item")

}

func (r *killmailRepository) CreateKillmailItemTxn(ctx context.Context, txn neo.Transactioner, item *neo.KillmailItem) (*neo.KillmailItem, error) {

	var t = txn.(*transaction)
	var bItem = new(boiler.KillmailItem)
	err := copier.Copy(bItem, item)
	if err != nil {
		return nil, errors.Wrap(err, "failed to copy item to orm")
	}

	err = bItem.Insert(ctx, t, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert item into db")
	}

	err = copier.Copy(item, bItem)

	return item, errors.Wrap(err, "failed to copy orm to item")

}

func (r *killmailRepository) CreateKillmailItems(ctx context.Context, items []*neo.KillmailItem) ([]*neo.KillmailItem, error) {

	for _, item := range items {
		var bItem = new(boiler.KillmailItem)
		err := copier.Copy(bItem, item)
		if err != nil {
			return nil, errors.Wrap(err, "failed to copy item to orm")
		}

		err = bItem.Insert(ctx, r.db, boil.Infer())
		if err != nil {
			return nil, errors.Wrap(err, "failed to insert item into db")
		}

		err = copier.Copy(item, bItem)
		if err != nil {
			return nil, errors.Wrap(err, "failed to copy orm to item")
		}

	}

	return items, nil

}

func (r *killmailRepository) CreateKillmailItemsTxn(ctx context.Context, txn neo.Transactioner, items []*neo.KillmailItem) ([]*neo.KillmailItem, error) {

	var t = txn.(*transaction)
	for _, item := range items {
		var bItem = new(boiler.KillmailItem)
		err := copier.Copy(bItem, item)
		if err != nil {
			return nil, errors.Wrap(err, "failed to copy item to orm")
		}

		err = bItem.Insert(ctx, t, boil.Infer())
		if err != nil {
			return nil, errors.Wrap(err, "failed to insert item into db")
		}

		err = copier.Copy(item, bItem)
		if err != nil {
			return nil, errors.Wrap(err, "failed to copy orm to item")
		}

	}

	return items, nil

}

func (r *killmailRepository) UpdateKillmailItems(ctx context.Context, items []*neo.KillmailItem) error {

	for _, item := range items {
		var bItem = new(boiler.KillmailItem)
		err := copier.Copy(bItem, item)
		if err != nil {
			return errors.Wrap(err, "failed to copy item to orm")
		}

		_, err = bItem.Update(ctx, r.db, boil.Infer())
		if err != nil {
			return errors.Wrap(err, "failed to update item in db")
		}

		err = copier.Copy(item, bItem)
		if err != nil {
			return errors.Wrap(err, "failed to copy orm to item")
		}

	}

	return nil

}

func (r *killmailRepository) UpdateKillmailItemsTxn(ctx context.Context, txn neo.Transactioner, items []*neo.KillmailItem) error {

	var t = txn.(*transaction)
	for _, item := range items {
		var bItem = new(boiler.KillmailItem)
		err := copier.Copy(bItem, item)
		if err != nil {
			return errors.Wrap(err, "failed to copy item to orm")
		}

		_, err = bItem.Update(ctx, t, boil.Infer())
		if err != nil {
			return errors.Wrap(err, "failed to update item in db")
		}

		err = copier.Copy(item, bItem)
		if err != nil {
			return errors.Wrap(err, "failed to copy orm to item")
		}

	}

	return nil

}

func (r *killmailRepository) KillmailAttackersByKillmailID(ctx context.Context, id uint64) ([]*neo.KillmailAttacker, error) {

	var attackers = make([]*neo.KillmailAttacker, 0)
	err := boiler.KillmailAttackers(
		boiler.KillmailAttackerWhere.KillmailID.EQ(id),
	).Bind(ctx, r.db, &attackers)

	return attackers, err

}

func (r *killmailRepository) KillmailVictimByKillmailID(ctx context.Context, id uint64) (*neo.KillmailVictim, error) {

	var victim = new(neo.KillmailVictim)
	err := boiler.KillmailVictims(
		boiler.KillmailVictimWhere.KillmailID.EQ(id),
	).Bind(ctx, r.db, victim)

	return victim, err

}

func (r *killmailRepository) KillmailVictimsByKillmailIDs(ctx context.Context, ids []uint64) ([]*neo.KillmailVictim, error) {

	var victims = make([]*neo.KillmailVictim, 0)
	err := boiler.KillmailVictims(
		qm.WhereIn(
			fmt.Sprintf(
				"%s IN ?",
				boiler.KillmailVictimColumns.KillmailID,
			),
			convertSliceUint64ToSliceInterface(ids)...,
		),
	).Bind(ctx, r.db, &victims)

	return victims, err

}

func (r *killmailRepository) CreateKillmailVictim(ctx context.Context, victim *neo.KillmailVictim) (*neo.KillmailVictim, error) {

	var bVictim = new(boiler.KillmailVictim)
	err := copier.Copy(bVictim, victim)
	if err != nil {
		return nil, errors.Wrap(err, "failed to copy victim to orm")
	}

	err = bVictim.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert victim into db")
	}

	err = copier.Copy(victim, bVictim)

	return victim, errors.Wrap(err, "failed to copy orm to victim")

}

func (r *killmailRepository) CreateKillmailVictimTxn(ctx context.Context, txn neo.Transactioner, victim *neo.KillmailVictim) (*neo.KillmailVictim, error) {

	var t = txn.(*transaction)
	var bVictim = new(boiler.KillmailVictim)
	err := copier.Copy(bVictim, victim)
	if err != nil {
		return nil, errors.Wrap(err, "failed to copy victim to orm")
	}

	err = bVictim.Insert(ctx, t, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert victim into db")
	}

	err = copier.Copy(victim, bVictim)

	return victim, errors.Wrap(err, "failed to copy orm to victim")

}

func (r *killmailRepository) UpdateKillmailVictim(ctx context.Context, victim *neo.KillmailVictim) error {
	var bVictim = new(boiler.KillmailVictim)
	err := copier.Copy(bVictim, victim)
	if err != nil {
		return errors.Wrap(err, "failed to copy victim to orm")
	}

	_, err = bVictim.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return errors.Wrap(err, "failed to update victim in db")
	}

	err = copier.Copy(victim, bVictim)

	return errors.Wrap(err, "failed to copy orm to victim")
}

func (r *killmailRepository) UpdateKillmailVictimTxn(ctx context.Context, txn neo.Transactioner, victim *neo.KillmailVictim) error {

	var t = txn.(*transaction)
	var bVictim = new(boiler.KillmailVictim)
	err := copier.Copy(bVictim, victim)
	if err != nil {
		return errors.Wrap(err, "failed to copy victim to orm")
	}

	_, err = bVictim.Update(ctx, t, boil.Infer())
	if err != nil {
		return errors.Wrap(err, "failed to update victim in db")
	}

	err = copier.Copy(victim, bVictim)

	return errors.Wrap(err, "failed to copy orm to victim")

}

func (r *killmailRepository) KillmailsByCharacterID(ctx context.Context, id uint64) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	err := boiler.Killmails(
		qm.InnerJoin(
			fmt.Sprintf(
				"%s ON %s.%s = %s.%s",
				boiler.TableNames.KillmailVictim,
				boiler.TableNames.KillmailVictim,
				boiler.KillmailVictimColumns.KillmailID,
				boiler.TableNames.Killmails,
				boiler.KillmailColumns.ID,
			),
		),
		qm.InnerJoin(
			fmt.Sprintf(
				"%s ON %s.%s = %s.%s",
				boiler.TableNames.KillmailAttackers,
				boiler.TableNames.KillmailAttackers,
				boiler.KillmailAttackerColumns.KillmailID,
				boiler.TableNames.Killmails,
				boiler.KillmailColumns.ID,
			),
		),
		qm.Where(
			fmt.Sprintf(
				"%s.%s = ?",
				boiler.TableNames.KillmailVictim,
				boiler.KillmailVictimColumns.CharacterID,
			),
			id,
		),
		qm.Where(
			fmt.Sprintf(
				"%s.%s = ?",
				boiler.TableNames.KillmailAttackers,
				boiler.KillmailAttackerColumns.CharacterID,
			),
			id,
		),
	).Bind(ctx, r.db, &killmails)

	return killmails, err

}

func (r *killmailRepository) KillmailsByCorporationID(ctx context.Context, id uint64) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	err := boiler.Killmails(
		qm.InnerJoin(
			fmt.Sprintf(
				"%s ON %s.%s = %s.%s",
				boiler.TableNames.KillmailVictim,
				boiler.TableNames.KillmailVictim,
				boiler.KillmailVictimColumns.KillmailID,
				boiler.TableNames.Killmails,
				boiler.KillmailColumns.ID,
			),
		),
		qm.InnerJoin(
			fmt.Sprintf(
				"%s ON %s.%s = %s.%s",
				boiler.TableNames.KillmailAttackers,
				boiler.TableNames.KillmailAttackers,
				boiler.KillmailAttackerColumns.KillmailID,
				boiler.TableNames.Killmails,
				boiler.KillmailColumns.ID,
			),
		),
		qm.Where(
			fmt.Sprintf(
				"%s.%s = ?",
				boiler.TableNames.KillmailVictim,
				boiler.KillmailVictimColumns.CharacterID,
			),
			id,
		),
		qm.Where(
			fmt.Sprintf(
				"%s.%s = ?",
				boiler.TableNames.KillmailAttackers,
				boiler.KillmailAttackerColumns.CharacterID,
			),
			id,
		),
	).Bind(ctx, r.db, &killmails)

	return killmails, err

}

func (r *killmailRepository) KillmailsByAllianceID(ctx context.Context, id uint64) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	err := boiler.Killmails(
		qm.InnerJoin(
			fmt.Sprintf(
				"%s ON %s.%s = %s.%s",
				boiler.TableNames.KillmailVictim,
				boiler.TableNames.KillmailVictim,
				boiler.KillmailVictimColumns.KillmailID,
				boiler.TableNames.Killmails,
				boiler.KillmailColumns.ID,
			),
		),
		qm.InnerJoin(
			fmt.Sprintf(
				"%s ON %s.%s = %s.%s",
				boiler.TableNames.KillmailAttackers,
				boiler.TableNames.KillmailAttackers,
				boiler.KillmailAttackerColumns.KillmailID,
				boiler.TableNames.Killmails,
				boiler.KillmailColumns.ID,
			),
		),
		qm.Where(
			fmt.Sprintf(
				"%s.%s = ?",
				boiler.TableNames.KillmailVictim,
				boiler.KillmailVictimColumns.CharacterID,
			),
			id,
		),
		qm.Where(
			fmt.Sprintf(
				"%s.%s = ?",
				boiler.TableNames.KillmailAttackers,
				boiler.KillmailAttackerColumns.CharacterID,
			),
			id,
		),
	).Bind(ctx, r.db, &killmails)

	return killmails, err

}
