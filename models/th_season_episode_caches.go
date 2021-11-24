// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// THSeasonEpisodeCach is an object representing the database table.
type THSeasonEpisodeCach struct {
	EpisodeID int64     `boil:"episode_id" json:"episode_id" toml:"episode_id" yaml:"episode_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	SeasonID  int64     `boil:"season_id" json:"season_id" toml:"season_id" yaml:"season_id"`

	R *thSeasonEpisodeCachR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L thSeasonEpisodeCachL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var THSeasonEpisodeCachColumns = struct {
	EpisodeID string
	CreatedAt string
	UpdatedAt string
	SeasonID  string
}{
	EpisodeID: "episode_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	SeasonID:  "season_id",
}

var THSeasonEpisodeCachTableColumns = struct {
	EpisodeID string
	CreatedAt string
	UpdatedAt string
	SeasonID  string
}{
	EpisodeID: "th_season_episode_caches.episode_id",
	CreatedAt: "th_season_episode_caches.created_at",
	UpdatedAt: "th_season_episode_caches.updated_at",
	SeasonID:  "th_season_episode_caches.season_id",
}

// Generated where

var THSeasonEpisodeCachWhere = struct {
	EpisodeID whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	SeasonID  whereHelperint64
}{
	EpisodeID: whereHelperint64{field: "\"th_season_episode_caches\".\"episode_id\""},
	CreatedAt: whereHelpertime_Time{field: "\"th_season_episode_caches\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"th_season_episode_caches\".\"updated_at\""},
	SeasonID:  whereHelperint64{field: "\"th_season_episode_caches\".\"season_id\""},
}

// THSeasonEpisodeCachRels is where relationship names are stored.
var THSeasonEpisodeCachRels = struct {
	Season string
}{
	Season: "Season",
}

// thSeasonEpisodeCachR is where relationships are stored.
type thSeasonEpisodeCachR struct {
	Season *THSeasonCach `boil:"Season" json:"Season" toml:"Season" yaml:"Season"`
}

// NewStruct creates a new relationship struct
func (*thSeasonEpisodeCachR) NewStruct() *thSeasonEpisodeCachR {
	return &thSeasonEpisodeCachR{}
}

// thSeasonEpisodeCachL is where Load methods for each relationship are stored.
type thSeasonEpisodeCachL struct{}

var (
	thSeasonEpisodeCachAllColumns            = []string{"episode_id", "created_at", "updated_at", "season_id"}
	thSeasonEpisodeCachColumnsWithoutDefault = []string{"episode_id", "created_at", "updated_at", "season_id"}
	thSeasonEpisodeCachColumnsWithDefault    = []string{}
	thSeasonEpisodeCachPrimaryKeyColumns     = []string{"episode_id"}
)

type (
	// THSeasonEpisodeCachSlice is an alias for a slice of pointers to THSeasonEpisodeCach.
	// This should almost always be used instead of []THSeasonEpisodeCach.
	THSeasonEpisodeCachSlice []*THSeasonEpisodeCach
	// THSeasonEpisodeCachHook is the signature for custom THSeasonEpisodeCach hook methods
	THSeasonEpisodeCachHook func(context.Context, boil.ContextExecutor, *THSeasonEpisodeCach) error

	thSeasonEpisodeCachQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	thSeasonEpisodeCachType                 = reflect.TypeOf(&THSeasonEpisodeCach{})
	thSeasonEpisodeCachMapping              = queries.MakeStructMapping(thSeasonEpisodeCachType)
	thSeasonEpisodeCachPrimaryKeyMapping, _ = queries.BindMapping(thSeasonEpisodeCachType, thSeasonEpisodeCachMapping, thSeasonEpisodeCachPrimaryKeyColumns)
	thSeasonEpisodeCachInsertCacheMut       sync.RWMutex
	thSeasonEpisodeCachInsertCache          = make(map[string]insertCache)
	thSeasonEpisodeCachUpdateCacheMut       sync.RWMutex
	thSeasonEpisodeCachUpdateCache          = make(map[string]updateCache)
	thSeasonEpisodeCachUpsertCacheMut       sync.RWMutex
	thSeasonEpisodeCachUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var thSeasonEpisodeCachBeforeInsertHooks []THSeasonEpisodeCachHook
var thSeasonEpisodeCachBeforeUpdateHooks []THSeasonEpisodeCachHook
var thSeasonEpisodeCachBeforeDeleteHooks []THSeasonEpisodeCachHook
var thSeasonEpisodeCachBeforeUpsertHooks []THSeasonEpisodeCachHook

var thSeasonEpisodeCachAfterInsertHooks []THSeasonEpisodeCachHook
var thSeasonEpisodeCachAfterSelectHooks []THSeasonEpisodeCachHook
var thSeasonEpisodeCachAfterUpdateHooks []THSeasonEpisodeCachHook
var thSeasonEpisodeCachAfterDeleteHooks []THSeasonEpisodeCachHook
var thSeasonEpisodeCachAfterUpsertHooks []THSeasonEpisodeCachHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *THSeasonEpisodeCach) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonEpisodeCachBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *THSeasonEpisodeCach) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonEpisodeCachBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *THSeasonEpisodeCach) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonEpisodeCachBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *THSeasonEpisodeCach) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonEpisodeCachBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *THSeasonEpisodeCach) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonEpisodeCachAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *THSeasonEpisodeCach) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonEpisodeCachAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *THSeasonEpisodeCach) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonEpisodeCachAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *THSeasonEpisodeCach) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonEpisodeCachAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *THSeasonEpisodeCach) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonEpisodeCachAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTHSeasonEpisodeCachHook registers your hook function for all future operations.
func AddTHSeasonEpisodeCachHook(hookPoint boil.HookPoint, thSeasonEpisodeCachHook THSeasonEpisodeCachHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		thSeasonEpisodeCachBeforeInsertHooks = append(thSeasonEpisodeCachBeforeInsertHooks, thSeasonEpisodeCachHook)
	case boil.BeforeUpdateHook:
		thSeasonEpisodeCachBeforeUpdateHooks = append(thSeasonEpisodeCachBeforeUpdateHooks, thSeasonEpisodeCachHook)
	case boil.BeforeDeleteHook:
		thSeasonEpisodeCachBeforeDeleteHooks = append(thSeasonEpisodeCachBeforeDeleteHooks, thSeasonEpisodeCachHook)
	case boil.BeforeUpsertHook:
		thSeasonEpisodeCachBeforeUpsertHooks = append(thSeasonEpisodeCachBeforeUpsertHooks, thSeasonEpisodeCachHook)
	case boil.AfterInsertHook:
		thSeasonEpisodeCachAfterInsertHooks = append(thSeasonEpisodeCachAfterInsertHooks, thSeasonEpisodeCachHook)
	case boil.AfterSelectHook:
		thSeasonEpisodeCachAfterSelectHooks = append(thSeasonEpisodeCachAfterSelectHooks, thSeasonEpisodeCachHook)
	case boil.AfterUpdateHook:
		thSeasonEpisodeCachAfterUpdateHooks = append(thSeasonEpisodeCachAfterUpdateHooks, thSeasonEpisodeCachHook)
	case boil.AfterDeleteHook:
		thSeasonEpisodeCachAfterDeleteHooks = append(thSeasonEpisodeCachAfterDeleteHooks, thSeasonEpisodeCachHook)
	case boil.AfterUpsertHook:
		thSeasonEpisodeCachAfterUpsertHooks = append(thSeasonEpisodeCachAfterUpsertHooks, thSeasonEpisodeCachHook)
	}
}

// One returns a single thSeasonEpisodeCach record from the query.
func (q thSeasonEpisodeCachQuery) One(ctx context.Context, exec boil.ContextExecutor) (*THSeasonEpisodeCach, error) {
	o := &THSeasonEpisodeCach{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for th_season_episode_caches")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all THSeasonEpisodeCach records from the query.
func (q thSeasonEpisodeCachQuery) All(ctx context.Context, exec boil.ContextExecutor) (THSeasonEpisodeCachSlice, error) {
	var o []*THSeasonEpisodeCach

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to THSeasonEpisodeCach slice")
	}

	if len(thSeasonEpisodeCachAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all THSeasonEpisodeCach records in the query.
func (q thSeasonEpisodeCachQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count th_season_episode_caches rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q thSeasonEpisodeCachQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if th_season_episode_caches exists")
	}

	return count > 0, nil
}

// Season pointed to by the foreign key.
func (o *THSeasonEpisodeCach) Season(mods ...qm.QueryMod) thSeasonCachQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"season_id\" = ?", o.SeasonID),
	}

	queryMods = append(queryMods, mods...)

	query := THSeasonCaches(queryMods...)
	queries.SetFrom(query.Query, "\"th_season_caches\"")

	return query
}

// LoadSeason allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (thSeasonEpisodeCachL) LoadSeason(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTHSeasonEpisodeCach interface{}, mods queries.Applicator) error {
	var slice []*THSeasonEpisodeCach
	var object *THSeasonEpisodeCach

	if singular {
		object = maybeTHSeasonEpisodeCach.(*THSeasonEpisodeCach)
	} else {
		slice = *maybeTHSeasonEpisodeCach.(*[]*THSeasonEpisodeCach)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &thSeasonEpisodeCachR{}
		}
		args = append(args, object.SeasonID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &thSeasonEpisodeCachR{}
			}

			for _, a := range args {
				if a == obj.SeasonID {
					continue Outer
				}
			}

			args = append(args, obj.SeasonID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`th_season_caches`),
		qm.WhereIn(`th_season_caches.season_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load THSeasonCach")
	}

	var resultSlice []*THSeasonCach
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice THSeasonCach")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for th_season_caches")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for th_season_caches")
	}

	if len(thSeasonEpisodeCachAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Season = foreign
		if foreign.R == nil {
			foreign.R = &thSeasonCachR{}
		}
		foreign.R.SeasonTHSeasonEpisodeCaches = append(foreign.R.SeasonTHSeasonEpisodeCaches, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SeasonID == foreign.SeasonID {
				local.R.Season = foreign
				if foreign.R == nil {
					foreign.R = &thSeasonCachR{}
				}
				foreign.R.SeasonTHSeasonEpisodeCaches = append(foreign.R.SeasonTHSeasonEpisodeCaches, local)
				break
			}
		}
	}

	return nil
}

// SetSeason of the thSeasonEpisodeCach to the related item.
// Sets o.R.Season to related.
// Adds o to related.R.SeasonTHSeasonEpisodeCaches.
func (o *THSeasonEpisodeCach) SetSeason(ctx context.Context, exec boil.ContextExecutor, insert bool, related *THSeasonCach) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"th_season_episode_caches\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"season_id"}),
		strmangle.WhereClause("\"", "\"", 2, thSeasonEpisodeCachPrimaryKeyColumns),
	)
	values := []interface{}{related.SeasonID, o.EpisodeID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SeasonID = related.SeasonID
	if o.R == nil {
		o.R = &thSeasonEpisodeCachR{
			Season: related,
		}
	} else {
		o.R.Season = related
	}

	if related.R == nil {
		related.R = &thSeasonCachR{
			SeasonTHSeasonEpisodeCaches: THSeasonEpisodeCachSlice{o},
		}
	} else {
		related.R.SeasonTHSeasonEpisodeCaches = append(related.R.SeasonTHSeasonEpisodeCaches, o)
	}

	return nil
}

// THSeasonEpisodeCaches retrieves all the records using an executor.
func THSeasonEpisodeCaches(mods ...qm.QueryMod) thSeasonEpisodeCachQuery {
	mods = append(mods, qm.From("\"th_season_episode_caches\""))
	return thSeasonEpisodeCachQuery{NewQuery(mods...)}
}

// FindTHSeasonEpisodeCach retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTHSeasonEpisodeCach(ctx context.Context, exec boil.ContextExecutor, episodeID int64, selectCols ...string) (*THSeasonEpisodeCach, error) {
	thSeasonEpisodeCachObj := &THSeasonEpisodeCach{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"th_season_episode_caches\" where \"episode_id\"=$1", sel,
	)

	q := queries.Raw(query, episodeID)

	err := q.Bind(ctx, exec, thSeasonEpisodeCachObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from th_season_episode_caches")
	}

	if err = thSeasonEpisodeCachObj.doAfterSelectHooks(ctx, exec); err != nil {
		return thSeasonEpisodeCachObj, err
	}

	return thSeasonEpisodeCachObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *THSeasonEpisodeCach) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no th_season_episode_caches provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(thSeasonEpisodeCachColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	thSeasonEpisodeCachInsertCacheMut.RLock()
	cache, cached := thSeasonEpisodeCachInsertCache[key]
	thSeasonEpisodeCachInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			thSeasonEpisodeCachAllColumns,
			thSeasonEpisodeCachColumnsWithDefault,
			thSeasonEpisodeCachColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(thSeasonEpisodeCachType, thSeasonEpisodeCachMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(thSeasonEpisodeCachType, thSeasonEpisodeCachMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"th_season_episode_caches\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"th_season_episode_caches\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into th_season_episode_caches")
	}

	if !cached {
		thSeasonEpisodeCachInsertCacheMut.Lock()
		thSeasonEpisodeCachInsertCache[key] = cache
		thSeasonEpisodeCachInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the THSeasonEpisodeCach.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *THSeasonEpisodeCach) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	thSeasonEpisodeCachUpdateCacheMut.RLock()
	cache, cached := thSeasonEpisodeCachUpdateCache[key]
	thSeasonEpisodeCachUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			thSeasonEpisodeCachAllColumns,
			thSeasonEpisodeCachPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update th_season_episode_caches, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"th_season_episode_caches\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, thSeasonEpisodeCachPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(thSeasonEpisodeCachType, thSeasonEpisodeCachMapping, append(wl, thSeasonEpisodeCachPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update th_season_episode_caches row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for th_season_episode_caches")
	}

	if !cached {
		thSeasonEpisodeCachUpdateCacheMut.Lock()
		thSeasonEpisodeCachUpdateCache[key] = cache
		thSeasonEpisodeCachUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q thSeasonEpisodeCachQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for th_season_episode_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for th_season_episode_caches")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o THSeasonEpisodeCachSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeasonEpisodeCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"th_season_episode_caches\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, thSeasonEpisodeCachPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in thSeasonEpisodeCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all thSeasonEpisodeCach")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *THSeasonEpisodeCach) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no th_season_episode_caches provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(thSeasonEpisodeCachColumnsWithDefault, o)

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

	thSeasonEpisodeCachUpsertCacheMut.RLock()
	cache, cached := thSeasonEpisodeCachUpsertCache[key]
	thSeasonEpisodeCachUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			thSeasonEpisodeCachAllColumns,
			thSeasonEpisodeCachColumnsWithDefault,
			thSeasonEpisodeCachColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			thSeasonEpisodeCachAllColumns,
			thSeasonEpisodeCachPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert th_season_episode_caches, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(thSeasonEpisodeCachPrimaryKeyColumns))
			copy(conflict, thSeasonEpisodeCachPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"th_season_episode_caches\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(thSeasonEpisodeCachType, thSeasonEpisodeCachMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(thSeasonEpisodeCachType, thSeasonEpisodeCachMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert th_season_episode_caches")
	}

	if !cached {
		thSeasonEpisodeCachUpsertCacheMut.Lock()
		thSeasonEpisodeCachUpsertCache[key] = cache
		thSeasonEpisodeCachUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single THSeasonEpisodeCach record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *THSeasonEpisodeCach) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no THSeasonEpisodeCach provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), thSeasonEpisodeCachPrimaryKeyMapping)
	sql := "DELETE FROM \"th_season_episode_caches\" WHERE \"episode_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from th_season_episode_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for th_season_episode_caches")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q thSeasonEpisodeCachQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no thSeasonEpisodeCachQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from th_season_episode_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for th_season_episode_caches")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o THSeasonEpisodeCachSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(thSeasonEpisodeCachBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeasonEpisodeCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"th_season_episode_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thSeasonEpisodeCachPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from thSeasonEpisodeCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for th_season_episode_caches")
	}

	if len(thSeasonEpisodeCachAfterDeleteHooks) != 0 {
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
func (o *THSeasonEpisodeCach) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTHSeasonEpisodeCach(ctx, exec, o.EpisodeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *THSeasonEpisodeCachSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := THSeasonEpisodeCachSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeasonEpisodeCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"th_season_episode_caches\".* FROM \"th_season_episode_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thSeasonEpisodeCachPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in THSeasonEpisodeCachSlice")
	}

	*o = slice

	return nil
}

// THSeasonEpisodeCachExists checks if the THSeasonEpisodeCach row exists.
func THSeasonEpisodeCachExists(ctx context.Context, exec boil.ContextExecutor, episodeID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"th_season_episode_caches\" where \"episode_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, episodeID)
	}
	row := exec.QueryRowContext(ctx, sql, episodeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if th_season_episode_caches exists")
	}

	return exists, nil
}
