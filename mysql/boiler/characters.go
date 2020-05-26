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

// Character is an object representing the database table.
type Character struct {
	ID               uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name             string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	CorporationID    uint64      `boil:"corporation_id" json:"corporationID" toml:"corporationID" yaml:"corporationID"`
	AllianceID       null.Uint64 `boil:"alliance_id" json:"allianceID,omitempty" toml:"allianceID" yaml:"allianceID,omitempty"`
	FactionID        null.Uint64 `boil:"faction_id" json:"factionID,omitempty" toml:"factionID" yaml:"factionID,omitempty"`
	NotModifiedCount uint        `boil:"not_modified_count" json:"notModifiedCount" toml:"notModifiedCount" yaml:"notModifiedCount"`
	UpdatePriority   uint        `boil:"update_priority" json:"updatePriority" toml:"updatePriority" yaml:"updatePriority"`
	Etag             string      `boil:"etag" json:"etag" toml:"etag" yaml:"etag"`
	CachedUntil      time.Time   `boil:"cached_until" json:"cachedUntil" toml:"cachedUntil" yaml:"cachedUntil"`
	CreatedAt        time.Time   `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt        time.Time   `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *characterR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L characterL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CharacterColumns = struct {
	ID               string
	Name             string
	CorporationID    string
	AllianceID       string
	FactionID        string
	NotModifiedCount string
	UpdatePriority   string
	Etag             string
	CachedUntil      string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "id",
	Name:             "name",
	CorporationID:    "corporation_id",
	AllianceID:       "alliance_id",
	FactionID:        "faction_id",
	NotModifiedCount: "not_modified_count",
	UpdatePriority:   "update_priority",
	Etag:             "etag",
	CachedUntil:      "cached_until",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// Generated where

type whereHelpernull_Uint64 struct{ field string }

func (w whereHelpernull_Uint64) EQ(x null.Uint64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Uint64) NEQ(x null.Uint64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Uint64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Uint64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Uint64) LT(x null.Uint64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Uint64) LTE(x null.Uint64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Uint64) GT(x null.Uint64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Uint64) GTE(x null.Uint64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CharacterWhere = struct {
	ID               whereHelperuint64
	Name             whereHelperstring
	CorporationID    whereHelperuint64
	AllianceID       whereHelpernull_Uint64
	FactionID        whereHelpernull_Uint64
	NotModifiedCount whereHelperuint
	UpdatePriority   whereHelperuint
	Etag             whereHelperstring
	CachedUntil      whereHelpertime_Time
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
}{
	ID:               whereHelperuint64{field: "`characters`.`id`"},
	Name:             whereHelperstring{field: "`characters`.`name`"},
	CorporationID:    whereHelperuint64{field: "`characters`.`corporation_id`"},
	AllianceID:       whereHelpernull_Uint64{field: "`characters`.`alliance_id`"},
	FactionID:        whereHelpernull_Uint64{field: "`characters`.`faction_id`"},
	NotModifiedCount: whereHelperuint{field: "`characters`.`not_modified_count`"},
	UpdatePriority:   whereHelperuint{field: "`characters`.`update_priority`"},
	Etag:             whereHelperstring{field: "`characters`.`etag`"},
	CachedUntil:      whereHelpertime_Time{field: "`characters`.`cached_until`"},
	CreatedAt:        whereHelpertime_Time{field: "`characters`.`created_at`"},
	UpdatedAt:        whereHelpertime_Time{field: "`characters`.`updated_at`"},
}

// CharacterRels is where relationship names are stored.
var CharacterRels = struct {
}{}

// characterR is where relationships are stored.
type characterR struct {
}

// NewStruct creates a new relationship struct
func (*characterR) NewStruct() *characterR {
	return &characterR{}
}

// characterL is where Load methods for each relationship are stored.
type characterL struct{}

var (
	characterAllColumns            = []string{"id", "name", "corporation_id", "alliance_id", "faction_id", "not_modified_count", "update_priority", "etag", "cached_until", "created_at", "updated_at"}
	characterColumnsWithoutDefault = []string{"id", "name", "corporation_id", "alliance_id", "faction_id", "etag", "cached_until", "created_at", "updated_at"}
	characterColumnsWithDefault    = []string{"not_modified_count", "update_priority"}
	characterPrimaryKeyColumns     = []string{"id"}
)

type (
	// CharacterSlice is an alias for a slice of pointers to Character.
	// This should generally be used opposed to []Character.
	CharacterSlice []*Character

	characterQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	characterType                 = reflect.TypeOf(&Character{})
	characterMapping              = queries.MakeStructMapping(characterType)
	characterPrimaryKeyMapping, _ = queries.BindMapping(characterType, characterMapping, characterPrimaryKeyColumns)
	characterInsertCacheMut       sync.RWMutex
	characterInsertCache          = make(map[string]insertCache)
	characterUpdateCacheMut       sync.RWMutex
	characterUpdateCache          = make(map[string]updateCache)
	characterUpsertCacheMut       sync.RWMutex
	characterUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single character record from the query.
func (q characterQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Character, error) {
	o := &Character{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for characters")
	}

	return o, nil
}

// All returns all Character records from the query.
func (q characterQuery) All(ctx context.Context, exec boil.ContextExecutor) (CharacterSlice, error) {
	var o []*Character

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to Character slice")
	}

	return o, nil
}

// Count returns the count of all Character records in the query.
func (q characterQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count characters rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q characterQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if characters exists")
	}

	return count > 0, nil
}

// Characters retrieves all the records using an executor.
func Characters(mods ...qm.QueryMod) characterQuery {
	mods = append(mods, qm.From("`characters`"))
	return characterQuery{NewQuery(mods...)}
}

// FindCharacter retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCharacter(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*Character, error) {
	characterObj := &Character{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `characters` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, characterObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from characters")
	}

	return characterObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Character) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no characters provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(characterColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	characterInsertCacheMut.RLock()
	cache, cached := characterInsertCache[key]
	characterInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			characterAllColumns,
			characterColumnsWithDefault,
			characterColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(characterType, characterMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(characterType, characterMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `characters` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `characters` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `characters` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, characterPrimaryKeyColumns))
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
		return errors.Wrap(err, "boiler: unable to insert into characters")
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
		return errors.Wrap(err, "boiler: unable to populate default values for characters")
	}

CacheNoHooks:
	if !cached {
		characterInsertCacheMut.Lock()
		characterInsertCache[key] = cache
		characterInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Character.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Character) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	characterUpdateCacheMut.RLock()
	cache, cached := characterUpdateCache[key]
	characterUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			characterAllColumns,
			characterPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update characters, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `characters` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, characterPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(characterType, characterMapping, append(wl, characterPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update characters row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for characters")
	}

	if !cached {
		characterUpdateCacheMut.Lock()
		characterUpdateCache[key] = cache
		characterUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q characterQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for characters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for characters")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CharacterSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `characters` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in character slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all character")
	}
	return rowsAff, nil
}

var mySQLCharacterUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Character) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no characters provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(characterColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCharacterUniqueColumns, o)

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

	characterUpsertCacheMut.RLock()
	cache, cached := characterUpsertCache[key]
	characterUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			characterAllColumns,
			characterColumnsWithDefault,
			characterColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			characterAllColumns,
			characterPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("boiler: unable to upsert characters, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "characters", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `characters` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(characterType, characterMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(characterType, characterMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert for characters")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(characterType, characterMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to retrieve unique values for characters")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for characters")
	}

CacheNoHooks:
	if !cached {
		characterUpsertCacheMut.Lock()
		characterUpsertCache[key] = cache
		characterUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Character record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Character) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no Character provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), characterPrimaryKeyMapping)
	sql := "DELETE FROM `characters` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from characters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for characters")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q characterQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no characterQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from characters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for characters")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CharacterSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `characters` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from character slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for characters")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Character) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCharacter(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CharacterSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CharacterSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `characters`.* FROM `characters` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in CharacterSlice")
	}

	*o = slice

	return nil
}

// CharacterExists checks if the Character row exists.
func CharacterExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `characters` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if characters exists")
	}

	return exists, nil
}
