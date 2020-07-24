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

// KillmailAttacker is an object representing the database table.
type KillmailAttacker struct {
	ID             uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	KillmailID     uint64      `boil:"killmail_id" json:"killmailID" toml:"killmailID" yaml:"killmailID"`
	AllianceID     null.Uint64 `boil:"alliance_id" json:"allianceID,omitempty" toml:"allianceID" yaml:"allianceID,omitempty"`
	CharacterID    null.Uint64 `boil:"character_id" json:"characterID,omitempty" toml:"characterID" yaml:"characterID,omitempty"`
	CorporationID  null.Uint64 `boil:"corporation_id" json:"corporationID,omitempty" toml:"corporationID" yaml:"corporationID,omitempty"`
	FactionID      null.Uint64 `boil:"faction_id" json:"factionID,omitempty" toml:"factionID" yaml:"factionID,omitempty"`
	DamageDone     uint64      `boil:"damage_done" json:"damageDone" toml:"damageDone" yaml:"damageDone"`
	FinalBlow      bool        `boil:"final_blow" json:"finalBlow" toml:"finalBlow" yaml:"finalBlow"`
	SecurityStatus float64     `boil:"security_status" json:"securityStatus" toml:"securityStatus" yaml:"securityStatus"`
	ShipTypeID     null.Uint64 `boil:"ship_type_id" json:"shipTypeID,omitempty" toml:"shipTypeID" yaml:"shipTypeID,omitempty"`
	WeaponTypeID   null.Uint64 `boil:"weapon_type_id" json:"weaponTypeID,omitempty" toml:"weaponTypeID" yaml:"weaponTypeID,omitempty"`
	CreatedAt      time.Time   `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt      time.Time   `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *killmailAttackerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L killmailAttackerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var KillmailAttackerColumns = struct {
	ID             string
	KillmailID     string
	AllianceID     string
	CharacterID    string
	CorporationID  string
	FactionID      string
	DamageDone     string
	FinalBlow      string
	SecurityStatus string
	ShipTypeID     string
	WeaponTypeID   string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	KillmailID:     "killmail_id",
	AllianceID:     "alliance_id",
	CharacterID:    "character_id",
	CorporationID:  "corporation_id",
	FactionID:      "faction_id",
	DamageDone:     "damage_done",
	FinalBlow:      "final_blow",
	SecurityStatus: "security_status",
	ShipTypeID:     "ship_type_id",
	WeaponTypeID:   "weapon_type_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// Generated where

var KillmailAttackerWhere = struct {
	ID             whereHelperuint64
	KillmailID     whereHelperuint64
	AllianceID     whereHelpernull_Uint64
	CharacterID    whereHelpernull_Uint64
	CorporationID  whereHelpernull_Uint64
	FactionID      whereHelpernull_Uint64
	DamageDone     whereHelperuint64
	FinalBlow      whereHelperbool
	SecurityStatus whereHelperfloat64
	ShipTypeID     whereHelpernull_Uint64
	WeaponTypeID   whereHelpernull_Uint64
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
}{
	ID:             whereHelperuint64{field: "`killmail_attackers`.`id`"},
	KillmailID:     whereHelperuint64{field: "`killmail_attackers`.`killmail_id`"},
	AllianceID:     whereHelpernull_Uint64{field: "`killmail_attackers`.`alliance_id`"},
	CharacterID:    whereHelpernull_Uint64{field: "`killmail_attackers`.`character_id`"},
	CorporationID:  whereHelpernull_Uint64{field: "`killmail_attackers`.`corporation_id`"},
	FactionID:      whereHelpernull_Uint64{field: "`killmail_attackers`.`faction_id`"},
	DamageDone:     whereHelperuint64{field: "`killmail_attackers`.`damage_done`"},
	FinalBlow:      whereHelperbool{field: "`killmail_attackers`.`final_blow`"},
	SecurityStatus: whereHelperfloat64{field: "`killmail_attackers`.`security_status`"},
	ShipTypeID:     whereHelpernull_Uint64{field: "`killmail_attackers`.`ship_type_id`"},
	WeaponTypeID:   whereHelpernull_Uint64{field: "`killmail_attackers`.`weapon_type_id`"},
	CreatedAt:      whereHelpertime_Time{field: "`killmail_attackers`.`created_at`"},
	UpdatedAt:      whereHelpertime_Time{field: "`killmail_attackers`.`updated_at`"},
}

// KillmailAttackerRels is where relationship names are stored.
var KillmailAttackerRels = struct {
	Killmail string
}{
	Killmail: "Killmail",
}

// killmailAttackerR is where relationships are stored.
type killmailAttackerR struct {
	Killmail *Killmail
}

// NewStruct creates a new relationship struct
func (*killmailAttackerR) NewStruct() *killmailAttackerR {
	return &killmailAttackerR{}
}

// killmailAttackerL is where Load methods for each relationship are stored.
type killmailAttackerL struct{}

var (
	killmailAttackerAllColumns            = []string{"id", "killmail_id", "alliance_id", "character_id", "corporation_id", "faction_id", "damage_done", "final_blow", "security_status", "ship_type_id", "weapon_type_id", "created_at", "updated_at"}
	killmailAttackerColumnsWithoutDefault = []string{"killmail_id", "alliance_id", "character_id", "corporation_id", "faction_id", "damage_done", "security_status", "ship_type_id", "weapon_type_id", "created_at", "updated_at"}
	killmailAttackerColumnsWithDefault    = []string{"id", "final_blow"}
	killmailAttackerPrimaryKeyColumns     = []string{"id"}
)

type (
	// KillmailAttackerSlice is an alias for a slice of pointers to KillmailAttacker.
	// This should generally be used opposed to []KillmailAttacker.
	KillmailAttackerSlice []*KillmailAttacker

	killmailAttackerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	killmailAttackerType                 = reflect.TypeOf(&KillmailAttacker{})
	killmailAttackerMapping              = queries.MakeStructMapping(killmailAttackerType)
	killmailAttackerPrimaryKeyMapping, _ = queries.BindMapping(killmailAttackerType, killmailAttackerMapping, killmailAttackerPrimaryKeyColumns)
	killmailAttackerInsertCacheMut       sync.RWMutex
	killmailAttackerInsertCache          = make(map[string]insertCache)
	killmailAttackerUpdateCacheMut       sync.RWMutex
	killmailAttackerUpdateCache          = make(map[string]updateCache)
	killmailAttackerUpsertCacheMut       sync.RWMutex
	killmailAttackerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single killmailAttacker record from the query.
func (q killmailAttackerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*KillmailAttacker, error) {
	o := &KillmailAttacker{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for killmail_attackers")
	}

	return o, nil
}

// All returns all KillmailAttacker records from the query.
func (q killmailAttackerQuery) All(ctx context.Context, exec boil.ContextExecutor) (KillmailAttackerSlice, error) {
	var o []*KillmailAttacker

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to KillmailAttacker slice")
	}

	return o, nil
}

// Count returns the count of all KillmailAttacker records in the query.
func (q killmailAttackerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count killmail_attackers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q killmailAttackerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if killmail_attackers exists")
	}

	return count > 0, nil
}

// Killmail pointed to by the foreign key.
func (o *KillmailAttacker) Killmail(mods ...qm.QueryMod) killmailQuery {
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
func (killmailAttackerL) LoadKillmail(ctx context.Context, e boil.ContextExecutor, singular bool, maybeKillmailAttacker interface{}, mods queries.Applicator) error {
	var slice []*KillmailAttacker
	var object *KillmailAttacker

	if singular {
		object = maybeKillmailAttacker.(*KillmailAttacker)
	} else {
		slice = *maybeKillmailAttacker.(*[]*KillmailAttacker)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &killmailAttackerR{}
		}
		args = append(args, object.KillmailID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &killmailAttackerR{}
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
		foreign.R.KillmailAttackers = append(foreign.R.KillmailAttackers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.KillmailID == foreign.ID {
				local.R.Killmail = foreign
				if foreign.R == nil {
					foreign.R = &killmailR{}
				}
				foreign.R.KillmailAttackers = append(foreign.R.KillmailAttackers, local)
				break
			}
		}
	}

	return nil
}

// SetKillmail of the killmailAttacker to the related item.
// Sets o.R.Killmail to related.
// Adds o to related.R.KillmailAttackers.
func (o *KillmailAttacker) SetKillmail(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Killmail) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer(), false); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `killmail_attackers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"killmail_id"}),
		strmangle.WhereClause("`", "`", 0, killmailAttackerPrimaryKeyColumns),
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
		o.R = &killmailAttackerR{
			Killmail: related,
		}
	} else {
		o.R.Killmail = related
	}

	if related.R == nil {
		related.R = &killmailR{
			KillmailAttackers: KillmailAttackerSlice{o},
		}
	} else {
		related.R.KillmailAttackers = append(related.R.KillmailAttackers, o)
	}

	return nil
}

// KillmailAttackers retrieves all the records using an executor.
func KillmailAttackers(mods ...qm.QueryMod) killmailAttackerQuery {
	mods = append(mods, qm.From("`killmail_attackers`"))
	return killmailAttackerQuery{NewQuery(mods...)}
}

// FindKillmailAttacker retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindKillmailAttacker(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*KillmailAttacker, error) {
	killmailAttackerObj := &KillmailAttacker{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `killmail_attackers` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, killmailAttackerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from killmail_attackers")
	}

	return killmailAttackerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *KillmailAttacker) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns, ignore bool) error {
	if o == nil {
		return errors.New("boiler: no killmail_attackers provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(killmailAttackerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	killmailAttackerInsertCacheMut.RLock()
	cache, cached := killmailAttackerInsertCache[key]
	killmailAttackerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			killmailAttackerAllColumns,
			killmailAttackerColumnsWithDefault,
			killmailAttackerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(killmailAttackerType, killmailAttackerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(killmailAttackerType, killmailAttackerMapping, returnColumns)
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
			cache.query = fmt.Sprintf("%s INTO `killmail_attackers` (`%s`) %%sVALUES (%s)%%s", insert, strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = fmt.Sprintf("%s INTO `killmail_attackers` () VALUES ()%s%s", insert)
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `killmail_attackers` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, killmailAttackerPrimaryKeyColumns))
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
		return errors.Wrap(err, "boiler: unable to insert into killmail_attackers")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == killmailAttackerMapping["id"] {
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
		return errors.Wrap(err, "boiler: unable to populate default values for killmail_attackers")
	}

CacheNoHooks:
	if !cached {
		killmailAttackerInsertCacheMut.Lock()
		killmailAttackerInsertCache[key] = cache
		killmailAttackerInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the KillmailAttacker.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *KillmailAttacker) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	killmailAttackerUpdateCacheMut.RLock()
	cache, cached := killmailAttackerUpdateCache[key]
	killmailAttackerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			killmailAttackerAllColumns,
			killmailAttackerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update killmail_attackers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `killmail_attackers` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, killmailAttackerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(killmailAttackerType, killmailAttackerMapping, append(wl, killmailAttackerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update killmail_attackers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for killmail_attackers")
	}

	if !cached {
		killmailAttackerUpdateCacheMut.Lock()
		killmailAttackerUpdateCache[key] = cache
		killmailAttackerUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q killmailAttackerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for killmail_attackers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for killmail_attackers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o KillmailAttackerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), killmailAttackerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `killmail_attackers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, killmailAttackerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in killmailAttacker slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all killmailAttacker")
	}
	return rowsAff, nil
}

var mySQLKillmailAttackerUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *KillmailAttacker) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no killmail_attackers provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(killmailAttackerColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLKillmailAttackerUniqueColumns, o)

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

	killmailAttackerUpsertCacheMut.RLock()
	cache, cached := killmailAttackerUpsertCache[key]
	killmailAttackerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			killmailAttackerAllColumns,
			killmailAttackerColumnsWithDefault,
			killmailAttackerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			killmailAttackerAllColumns,
			killmailAttackerPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("boiler: unable to upsert killmail_attackers, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "killmail_attackers", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `killmail_attackers` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(killmailAttackerType, killmailAttackerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(killmailAttackerType, killmailAttackerMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert for killmail_attackers")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == killmailAttackerMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(killmailAttackerType, killmailAttackerMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to retrieve unique values for killmail_attackers")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for killmail_attackers")
	}

CacheNoHooks:
	if !cached {
		killmailAttackerUpsertCacheMut.Lock()
		killmailAttackerUpsertCache[key] = cache
		killmailAttackerUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single KillmailAttacker record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *KillmailAttacker) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no KillmailAttacker provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), killmailAttackerPrimaryKeyMapping)
	sql := "DELETE FROM `killmail_attackers` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from killmail_attackers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for killmail_attackers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q killmailAttackerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no killmailAttackerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from killmail_attackers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for killmail_attackers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o KillmailAttackerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), killmailAttackerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `killmail_attackers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, killmailAttackerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from killmailAttacker slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for killmail_attackers")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *KillmailAttacker) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindKillmailAttacker(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *KillmailAttackerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := KillmailAttackerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), killmailAttackerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `killmail_attackers`.* FROM `killmail_attackers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, killmailAttackerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in KillmailAttackerSlice")
	}

	*o = slice

	return nil
}

// KillmailAttackerExists checks if the KillmailAttacker row exists.
func KillmailAttackerExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `killmail_attackers` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if killmail_attackers exists")
	}

	return exists, nil
}
