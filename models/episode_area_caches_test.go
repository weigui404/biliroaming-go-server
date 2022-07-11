// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testEpisodeAreaCaches(t *testing.T) {
	t.Parallel()

	query := EpisodeAreaCaches()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEpisodeAreaCachesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EpisodeAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEpisodeAreaCachesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EpisodeAreaCaches().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EpisodeAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEpisodeAreaCachesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EpisodeAreaCachSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EpisodeAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEpisodeAreaCachesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EpisodeAreaCachExists(ctx, tx, o.EpisodeID)
	if err != nil {
		t.Errorf("Unable to check if EpisodeAreaCach exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EpisodeAreaCachExists to return true, but got false.")
	}
}

func testEpisodeAreaCachesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	episodeAreaCachFound, err := FindEpisodeAreaCach(ctx, tx, o.EpisodeID)
	if err != nil {
		t.Error(err)
	}

	if episodeAreaCachFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEpisodeAreaCachesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EpisodeAreaCaches().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEpisodeAreaCachesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EpisodeAreaCaches().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEpisodeAreaCachesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	episodeAreaCachOne := &EpisodeAreaCach{}
	episodeAreaCachTwo := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, episodeAreaCachOne, episodeAreaCachDBTypes, false, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}
	if err = randomize.Struct(seed, episodeAreaCachTwo, episodeAreaCachDBTypes, false, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = episodeAreaCachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = episodeAreaCachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EpisodeAreaCaches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEpisodeAreaCachesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	episodeAreaCachOne := &EpisodeAreaCach{}
	episodeAreaCachTwo := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, episodeAreaCachOne, episodeAreaCachDBTypes, false, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}
	if err = randomize.Struct(seed, episodeAreaCachTwo, episodeAreaCachDBTypes, false, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = episodeAreaCachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = episodeAreaCachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EpisodeAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func episodeAreaCachBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *EpisodeAreaCach) error {
	*o = EpisodeAreaCach{}
	return nil
}

func episodeAreaCachAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *EpisodeAreaCach) error {
	*o = EpisodeAreaCach{}
	return nil
}

func episodeAreaCachAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *EpisodeAreaCach) error {
	*o = EpisodeAreaCach{}
	return nil
}

func episodeAreaCachBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EpisodeAreaCach) error {
	*o = EpisodeAreaCach{}
	return nil
}

func episodeAreaCachAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EpisodeAreaCach) error {
	*o = EpisodeAreaCach{}
	return nil
}

func episodeAreaCachBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EpisodeAreaCach) error {
	*o = EpisodeAreaCach{}
	return nil
}

func episodeAreaCachAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EpisodeAreaCach) error {
	*o = EpisodeAreaCach{}
	return nil
}

func episodeAreaCachBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EpisodeAreaCach) error {
	*o = EpisodeAreaCach{}
	return nil
}

func episodeAreaCachAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EpisodeAreaCach) error {
	*o = EpisodeAreaCach{}
	return nil
}

func testEpisodeAreaCachesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &EpisodeAreaCach{}
	o := &EpisodeAreaCach{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach object: %s", err)
	}

	AddEpisodeAreaCachHook(boil.BeforeInsertHook, episodeAreaCachBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	episodeAreaCachBeforeInsertHooks = []EpisodeAreaCachHook{}

	AddEpisodeAreaCachHook(boil.AfterInsertHook, episodeAreaCachAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	episodeAreaCachAfterInsertHooks = []EpisodeAreaCachHook{}

	AddEpisodeAreaCachHook(boil.AfterSelectHook, episodeAreaCachAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	episodeAreaCachAfterSelectHooks = []EpisodeAreaCachHook{}

	AddEpisodeAreaCachHook(boil.BeforeUpdateHook, episodeAreaCachBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	episodeAreaCachBeforeUpdateHooks = []EpisodeAreaCachHook{}

	AddEpisodeAreaCachHook(boil.AfterUpdateHook, episodeAreaCachAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	episodeAreaCachAfterUpdateHooks = []EpisodeAreaCachHook{}

	AddEpisodeAreaCachHook(boil.BeforeDeleteHook, episodeAreaCachBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	episodeAreaCachBeforeDeleteHooks = []EpisodeAreaCachHook{}

	AddEpisodeAreaCachHook(boil.AfterDeleteHook, episodeAreaCachAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	episodeAreaCachAfterDeleteHooks = []EpisodeAreaCachHook{}

	AddEpisodeAreaCachHook(boil.BeforeUpsertHook, episodeAreaCachBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	episodeAreaCachBeforeUpsertHooks = []EpisodeAreaCachHook{}

	AddEpisodeAreaCachHook(boil.AfterUpsertHook, episodeAreaCachAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	episodeAreaCachAfterUpsertHooks = []EpisodeAreaCachHook{}
}

func testEpisodeAreaCachesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EpisodeAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEpisodeAreaCachesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(episodeAreaCachColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EpisodeAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEpisodeAreaCachesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEpisodeAreaCachesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EpisodeAreaCachSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEpisodeAreaCachesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EpisodeAreaCaches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	episodeAreaCachDBTypes = map[string]string{`EpisodeID`: `bigint`, `CN`: `boolean`, `HK`: `boolean`, `TW`: `boolean`, `TH`: `boolean`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                      = bytes.MinRead
)

func testEpisodeAreaCachesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(episodeAreaCachPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(episodeAreaCachAllColumns) == len(episodeAreaCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EpisodeAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEpisodeAreaCachesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(episodeAreaCachAllColumns) == len(episodeAreaCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EpisodeAreaCach{}
	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EpisodeAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, episodeAreaCachDBTypes, true, episodeAreaCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(episodeAreaCachAllColumns, episodeAreaCachPrimaryKeyColumns) {
		fields = episodeAreaCachAllColumns
	} else {
		fields = strmangle.SetComplement(
			episodeAreaCachAllColumns,
			episodeAreaCachPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := EpisodeAreaCachSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEpisodeAreaCachesUpsert(t *testing.T) {
	t.Parallel()

	if len(episodeAreaCachAllColumns) == len(episodeAreaCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EpisodeAreaCach{}
	if err = randomize.Struct(seed, &o, episodeAreaCachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EpisodeAreaCach: %s", err)
	}

	count, err := EpisodeAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, episodeAreaCachDBTypes, false, episodeAreaCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EpisodeAreaCach struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EpisodeAreaCach: %s", err)
	}

	count, err = EpisodeAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
