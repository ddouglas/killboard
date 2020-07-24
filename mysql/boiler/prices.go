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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Price is an object representing the database table.
type Price struct {
	TypeID    uint64    `boil:"type_id" json:"typeID" toml:"typeID" yaml:"typeID"`
	Date      time.Time `boil:"date" json:"date" toml:"date" yaml:"date"`
	Price     float64   `boil:"price" json:"price" toml:"price" yaml:"price"`
	CreatedAt time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *priceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L priceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PriceColumns = struct {
	TypeID    string
	Date      string
	Price     string
	CreatedAt string
	UpdatedAt string
}{
	TypeID:    "type_id",
	Date:      "date",
	Price:     "price",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var PriceWhere = struct {
	TypeID    whereHelperuint64
	Date      whereHelpertime_Time
	Price     whereHelperfloat64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	TypeID:    whereHelperuint64{field: "`prices`.`type_id`"},
	Date:      whereHelpertime_Time{field: "`prices`.`date`"},
	Price:     whereHelperfloat64{field: "`prices`.`price`"},
	CreatedAt: whereHelpertime_Time{field: "`prices`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`prices`.`updated_at`"},
}

// PriceRels is where relationship names are stored.
var PriceRels = struct {
}{}

// priceR is where relationships are stored.
type priceR struct {
}

// NewStruct creates a new relationship struct
func (*priceR) NewStruct() *priceR {
	return &priceR{}
}

// priceL is where Load methods for each relationship are stored.
type priceL struct{}

var (
	priceAllColumns            = []string{"type_id", "date", "price", "created_at", "updated_at"}
	priceColumnsWithoutDefault = []string{"type_id", "date", "created_at", "updated_at"}
	priceColumnsWithDefault    = []string{"price"}
	pricePrimaryKeyColumns     = []string{"type_id", "date"}
)

type (
	// PriceSlice is an alias for a slice of pointers to Price.
	// This should generally be used opposed to []Price.
	PriceSlice []*Price

	priceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	priceType                 = reflect.TypeOf(&Price{})
	priceMapping              = queries.MakeStructMapping(priceType)
	pricePrimaryKeyMapping, _ = queries.BindMapping(priceType, priceMapping, pricePrimaryKeyColumns)
	priceInsertCacheMut       sync.RWMutex
	priceInsertCache          = make(map[string]insertCache)
	priceUpdateCacheMut       sync.RWMutex
	priceUpdateCache          = make(map[string]updateCache)
	priceUpsertCacheMut       sync.RWMutex
	priceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single price record from the query.
func (q priceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Price, error) {
	o := &Price{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for prices")
	}

	return o, nil
}

// All returns all Price records from the query.
func (q priceQuery) All(ctx context.Context, exec boil.ContextExecutor) (PriceSlice, error) {
	var o []*Price

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to Price slice")
	}

	return o, nil
}

// Count returns the count of all Price records in the query.
func (q priceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count prices rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q priceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if prices exists")
	}

	return count > 0, nil
}

// Prices retrieves all the records using an executor.
func Prices(mods ...qm.QueryMod) priceQuery {
	mods = append(mods, qm.From("`prices`"))
	return priceQuery{NewQuery(mods...)}
}

// FindPrice retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPrice(ctx context.Context, exec boil.ContextExecutor, typeID uint64, date time.Time, selectCols ...string) (*Price, error) {
	priceObj := &Price{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `prices` where `type_id`=? AND `date`=?", sel,
	)

	q := queries.Raw(query, typeID, date)

	err := q.Bind(ctx, exec, priceObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from prices")
	}

	return priceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Price) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns, ignore bool) error {
	if o == nil {
		return errors.New("boiler: no prices provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(priceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	priceInsertCacheMut.RLock()
	cache, cached := priceInsertCache[key]
	priceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			priceAllColumns,
			priceColumnsWithDefault,
			priceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(priceType, priceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(priceType, priceMapping, returnColumns)
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
			cache.query = fmt.Sprintf("%s INTO `prices` (`%s`) %%sVALUES (%s)%%s", insert, strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = fmt.Sprintf("%s INTO `prices` () VALUES ()%s%s", insert)
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `prices` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, pricePrimaryKeyColumns))
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
		return errors.Wrap(err, "boiler: unable to insert into prices")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.TypeID,
		o.Date,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for prices")
	}

CacheNoHooks:
	if !cached {
		priceInsertCacheMut.Lock()
		priceInsertCache[key] = cache
		priceInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Price.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Price) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	priceUpdateCacheMut.RLock()
	cache, cached := priceUpdateCache[key]
	priceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			priceAllColumns,
			pricePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update prices, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `prices` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, pricePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(priceType, priceMapping, append(wl, pricePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update prices row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for prices")
	}

	if !cached {
		priceUpdateCacheMut.Lock()
		priceUpdateCache[key] = cache
		priceUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q priceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for prices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for prices")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PriceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `prices` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pricePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in price slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all price")
	}
	return rowsAff, nil
}

var mySQLPriceUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Price) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no prices provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(priceColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPriceUniqueColumns, o)

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

	priceUpsertCacheMut.RLock()
	cache, cached := priceUpsertCache[key]
	priceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			priceAllColumns,
			priceColumnsWithDefault,
			priceColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			priceAllColumns,
			pricePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("boiler: unable to upsert prices, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "prices", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `prices` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(priceType, priceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(priceType, priceMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert for prices")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(priceType, priceMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to retrieve unique values for prices")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for prices")
	}

CacheNoHooks:
	if !cached {
		priceUpsertCacheMut.Lock()
		priceUpsertCache[key] = cache
		priceUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Price record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Price) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no Price provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pricePrimaryKeyMapping)
	sql := "DELETE FROM `prices` WHERE `type_id`=? AND `date`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from prices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for prices")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q priceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no priceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from prices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for prices")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PriceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `prices` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pricePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from price slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for prices")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Price) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPrice(ctx, exec, o.TypeID, o.Date)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PriceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PriceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `prices`.* FROM `prices` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pricePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in PriceSlice")
	}

	*o = slice

	return nil
}

// PriceExists checks if the Price row exists.
func PriceExists(ctx context.Context, exec boil.ContextExecutor, typeID uint64, date time.Time) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `prices` where `type_id`=? AND `date`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, typeID, date)
	}
	row := exec.QueryRowContext(ctx, sql, typeID, date)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if prices exists")
	}

	return exists, nil
}
