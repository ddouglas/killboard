package mysql

import (
	"context"
	"fmt"

	"github.com/eveisesi/neo"
	"github.com/eveisesi/neo/mysql/boiler"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type universeRepository struct {
	db *sqlx.DB
}

func NewUniverseRepository(db *sqlx.DB) neo.UniverseRepository {
	return &universeRepository{
		db,
	}
}

func (r *universeRepository) Type(ctx context.Context, id uint64) (*neo.Type, error) {

	var invType = neo.Type{}
	err := boiler.Types(
		boiler.TypeWhere.ID.EQ(id),
	).Bind(ctx, r.db, &invType)

	return &invType, err

}

func (r *universeRepository) TypesByTypeIDs(ctx context.Context, ids []uint64) ([]*neo.Type, error) {

	var invTypes = make([]*neo.Type, 0)
	err := boiler.Types(
		qm.WhereIn(
			fmt.Sprintf(
				"%s IN ?",
				boiler.TypeColumns.ID,
			),
			convertSliceUint64ToSliceInterface(ids)...,
		),
	).Bind(ctx, r.db, &invTypes)

	return invTypes, err
}

func (r *universeRepository) SolarSystem(ctx context.Context, id uint64) (*neo.SolarSystem, error) {

	var solarSystem = neo.SolarSystem{}
	err := boiler.SolarSystems(
		boiler.SolarSystemWhere.ID.EQ(id),
	).Bind(ctx, r.db, &solarSystem)

	return &solarSystem, err

}

func (r *universeRepository) SolarSystemsBySolarSystemIDs(ctx context.Context, ids []uint64) ([]*neo.SolarSystem, error) {

	var systems = make([]*neo.SolarSystem, 0)
	err := boiler.SolarSystems(
		qm.WhereIn(
			fmt.Sprintf(
				"%s IN ?",
				boiler.SolarSystemColumns.ID,
			),
			convertSliceUint64ToSliceInterface(ids)...,
		),
	).Bind(ctx, r.db, &systems)

	return systems, err
}

func (r *universeRepository) TypeFlag(ctx context.Context, id uint64) (*neo.TypeFlag, error) {

	flag := neo.TypeFlag{}
	err := boiler.TypeFlags(
		boiler.TypeFlagWhere.ID.EQ(id),
	).Bind(ctx, r.db, &flag)

	return &flag, err

}
func (r *universeRepository) TypeFlagsByTypeFlagIDs(ctx context.Context, ids []uint64) ([]*neo.TypeFlag, error) {

	flags := make([]*neo.TypeFlag, 0)
	err := boiler.TypeFlags(
		qm.WhereIn(
			fmt.Sprintf(
				"%s IN ?",
				boiler.TypeFlagColumns.ID,
			),
			convertSliceUint64ToSliceInterface(ids)...,
		),
	).Bind(ctx, r.db, &flags)

	return flags, err

}
