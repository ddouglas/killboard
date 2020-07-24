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

// KillmailItem is an object representing the database table.
type KillmailItem struct {
	ID                uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ParentID          null.Uint64 `boil:"parent_id" json:"parentID,omitempty" toml:"parentID" yaml:"parentID,omitempty"`
	KillmailID        uint64      `boil:"killmail_id" json:"killmailID" toml:"killmailID" yaml:"killmailID"`
	Flag              uint64      `boil:"flag" json:"flag" toml:"flag" yaml:"flag"`
	ItemTypeID        uint64      `boil:"item_type_id" json:"itemTypeID" toml:"itemTypeID" yaml:"itemTypeID"`
	QuantityDropped   null.Uint64 `boil:"quantity_dropped" json:"quantityDropped,omitempty" toml:"quantityDropped" yaml:"quantityDropped,omitempty"`
	QuantityDestroyed null.Uint64 `boil:"quantity_destroyed" json:"quantityDestroyed,omitempty" toml:"quantityDestroyed" yaml:"quantityDestroyed,omitempty"`
	ItemValue         float64     `boil:"item_value" json:"itemValue" toml:"itemValue" yaml:"itemValue"`
	Singleton         uint64      `boil:"singleton" json:"singleton" toml:"singleton" yaml:"singleton"`
	IsParent          bool        `boil:"is_parent" json:"isParent" toml:"isParent" yaml:"isParent"`
	CreatedAt         time.Time   `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt         time.Time   `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *killmailItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L killmailItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var KillmailItemColumns = struct {
	ID                string
	ParentID          string
	KillmailID        string
	Flag              string
	ItemTypeID        string
	QuantityDropped   string
	QuantityDestroyed string
	ItemValue         string
	Singleton         string
	IsParent          string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	ParentID:          "parent_id",
	KillmailID:        "killmail_id",
	Flag:              "flag",
	ItemTypeID:        "item_type_id",
	QuantityDropped:   "quantity_dropped",
	QuantityDestroyed: "quantity_destroyed",
	ItemValue:         "item_value",
	Singleton:         "singleton",
	IsParent:          "is_parent",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// Generated where

var KillmailItemWhere = struct {
	ID                whereHelperuint64
	ParentID          whereHelpernull_Uint64
	KillmailID        whereHelperuint64
	Flag              whereHelperuint64
	ItemTypeID        whereHelperuint64
	QuantityDropped   whereHelpernull_Uint64
	QuantityDestroyed whereHelpernull_Uint64
	ItemValue         whereHelperfloat64
	Singleton         whereHelperuint64
	IsParent          whereHelperbool
	CreatedAt         whereHelpertime_Time
	UpdatedAt         whereHelpertime_Time
}{
	ID:                whereHelperuint64{field: "`killmail_items`.`id`"},
	ParentID:          whereHelpernull_Uint64{field: "`killmail_items`.`parent_id`"},
	KillmailID:        whereHelperuint64{field: "`killmail_items`.`killmail_id`"},
	Flag:              whereHelperuint64{field: "`killmail_items`.`flag`"},
	ItemTypeID:        whereHelperuint64{field: "`killmail_items`.`item_type_id`"},
	QuantityDropped:   whereHelpernull_Uint64{field: "`killmail_items`.`quantity_dropped`"},
	QuantityDestroyed: whereHelpernull_Uint64{field: "`killmail_items`.`quantity_destroyed`"},
	ItemValue:         whereHelperfloat64{field: "`killmail_items`.`item_value`"},
	Singleton:         whereHelperuint64{field: "`killmail_items`.`singleton`"},
	IsParent:          whereHelperbool{field: "`killmail_items`.`is_parent`"},
	CreatedAt:         whereHelpertime_Time{field: "`killmail_items`.`created_at`"},
	UpdatedAt:         whereHelpertime_Time{field: "`killmail_items`.`updated_at`"},
}

// KillmailItemRels is where relationship names are stored.
var KillmailItemRels = struct {
	Killmail string
}{
	Killmail: "Killmail",
}

// killmailItemR is where relationships are stored.
type killmailItemR struct {
	Killmail *Killmail
}

// NewStruct creates a new relationship struct
func (*killmailItemR) NewStruct() *killmailItemR {
	return &killmailItemR{}
}

// killmailItemL is where Load methods for each relationship are stored.
type killmailItemL struct{}

var (
	killmailItemAllColumns            = []string{"id", "parent_id", "killmail_id", "flag", "item_type_id", "quantity_dropped", "quantity_destroyed", "item_value", "singleton", "is_parent", "created_at", "updated_at"}
	killmailItemColumnsWithoutDefault = []string{"parent_id", "killmail_id", "flag", "item_type_id", "quantity_dropped", "quantity_destroyed", "singleton", "is_parent", "created_at", "updated_at"}
	killmailItemColumnsWithDefault    = []string{"id", "item_value"}
	killmailItemPrimaryKeyColumns     = []string{"id"}
)

type (
	// KillmailItemSlice is an alias for a slice of pointers to KillmailItem.
	// This should generally be used opposed to []KillmailItem.
	KillmailItemSlice []*KillmailItem

	killmailItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	killmailItemType                 = reflect.TypeOf(&KillmailItem{})
	killmailItemMapping              = queries.MakeStructMapping(killmailItemType)
	killmailItemPrimaryKeyMapping, _ = queries.BindMapping(killmailItemType, killmailItemMapping, killmailItemPrimaryKeyColumns)
	killmailItemInsertCacheMut       sync.RWMutex
	killmailItemInsertCache          = make(map[string]insertCache)
	killmailItemUpdateCacheMut       sync.RWMutex
	killmailItemUpdateCache          = make(map[string]updateCache)
	killmailItemUpsertCacheMut       sync.RWMutex
	killmailItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single killmailItem record from the query.
func (q killmailItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*KillmailItem, error) {
	o := &KillmailItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for killmail_items")
	}

	return o, nil
}

// All returns all KillmailItem records from the query.
func (q killmailItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (KillmailItemSlice, error) {
	var o []*KillmailItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to KillmailItem slice")
	}

	return o, nil
}

// Count returns the count of all KillmailItem records in the query.
func (q killmailItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count killmail_items rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q killmailItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if killmail_items exists")
	}

	return count > 0, nil
}

// Killmail pointed to by the foreign key.
func (o *KillmailItem) Killmail(mods ...qm.QueryMod) killmailQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.KillmailID),
	}

	queryMods = append(queryMods, mods...)

	query := Killmails(queryMods...)
	queries.SetFrom(query.Query, "`killmails`")

	return query
}

// LoadKillmail allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (killmailItemL) LoadKillmail(ctx context.Context, e boil.ContextExecutor, singular bool, maybeKillmailItem interface{}, mods queries.Applicator) error {
	var slice []*KillmailItem
	var object *KillmailItem

	if singular {
		object = maybeKillmailItem.(*KillmailItem)
	} else {
		slice = *maybeKillmailItem.(*[]*KillmailItem)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &killmailItemR{}
		}
		args = append(args, object.KillmailID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &killmailItemR{}
			}

			for _, a := range args {
				if a == obj.KillmailID {
					continue Outer
				}
			}

			args = append(args, obj.KillmailID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`killmails`), qm.WhereIn(`killmails.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Killmail")
	}

	var resultSlice []*Killmail
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Killmail")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for killmails")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for killmails")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Killmail = foreign
		if foreign.R == nil {
			foreign.R = &killmailR{}
		}
		foreign.R.KillmailItems = append(foreign.R.KillmailItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.KillmailID == foreign.ID {
				local.R.Killmail = foreign
				if foreign.R == nil {
					foreign.R = &killmailR{}
				}
				foreign.R.KillmailItems = append(foreign.R.KillmailItems, local)
				break
			}
		}
	}

	return nil
}

// SetKillmail of the killmailItem to the related item.
// Sets o.R.Killmail to related.
// Adds o to related.R.KillmailItems.
func (o *KillmailItem) SetKillmail(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Killmail) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer(), false); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `killmail_items` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"killmail_id"}),
		strmangle.WhereClause("`", "`", 0, killmailItemPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.KillmailID = related.ID
	if o.R == nil {
		o.R = &killmailItemR{
			Killmail: related,
		}
	} else {
		o.R.Killmail = related
	}

	if related.R == nil {
		related.R = &killmailR{
			KillmailItems: KillmailItemSlice{o},
		}
	} else {
		related.R.KillmailItems = append(related.R.KillmailItems, o)
	}

	return nil
}

// KillmailItems retrieves all the records using an executor.
func KillmailItems(mods ...qm.QueryMod) killmailItemQuery {
	mods = append(mods, qm.From("`killmail_items`"))
	return killmailItemQuery{NewQuery(mods...)}
}

// FindKillmailItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindKillmailItem(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*KillmailItem, error) {
	killmailItemObj := &KillmailItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `killmail_items` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, killmailItemObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from killmail_items")
	}

	return killmailItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *KillmailItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns, ignore bool) error {
	if o == nil {
		return errors.New("boiler: no killmail_items provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(killmailItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	killmailItemInsertCacheMut.RLock()
	cache, cached := killmailItemInsertCache[key]
	killmailItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			killmailItemAllColumns,
			killmailItemColumnsWithDefault,
			killmailItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(killmailItemType, killmailItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(killmailItemType, killmailItemMapping, returnColumns)
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
			cache.query = fmt.Sprintf("%s INTO `killmail_items` (`%s`) %%sVALUES (%s)%%s", insert, strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = fmt.Sprintf("%s INTO `killmail_items` () VALUES ()%s%s", insert)
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `killmail_items` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, killmailItemPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "boiler: unable to insert into killmail_items")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == killmailItemMapping["id"] {
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
		return errors.Wrap(err, "boiler: unable to populate default values for killmail_items")
	}

CacheNoHooks:
	if !cached {
		killmailItemInsertCacheMut.Lock()
		killmailItemInsertCache[key] = cache
		killmailItemInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the KillmailItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *KillmailItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	killmailItemUpdateCacheMut.RLock()
	cache, cached := killmailItemUpdateCache[key]
	killmailItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			killmailItemAllColumns,
			killmailItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update killmail_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `killmail_items` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, killmailItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(killmailItemType, killmailItemMapping, append(wl, killmailItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update killmail_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for killmail_items")
	}

	if !cached {
		killmailItemUpdateCacheMut.Lock()
		killmailItemUpdateCache[key] = cache
		killmailItemUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q killmailItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for killmail_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for killmail_items")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o KillmailItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), killmailItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `killmail_items` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, killmailItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in killmailItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all killmailItem")
	}
	return rowsAff, nil
}

var mySQLKillmailItemUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *KillmailItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no killmail_items provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(killmailItemColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLKillmailItemUniqueColumns, o)

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

	killmailItemUpsertCacheMut.RLock()
	cache, cached := killmailItemUpsertCache[key]
	killmailItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			killmailItemAllColumns,
			killmailItemColumnsWithDefault,
			killmailItemColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			killmailItemAllColumns,
			killmailItemPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("boiler: unable to upsert killmail_items, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "killmail_items", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `killmail_items` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(killmailItemType, killmailItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(killmailItemType, killmailItemMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "boiler: unable to upsert for killmail_items")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == killmailItemMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(killmailItemType, killmailItemMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to retrieve unique values for killmail_items")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for killmail_items")
	}

CacheNoHooks:
	if !cached {
		killmailItemUpsertCacheMut.Lock()
		killmailItemUpsertCache[key] = cache
		killmailItemUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single KillmailItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *KillmailItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no KillmailItem provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), killmailItemPrimaryKeyMapping)
	sql := "DELETE FROM `killmail_items` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from killmail_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for killmail_items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q killmailItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no killmailItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from killmail_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for killmail_items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o KillmailItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), killmailItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `killmail_items` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, killmailItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from killmailItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for killmail_items")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *KillmailItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindKillmailItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *KillmailItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := KillmailItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), killmailItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `killmail_items`.* FROM `killmail_items` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, killmailItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in KillmailItemSlice")
	}

	*o = slice

	return nil
}

// KillmailItemExists checks if the KillmailItem row exists.
func KillmailItemExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `killmail_items` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if killmail_items exists")
	}

	return exists, nil
}
