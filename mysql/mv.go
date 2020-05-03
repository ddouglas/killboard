package mysql

import (
	"context"

	"github.com/eveisesi/neo"
	"github.com/eveisesi/neo/mysql/boiler"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type mvRepository struct {
	db *sqlx.DB
}

func NewMVRepository(db *sqlx.DB) neo.MVRepository {
	return &mvRepository{
		db,
	}
}

func (r *mvRepository) All(ctx context.Context, limit, age int) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)
	err := boiler.Killmails(
		qm.Where("killmail_time >= CURDATE() - INTERVAL ? DAY", age),
		qm.Limit(limit),
		qm.OrderBy(boiler.KillmailColumns.TotalValue+" DESC"),
	).Bind(ctx, r.db, &killmails)

	return killmails, err

}

func (r *mvRepository) KillsByCharacterID(ctx context.Context, id uint64, limit, age int) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	query := `
		SELECT
			killmails.id,
			killmails.hash,
			killmails.moon_id,
			killmails.solar_system_id,
			killmails.war_id,
			killmails.is_npc,
			killmails.is_awox,
			killmails.is_solo,
			killmails.dropped_value,
			killmails.destroyed_value,
			killmails.fitted_value,
			killmails.total_value,
			killmails.killmail_time
		FROM (
			SELECT killmail_attackers.killmail_id FROM killmail_attackers WHERE killmail_attackers.character_id = ?
		) SELECTED_KMS
		LEFT JOIN killmails ON killmails.id = SELECTED_KMS.killmail_id
		WHERE
			killmails.killmail_time >= CURDATE() - INTERVAL ? DAY
		ORDER BY killmails.total_value DESC
		LIMIT ?
	`

	err := r.db.SelectContext(ctx, &killmails, query, id, age, limit)

	return killmails, err
}

func (r *mvRepository) LossesByCharacterID(ctx context.Context, id uint64, limit, age int) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	query := `
		SELECT
			killmails.id,
			killmails.hash,
			killmails.moon_id,
			killmails.solar_system_id,
			killmails.war_id,
			killmails.is_npc,
			killmails.is_awox,
			killmails.is_solo,
			killmails.dropped_value,
			killmails.destroyed_value,
			killmails.fitted_value,
			killmails.total_value,
			killmails.killmail_time
		FROM (
			SELECT killmail_victim.killmail_id FROM killmail_victim WHERE killmail_victim.character_id = ?
		) SELECTED_KMS
		LEFT JOIN killmails ON killmails.id = SELECTED_KMS.killmail_id
		WHERE
			killmails.killmail_time >= CURDATE() - INTERVAL ? DAY
		ORDER BY killmails.total_value DESC
		LIMIT ?
	`

	err := r.db.SelectContext(ctx, &killmails, query, id, age, limit)

	return killmails, err
}

func (r *mvRepository) KillsByCorporationID(ctx context.Context, id uint64, limit, age int) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	query := `
		SELECT
			killmails.id,
			killmails.hash,
			killmails.moon_id,
			killmails.solar_system_id,
			killmails.war_id,
			killmails.is_npc,
			killmails.is_awox,
			killmails.is_solo,
			killmails.dropped_value,
			killmails.destroyed_value,
			killmails.fitted_value,
			killmails.total_value,
			killmails.killmail_time
		FROM (
			SELECT killmail_attackers.killmail_id FROM killmail_attackers WHERE killmail_attackers.corporation_id = ?
		) SELECTED_KMS
		LEFT JOIN killmails ON killmails.id = SELECTED_KMS.killmail_id
		WHERE
			killmails.killmail_time >= CURDATE() - INTERVAL ? DAY
		ORDER BY killmails.total_value DESC
		LIMIT ?
	`

	err := r.db.SelectContext(ctx, &killmails, query, id, age, limit)

	return killmails, err
}

func (r *mvRepository) LossesByCorporationID(ctx context.Context, id uint64, limit, age int) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	query := `
		SELECT
			killmails.id,
			killmails.hash,
			killmails.moon_id,
			killmails.solar_system_id,
			killmails.war_id,
			killmails.is_npc,
			killmails.is_awox,
			killmails.is_solo,
			killmails.dropped_value,
			killmails.destroyed_value,
			killmails.fitted_value,
			killmails.total_value,
			killmails.killmail_time
		FROM (
			SELECT killmail_victim.killmail_id FROM killmail_victim WHERE killmail_victim.corporation_id = ?
		) SELECTED_KMS
		LEFT JOIN killmails ON killmails.id = SELECTED_KMS.killmail_id
		WHERE
			killmails.killmail_time >= CURDATE() - INTERVAL ? DAY
		ORDER BY killmails.total_value DESC
		LIMIT ?
	`

	err := r.db.SelectContext(ctx, &killmails, query, id, age, limit)

	return killmails, err
}

func (r *mvRepository) KillsByAllianceID(ctx context.Context, id uint64, limit, age int) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	query := `
		SELECT
			killmails.id,
			killmails.hash,
			killmails.moon_id,
			killmails.solar_system_id,
			killmails.war_id,
			killmails.is_npc,
			killmails.is_awox,
			killmails.is_solo,
			killmails.dropped_value,
			killmails.destroyed_value,
			killmails.fitted_value,
			killmails.total_value,
			killmails.killmail_time
		FROM (
			SELECT killmail_attackers.killmail_id FROM killmail_attackers WHERE killmail_attackers.alliance_id = ?
		) SELECTED_KMS
		LEFT JOIN killmails ON killmails.id = SELECTED_KMS.killmail_id
		WHERE
			killmails.killmail_time >= CURDATE() - INTERVAL ? DAY
		ORDER BY killmails.total_value DESC
		LIMIT ?
	`

	err := r.db.SelectContext(ctx, &killmails, query, id, age, limit)

	return killmails, err
}

func (r *mvRepository) LossesByAllianceID(ctx context.Context, id uint64, limit, age int) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	query := `
		SELECT
			killmails.id,
			killmails.hash,
			killmails.moon_id,
			killmails.solar_system_id,
			killmails.war_id,
			killmails.is_npc,
			killmails.is_awox,
			killmails.is_solo,
			killmails.dropped_value,
			killmails.destroyed_value,
			killmails.fitted_value,
			killmails.total_value,
			killmails.killmail_time
		FROM (
			SELECT killmail_victim.killmail_id FROM killmail_victim WHERE killmail_victim.alliance_id = ?
		) SELECTED_KMS
		LEFT JOIN killmails ON killmails.id = SELECTED_KMS.killmail_id
		WHERE
			killmails.killmail_time >= CURDATE() - INTERVAL ? DAY
		ORDER BY killmails.total_value DESC
		LIMIT ?
	`

	err := r.db.SelectContext(ctx, &killmails, query, id, age, limit)

	return killmails, err
}

func (r *mvRepository) KillsByShipID(ctx context.Context, id uint64, limit, age int) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	query := `
		SELECT
			killmails.id,
			killmails.hash,
			killmails.moon_id,
			killmails.solar_system_id,
			killmails.war_id,
			killmails.is_npc,
			killmails.is_awox,
			killmails.is_solo,
			killmails.dropped_value,
			killmails.destroyed_value,
			killmails.fitted_value,
			killmails.total_value,
			killmails.killmail_time
		FROM (
			SELECT killmail_attackers.killmail_id FROM killmail_attackers WHERE killmail_attackers.ship_type_id = ?
		) SELECTED_KMS
		LEFT JOIN killmails ON killmails.id = SELECTED_KMS.killmail_id
		WHERE
			killmails.killmail_time >= CURDATE() - INTERVAL ? DAY
		ORDER BY killmails.total_value DESC
		LIMIT ?
	`

	err := r.db.SelectContext(ctx, &killmails, query, id, age, limit)

	return killmails, err
}

func (r *mvRepository) LossesByShipID(ctx context.Context, id uint64, limit, age int) ([]*neo.Killmail, error) {

	var killmails = make([]*neo.Killmail, 0)

	query := `
		SELECT
			killmails.id,
			killmails.hash,
			killmails.moon_id,
			killmails.solar_system_id,
			killmails.war_id,
			killmails.is_npc,
			killmails.is_awox,
			killmails.is_solo,
			killmails.dropped_value,
			killmails.destroyed_value,
			killmails.fitted_value,
			killmails.total_value,
			killmails.killmail_time
		FROM (
			SELECT killmail_victim.killmail_id FROM killmail_victim WHERE killmail_victim.ship_type_id = ?
		) SELECTED_KMS
		LEFT JOIN killmails ON killmails.id = SELECTED_KMS.killmail_id
		WHERE
			killmails.killmail_time >= CURDATE() - INTERVAL ? DAY
		ORDER BY killmails.total_value DESC
		LIMIT ?
	`

	err := r.db.SelectContext(ctx, &killmails, query, id, age, limit)

	return killmails, err
}
