// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// SeasonAreaCach is an object representing the database table.
type SeasonAreaCach struct {
	SeasonID  int64     `boil:"season_id" json:"season_id" toml:"season_id" yaml:"season_id"`
	CN        null.Bool `boil:"cn" json:"cn,omitempty" toml:"cn" yaml:"cn,omitempty"`
	HK        null.Bool `boil:"hk" json:"hk,omitempty" toml:"hk" yaml:"hk,omitempty"`
	TW        null.Bool `boil:"tw" json:"tw,omitempty" toml:"tw" yaml:"tw,omitempty"`
	TH        null.Bool `boil:"th" json:"th,omitempty" toml:"th" yaml:"th,omitempty"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *seasonAreaCachR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L seasonAreaCachL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SeasonAreaCachColumns = struct {
	SeasonID  string
	CN        string
	HK        string
	TW        string
	TH        string
	CreatedAt string
	UpdatedAt string
}{
	SeasonID:  "season_id",
	CN:        "cn",
	HK:        "hk",
	TW:        "tw",
	TH:        "th",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var SeasonAreaCachTableColumns = struct {
	SeasonID  string
	CN        string
	HK        string
	TW        string
	TH        string
	CreatedAt string
	UpdatedAt string
}{
	SeasonID:  "season_area_caches.season_id",
	CN:        "season_area_caches.cn",
	HK:        "season_area_caches.hk",
	TW:        "season_area_caches.tw",
	TH:        "season_area_caches.th",
	CreatedAt: "season_area_caches.created_at",
	UpdatedAt: "season_area_caches.updated_at",
}

// Generated where

var SeasonAreaCachWhere = struct {
	SeasonID  whereHelperint64
	CN        whereHelpernull_Bool
	HK        whereHelpernull_Bool
	TW        whereHelpernull_Bool
	TH        whereHelpernull_Bool
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	SeasonID:  whereHelperint64{field: "\"season_area_caches\".\"season_id\""},
	CN:        whereHelpernull_Bool{field: "\"season_area_caches\".\"cn\""},
	HK:        whereHelpernull_Bool{field: "\"season_area_caches\".\"hk\""},
	TW:        whereHelpernull_Bool{field: "\"season_area_caches\".\"tw\""},
	TH:        whereHelpernull_Bool{field: "\"season_area_caches\".\"th\""},
	CreatedAt: whereHelpertime_Time{field: "\"season_area_caches\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"season_area_caches\".\"updated_at\""},
}

// SeasonAreaCachRels is where relationship names are stored.
var SeasonAreaCachRels = struct {
}{}

// seasonAreaCachR is where relationships are stored.
type seasonAreaCachR struct {
}

// NewStruct creates a new relationship struct
func (*seasonAreaCachR) NewStruct() *seasonAreaCachR {
	return &seasonAreaCachR{}
}

// seasonAreaCachL is where Load methods for each relationship are stored.
type seasonAreaCachL struct{}

var (
	seasonAreaCachAllColumns            = []string{"season_id", "cn", "hk", "tw", "th", "created_at", "updated_at"}
	seasonAreaCachColumnsWithoutDefault = []string{"season_id", "created_at", "updated_at"}
	seasonAreaCachColumnsWithDefault    = []string{"cn", "hk", "tw", "th"}
	seasonAreaCachPrimaryKeyColumns     = []string{"season_id"}
	seasonAreaCachGeneratedColumns      = []string{}
)

type (
	// SeasonAreaCachSlice is an alias for a slice of pointers to SeasonAreaCach.
	// This should almost always be used instead of []SeasonAreaCach.
	SeasonAreaCachSlice []*SeasonAreaCach
	// SeasonAreaCachHook is the signature for custom SeasonAreaCach hook methods
	SeasonAreaCachHook func(context.Context, boil.ContextExecutor, *SeasonAreaCach) error

	seasonAreaCachQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	seasonAreaCachType                 = reflect.TypeOf(&SeasonAreaCach{})
	seasonAreaCachMapping              = queries.MakeStructMapping(seasonAreaCachType)
	seasonAreaCachPrimaryKeyMapping, _ = queries.BindMapping(seasonAreaCachType, seasonAreaCachMapping, seasonAreaCachPrimaryKeyColumns)
	seasonAreaCachInsertCacheMut       sync.RWMutex
	seasonAreaCachInsertCache          = make(map[string]insertCache)
	seasonAreaCachUpdateCacheMut       sync.RWMutex
	seasonAreaCachUpdateCache          = make(map[string]updateCache)
	seasonAreaCachUpsertCacheMut       sync.RWMutex
	seasonAreaCachUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var seasonAreaCachAfterSelectHooks []SeasonAreaCachHook

var seasonAreaCachBeforeInsertHooks []SeasonAreaCachHook
var seasonAreaCachAfterInsertHooks []SeasonAreaCachHook

var seasonAreaCachBeforeUpdateHooks []SeasonAreaCachHook
var seasonAreaCachAfterUpdateHooks []SeasonAreaCachHook

var seasonAreaCachBeforeDeleteHooks []SeasonAreaCachHook
var seasonAreaCachAfterDeleteHooks []SeasonAreaCachHook

var seasonAreaCachBeforeUpsertHooks []SeasonAreaCachHook
var seasonAreaCachAfterUpsertHooks []SeasonAreaCachHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SeasonAreaCach) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range seasonAreaCachAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SeasonAreaCach) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range seasonAreaCachBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SeasonAreaCach) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range seasonAreaCachAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SeasonAreaCach) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range seasonAreaCachBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SeasonAreaCach) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range seasonAreaCachAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SeasonAreaCach) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range seasonAreaCachBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SeasonAreaCach) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range seasonAreaCachAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SeasonAreaCach) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range seasonAreaCachBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SeasonAreaCach) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range seasonAreaCachAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSeasonAreaCachHook registers your hook function for all future operations.
func AddSeasonAreaCachHook(hookPoint boil.HookPoint, seasonAreaCachHook SeasonAreaCachHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		seasonAreaCachAfterSelectHooks = append(seasonAreaCachAfterSelectHooks, seasonAreaCachHook)
	case boil.BeforeInsertHook:
		seasonAreaCachBeforeInsertHooks = append(seasonAreaCachBeforeInsertHooks, seasonAreaCachHook)
	case boil.AfterInsertHook:
		seasonAreaCachAfterInsertHooks = append(seasonAreaCachAfterInsertHooks, seasonAreaCachHook)
	case boil.BeforeUpdateHook:
		seasonAreaCachBeforeUpdateHooks = append(seasonAreaCachBeforeUpdateHooks, seasonAreaCachHook)
	case boil.AfterUpdateHook:
		seasonAreaCachAfterUpdateHooks = append(seasonAreaCachAfterUpdateHooks, seasonAreaCachHook)
	case boil.BeforeDeleteHook:
		seasonAreaCachBeforeDeleteHooks = append(seasonAreaCachBeforeDeleteHooks, seasonAreaCachHook)
	case boil.AfterDeleteHook:
		seasonAreaCachAfterDeleteHooks = append(seasonAreaCachAfterDeleteHooks, seasonAreaCachHook)
	case boil.BeforeUpsertHook:
		seasonAreaCachBeforeUpsertHooks = append(seasonAreaCachBeforeUpsertHooks, seasonAreaCachHook)
	case boil.AfterUpsertHook:
		seasonAreaCachAfterUpsertHooks = append(seasonAreaCachAfterUpsertHooks, seasonAreaCachHook)
	}
}

// One returns a single seasonAreaCach record from the query.
func (q seasonAreaCachQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SeasonAreaCach, error) {
	o := &SeasonAreaCach{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for season_area_caches")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SeasonAreaCach records from the query.
func (q seasonAreaCachQuery) All(ctx context.Context, exec boil.ContextExecutor) (SeasonAreaCachSlice, error) {
	var o []*SeasonAreaCach

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SeasonAreaCach slice")
	}

	if len(seasonAreaCachAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SeasonAreaCach records in the query.
func (q seasonAreaCachQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count season_area_caches rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q seasonAreaCachQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if season_area_caches exists")
	}

	return count > 0, nil
}

// SeasonAreaCaches retrieves all the records using an executor.
func SeasonAreaCaches(mods ...qm.QueryMod) seasonAreaCachQuery {
	mods = append(mods, qm.From("\"season_area_caches\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"season_area_caches\".*"})
	}

	return seasonAreaCachQuery{q}
}

// FindSeasonAreaCach retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSeasonAreaCach(ctx context.Context, exec boil.ContextExecutor, seasonID int64, selectCols ...string) (*SeasonAreaCach, error) {
	seasonAreaCachObj := &SeasonAreaCach{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"season_area_caches\" where \"season_id\"=$1", sel,
	)

	q := queries.Raw(query, seasonID)

	err := q.Bind(ctx, exec, seasonAreaCachObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from season_area_caches")
	}

	if err = seasonAreaCachObj.doAfterSelectHooks(ctx, exec); err != nil {
		return seasonAreaCachObj, err
	}

	return seasonAreaCachObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SeasonAreaCach) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no season_area_caches provided for insertion")
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

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(seasonAreaCachColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	seasonAreaCachInsertCacheMut.RLock()
	cache, cached := seasonAreaCachInsertCache[key]
	seasonAreaCachInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			seasonAreaCachAllColumns,
			seasonAreaCachColumnsWithDefault,
			seasonAreaCachColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(seasonAreaCachType, seasonAreaCachMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(seasonAreaCachType, seasonAreaCachMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"season_area_caches\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"season_area_caches\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into season_area_caches")
	}

	if !cached {
		seasonAreaCachInsertCacheMut.Lock()
		seasonAreaCachInsertCache[key] = cache
		seasonAreaCachInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SeasonAreaCach.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SeasonAreaCach) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	seasonAreaCachUpdateCacheMut.RLock()
	cache, cached := seasonAreaCachUpdateCache[key]
	seasonAreaCachUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			seasonAreaCachAllColumns,
			seasonAreaCachPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update season_area_caches, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"season_area_caches\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, seasonAreaCachPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(seasonAreaCachType, seasonAreaCachMapping, append(wl, seasonAreaCachPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update season_area_caches row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for season_area_caches")
	}

	if !cached {
		seasonAreaCachUpdateCacheMut.Lock()
		seasonAreaCachUpdateCache[key] = cache
		seasonAreaCachUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q seasonAreaCachQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for season_area_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for season_area_caches")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SeasonAreaCachSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), seasonAreaCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"season_area_caches\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, seasonAreaCachPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in seasonAreaCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all seasonAreaCach")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SeasonAreaCach) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no season_area_caches provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(seasonAreaCachColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	seasonAreaCachUpsertCacheMut.RLock()
	cache, cached := seasonAreaCachUpsertCache[key]
	seasonAreaCachUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			seasonAreaCachAllColumns,
			seasonAreaCachColumnsWithDefault,
			seasonAreaCachColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			seasonAreaCachAllColumns,
			seasonAreaCachPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert season_area_caches, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(seasonAreaCachPrimaryKeyColumns))
			copy(conflict, seasonAreaCachPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"season_area_caches\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(seasonAreaCachType, seasonAreaCachMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(seasonAreaCachType, seasonAreaCachMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert season_area_caches")
	}

	if !cached {
		seasonAreaCachUpsertCacheMut.Lock()
		seasonAreaCachUpsertCache[key] = cache
		seasonAreaCachUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SeasonAreaCach record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SeasonAreaCach) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SeasonAreaCach provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), seasonAreaCachPrimaryKeyMapping)
	sql := "DELETE FROM \"season_area_caches\" WHERE \"season_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from season_area_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for season_area_caches")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q seasonAreaCachQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no seasonAreaCachQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from season_area_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for season_area_caches")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SeasonAreaCachSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(seasonAreaCachBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), seasonAreaCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"season_area_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, seasonAreaCachPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from seasonAreaCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for season_area_caches")
	}

	if len(seasonAreaCachAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SeasonAreaCach) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSeasonAreaCach(ctx, exec, o.SeasonID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SeasonAreaCachSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SeasonAreaCachSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), seasonAreaCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"season_area_caches\".* FROM \"season_area_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, seasonAreaCachPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SeasonAreaCachSlice")
	}

	*o = slice

	return nil
}

// SeasonAreaCachExists checks if the SeasonAreaCach row exists.
func SeasonAreaCachExists(ctx context.Context, exec boil.ContextExecutor, seasonID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"season_area_caches\" where \"season_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, seasonID)
	}
	row := exec.QueryRowContext(ctx, sql, seasonID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if season_area_caches exists")
	}

	return exists, nil
}
