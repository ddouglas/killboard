// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiler

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// SolarSystem is an object representing the database table.
type SolarSystem struct {
	ID              uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name            string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	ConstellationID uint64      `boil:"constellation_id" json:"constellationID" toml:"constellationID" yaml:"constellationID"`
	FactionID       null.Uint64 `boil:"faction_id" json:"factionID,omitempty" toml:"factionID" yaml:"factionID,omitempty"`
	SunTypeID       null.Uint64 `boil:"sun_type_id" json:"sunTypeID,omitempty" toml:"sunTypeID" yaml:"sunTypeID,omitempty"`
	PosX            float64     `boil:"pos_x" json:"posX" toml:"posX" yaml:"posX"`
	PosY            float64     `boil:"pos_y" json:"posY" toml:"posY" yaml:"posY"`
	PosZ            float64     `boil:"pos_z" json:"posZ" toml:"posZ" yaml:"posZ"`
	Security        float64     `boil:"security" json:"security" toml:"security" yaml:"security"`
	CreatedAt       time.Time   `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt       time.Time   `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *solarSystemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L solarSystemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SolarSystemColumns = struct {
	ID              string
	Name            string
	ConstellationID string
	FactionID       string
	SunTypeID       string
	PosX            string
	PosY            string
	PosZ            string
	Security        string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "id",
	Name:            "name",
	ConstellationID: "constellation_id",
	FactionID:       "faction_id",
	SunTypeID:       "sun_type_id",
	PosX:            "pos_x",
	PosY:            "pos_y",
	PosZ:            "pos_z",
	Security:        "security",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// Generated where

var SolarSystemWhere = struct {
	ID              whereHelperuint64
	Name            whereHelperstring
	ConstellationID whereHelperuint64
	FactionID       whereHelpernull_Uint64
	SunTypeID       whereHelpernull_Uint64
	PosX            whereHelperfloat64
	PosY            whereHelperfloat64
	PosZ            whereHelperfloat64
	Security        whereHelperfloat64
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpertime_Time
}{
	ID:              whereHelperuint64{field: "`solar_systems`.`id`"},
	Name:            whereHelperstring{field: "`solar_systems`.`name`"},
	ConstellationID: whereHelperuint64{field: "`solar_systems`.`constellation_id`"},
	FactionID:       whereHelpernull_Uint64{field: "`solar_systems`.`faction_id`"},
	SunTypeID:       whereHelpernull_Uint64{field: "`solar_systems`.`sun_type_id`"},
	PosX:            whereHelperfloat64{field: "`solar_systems`.`pos_x`"},
	PosY:            whereHelperfloat64{field: "`solar_systems`.`pos_y`"},
	PosZ:            whereHelperfloat64{field: "`solar_systems`.`pos_z`"},
	Security:        whereHelperfloat64{field: "`solar_systems`.`security`"},
	CreatedAt:       whereHelpertime_Time{field: "`solar_systems`.`created_at`"},
	UpdatedAt:       whereHelpertime_Time{field: "`solar_systems`.`updated_at`"},
}

// SolarSystemRels is where relationship names are stored.
var SolarSystemRels = struct {
}{}

// solarSystemR is where relationships are stored.
type solarSystemR struct {
}

// NewStruct creates a new relationship struct
func (*solarSystemR) NewStruct() *solarSystemR {
	return &solarSystemR{}
}

// solarSystemL is where Load methods for each relationship are stored.
type solarSystemL struct{}

var (
	solarSystemAllColumns            = []string{"id", "name", "constellation_id", "faction_id", "sun_type_id", "pos_x", "pos_y", "pos_z", "security", "created_at", "updated_at"}
	solarSystemColumnsWithoutDefault = []string{"id", "name", "constellation_id", "faction_id", "sun_type_id", "pos_x", "pos_y", "pos_z", "security", "created_at", "updated_at"}
	solarSystemColumnsWithDefault    = []string{}
	solarSystemPrimaryKeyColumns     = []string{"id"}
)

type (
	// SolarSystemSlice is an alias for a slice of pointers to SolarSystem.
	// This should generally be used opposed to []SolarSystem.
	SolarSystemSlice []*SolarSystem

	solarSystemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	solarSystemType                 = reflect.TypeOf(&SolarSystem{})
	solarSystemMapping              = queries.MakeStructMapping(solarSystemType)
	solarSystemPrimaryKeyMapping, _ = queries.BindMapping(solarSystemType, solarSystemMapping, solarSystemPrimaryKeyColumns)
	solarSystemInsertCacheMut       sync.RWMutex
	solarSystemInsertCache          = make(map[string]insertCache)
	solarSystemUpdateCacheMut       sync.RWMutex
	solarSystemUpdateCache          = make(map[string]updateCache)
	solarSystemUpsertCacheMut       sync.RWMutex
	solarSystemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single solarSystem record from the query.
func (q solarSystemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SolarSystem, error) {
	o := &SolarSystem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for solar_systems")
	}

	return o, nil
}

// All returns all SolarSystem records from the query.
func (q solarSystemQuery) All(ctx context.Context, exec boil.ContextExecutor) (SolarSystemSlice, error) {
	var o []*SolarSystem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to SolarSystem slice")
	}

	return o, nil
}

// Count returns the count of all SolarSystem records in the query.
func (q solarSystemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count solar_systems rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q solarSystemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if solar_systems exists")
	}

	return count > 0, nil
}

// SolarSystems retrieves all the records using an executor.
func SolarSystems(mods ...qm.QueryMod) solarSystemQuery {
	mods = append(mods, qm.From("`solar_systems`"))
	return solarSystemQuery{NewQuery(mods...)}
}

// FindSolarSystem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSolarSystem(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*SolarSystem, error) {
	solarSystemObj := &SolarSystem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `solar_systems` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, solarSystemObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from solar_systems")
	}

	return solarSystemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SolarSystem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns, ignore bool) error {
	if o == nil {
		return errors.New("boiler: no solar_systems provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(solarSystemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	solarSystemInsertCacheMut.RLock()
	cache, cached := solarSystemInsertCache[key]
	solarSystemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			solarSystemAllColumns,
			solarSystemColumnsWithDefault,
			solarSystemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(solarSystemType, solarSystemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(solarSystemType, solarSystemMapping, returnColumns)
		if err != nil {
			return err
		}
		insert := "INSERT%s"
		if ignore {
			insert = fmt.Sprintf(insert, " IGNORE")
		} else {
			insert = fmt.Sprintf(insert, "")
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("%s INTO `solar_systems` (`%s`) %%sVALUES (%s)%%s", insert, strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = fmt.Sprintf("%s INTO `solar_systems` () VALUES ()%s%s", insert)
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `solar_systems` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, solarSystemPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "boiler: unable to insert into solar_systems")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for solar_systems")
	}

CacheNoHooks:
	if !cached {
		solarSystemInsertCacheMut.Lock()
		solarSystemInsertCache[key] = cache
		solarSystemInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the SolarSystem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SolarSystem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	solarSystemUpdateCacheMut.RLock()
	cache, cached := solarSystemUpdateCache[key]
	solarSystemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			solarSystemAllColumns,
			solarSystemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update solar_systems, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `solar_systems` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, solarSystemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(solarSystemType, solarSystemMapping, append(wl, solarSystemPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update solar_systems row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for solar_systems")
	}

	if !cached {
		solarSystemUpdateCacheMut.Lock()
		solarSystemUpdateCache[key] = cache
		solarSystemUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q solarSystemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for solar_systems")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for solar_systems")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SolarSystemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("boiler: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solarSystemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `solar_systems` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, solarSystemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in solarSystem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all solarSystem")
	}
	return rowsAff, nil
}

var mySQLSolarSystemUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SolarSystem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no solar_systems provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(solarSystemColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSolarSystemUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	solarSystemUpsertCacheMut.RLock()
	cache, cached := solarSystemUpsertCache[key]
	solarSystemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			solarSystemAllColumns,
			solarSystemColumnsWithDefault,
			solarSystemColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			solarSystemAllColumns,
			solarSystemPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("boiler: unable to upsert solar_systems, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "solar_systems", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `solar_systems` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(solarSystemType, solarSystemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(solarSystemType, solarSystemMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "boiler: unable to upsert for solar_systems")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(solarSystemType, solarSystemMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to retrieve unique values for solar_systems")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for solar_systems")
	}

CacheNoHooks:
	if !cached {
		solarSystemUpsertCacheMut.Lock()
		solarSystemUpsertCache[key] = cache
		solarSystemUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single SolarSystem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SolarSystem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no SolarSystem provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), solarSystemPrimaryKeyMapping)
	sql := "DELETE FROM `solar_systems` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from solar_systems")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for solar_systems")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q solarSystemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no solarSystemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from solar_systems")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for solar_systems")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SolarSystemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solarSystemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `solar_systems` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, solarSystemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from solarSystem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for solar_systems")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SolarSystem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSolarSystem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SolarSystemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SolarSystemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solarSystemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `solar_systems`.* FROM `solar_systems` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, solarSystemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in SolarSystemSlice")
	}

	*o = slice

	return nil
}

// SolarSystemExists checks if the SolarSystem row exists.
func SolarSystemExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `solar_systems` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if solar_systems exists")
	}

	return exists, nil
}
