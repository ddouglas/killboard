package mysql

import (
	"context"
	"fmt"

	"github.com/eveisesi/neo"
	"github.com/eveisesi/neo/mysql/boiler"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type allianceRepository struct {
	db *sqlx.DB
}

func NewAllianceRepository(db *sqlx.DB) neo.AllianceRespository {
	return &allianceRepository{
		db,
	}
}

func (r *allianceRepository) Alliance(ctx context.Context, id uint64) (*neo.Alliance, error) {

	var alliance = neo.Alliance{}
	err := boiler.Alliances(
		boiler.AllianceWhere.ID.EQ(id),
	).Bind(ctx, r.db, &alliance)

	return &alliance, err

}

func (r *allianceRepository) AlliancesByAllianceIDs(ctx context.Context, ids []uint64) ([]*neo.Alliance, error) {

	var alliances = make([]*neo.Alliance, 0)
	err := boiler.Alliances(
		qm.WhereIn(
			fmt.Sprintf(
				"%s IN ?",
				boiler.AllianceColumns.ID,
			),
			convertSliceUint64ToSliceInterface(ids)...,
		),
	).Bind(ctx, r.db, &alliances)

	return alliances, err
}
