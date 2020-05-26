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

// Corporation is an object representing the database table.
type Corporation struct {
	ID               uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name             string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Ticker           string      `boil:"ticker" json:"ticker" toml:"ticker" yaml:"ticker"`
	AllianceID       null.Uint64 `boil:"alliance_id" json:"allianceID,omitempty" toml:"allianceID" yaml:"allianceID,omitempty"`
	NotModifiedCount uint        `boil:"not_modified_count" json:"notModifiedCount" toml:"notModifiedCount" yaml:"notModifiedCount"`
	UpdatePriority   uint        `boil:"update_priority" json:"updatePriority" toml:"updatePriority" yaml:"updatePriority"`
	Etag             string      `boil:"etag" json:"etag" toml:"etag" yaml:"etag"`
	CachedUntil      time.Time   `boil:"cached_until" json:"cachedUntil" toml:"cachedUntil" yaml:"cachedUntil"`
	CreatedAt        time.Time   `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt        time.Time   `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *corporationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L corporationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CorporationColumns = struct {
	ID               string
	Name             string
	Ticker           string
	AllianceID       string
	NotModifiedCount string
	UpdatePriority   string
	Etag             string
	CachedUntil      string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "id",
	Name:             "name",
	Ticker:           "ticker",
	AllianceID:       "alliance_id",
	NotModifiedCount: "not_modified_count",
	UpdatePriority:   "update_priority",
	Etag:             "etag",
	CachedUntil:      "cached_until",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// Generated where

var CorporationWhere = struct {
	ID               whereHelperuint64
	Name             whereHelperstring
	Ticker           whereHelperstring
	AllianceID       whereHelpernull_Uint64
	NotModifiedCount whereHelperuint
	UpdatePriority   whereHelperuint
	Etag             whereHelperstring
	CachedUntil      whereHelpertime_Time
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
}{
	ID:               whereHelperuint64{field: "`corporations`.`id`"},
	Name:             whereHelperstring{field: "`corporations`.`name`"},
	Ticker:           whereHelperstring{field: "`corporations`.`ticker`"},
	AllianceID:       whereHelpernull_Uint64{field: "`corporations`.`alliance_id`"},
	NotModifiedCount: whereHelperuint{field: "`corporations`.`not_modified_count`"},
	UpdatePriority:   whereHelperuint{field: "`corporations`.`update_priority`"},
	Etag:             whereHelperstring{field: "`corporations`.`etag`"},
	CachedUntil:      whereHelpertime_Time{field: "`corporations`.`cached_until`"},
	CreatedAt:        whereHelpertime_Time{field: "`corporations`.`created_at`"},
	UpdatedAt:        whereHelpertime_Time{field: "`corporations`.`updated_at`"},
}

// CorporationRels is where relationship names are stored.
var CorporationRels = struct {
}{}

// corporationR is where relationships are stored.
type corporationR struct {
}

// NewStruct creates a new relationship struct
func (*corporationR) NewStruct() *corporationR {
	return &corporationR{}
}

// corporationL is where Load methods for each relationship are stored.
type corporationL struct{}

var (
	corporationAllColumns            = []string{"id", "name", "ticker", "alliance_id", "not_modified_count", "update_priority", "etag", "cached_until", "created_at", "updated_at"}
	corporationColumnsWithoutDefault = []string{"id", "name", "ticker", "alliance_id", "etag", "cached_until", "created_at", "updated_at"}
	corporationColumnsWithDefault    = []string{"not_modified_count", "update_priority"}
	corporationPrimaryKeyColumns     = []string{"id"}
)

type (
	// CorporationSlice is an alias for a slice of pointers to Corporation.
	// This should generally be used opposed to []Corporation.
	CorporationSlice []*Corporation

	corporationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	corporationType                 = reflect.TypeOf(&Corporation{})
	corporationMapping              = queries.MakeStructMapping(corporationType)
	corporationPrimaryKeyMapping, _ = queries.BindMapping(corporationType, corporationMapping, corporationPrimaryKeyColumns)
	corporationInsertCacheMut       sync.RWMutex
	corporationInsertCache          = make(map[string]insertCache)
	corporationUpdateCacheMut       sync.RWMutex
	corporationUpdateCache          = make(map[string]updateCache)
	corporationUpsertCacheMut       sync.RWMutex
	corporationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single corporation record from the query.
func (q corporationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Corporation, error) {
	o := &Corporation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for corporations")
	}

	return o, nil
}

// All returns all Corporation records from the query.
func (q corporationQuery) All(ctx context.Context, exec boil.ContextExecutor) (CorporationSlice, error) {
	var o []*Corporation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to Corporation slice")
	}

	return o, nil
}

// Count returns the count of all Corporation records in the query.
func (q corporationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count corporations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q corporationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if corporations exists")
	}

	return count > 0, nil
}

// Corporations retrieves all the records using an executor.
func Corporations(mods ...qm.QueryMod) corporationQuery {
	mods = append(mods, qm.From("`corporations`"))
	return corporationQuery{NewQuery(mods...)}
}

// FindCorporation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCorporation(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*Corporation, error) {
	corporationObj := &Corporation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `corporations` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, corporationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from corporations")
	}

	return corporationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Corporation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no corporations provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(corporationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	corporationInsertCacheMut.RLock()
	cache, cached := corporationInsertCache[key]
	corporationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			corporationAllColumns,
			corporationColumnsWithDefault,
			corporationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(corporationType, corporationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(corporationType, corporationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `corporations` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `corporations` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `corporations` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, corporationPrimaryKeyColumns))
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
		return errors.Wrap(err, "boiler: unable to insert into corporations")
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
		return errors.Wrap(err, "boiler: unable to populate default values for corporations")
	}

CacheNoHooks:
	if !cached {
		corporationInsertCacheMut.Lock()
		corporationInsertCache[key] = cache
		corporationInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Corporation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Corporation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	corporationUpdateCacheMut.RLock()
	cache, cached := corporationUpdateCache[key]
	corporationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			corporationAllColumns,
			corporationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update corporations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `corporations` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, corporationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(corporationType, corporationMapping, append(wl, corporationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update corporations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for corporations")
	}

	if !cached {
		corporationUpdateCacheMut.Lock()
		corporationUpdateCache[key] = cache
		corporationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q corporationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for corporations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for corporations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CorporationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), corporationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `corporations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, corporationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in corporation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all corporation")
	}
	return rowsAff, nil
}

var mySQLCorporationUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Corporation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no corporations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(corporationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCorporationUniqueColumns, o)

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

	corporationUpsertCacheMut.RLock()
	cache, cached := corporationUpsertCache[key]
	corporationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			corporationAllColumns,
			corporationColumnsWithDefault,
			corporationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			corporationAllColumns,
			corporationPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("boiler: unable to upsert corporations, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "corporations", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `corporations` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(corporationType, corporationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(corporationType, corporationMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert for corporations")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(corporationType, corporationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to retrieve unique values for corporations")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for corporations")
	}

CacheNoHooks:
	if !cached {
		corporationUpsertCacheMut.Lock()
		corporationUpsertCache[key] = cache
		corporationUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Corporation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Corporation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no Corporation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), corporationPrimaryKeyMapping)
	sql := "DELETE FROM `corporations` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from corporations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for corporations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q corporationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no corporationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from corporations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for corporations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CorporationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), corporationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `corporations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, corporationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from corporation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for corporations")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Corporation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCorporation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CorporationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CorporationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), corporationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `corporations`.* FROM `corporations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, corporationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in CorporationSlice")
	}

	*o = slice

	return nil
}

// CorporationExists checks if the Corporation row exists.
func CorporationExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `corporations` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if corporations exists")
	}

	return exists, nil
}
