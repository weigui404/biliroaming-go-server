// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testAccessKeys(t *testing.T) {
	t.Parallel()

	query := AccessKeys()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAccessKeysDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
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

	count, err := AccessKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccessKeysQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AccessKeys().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AccessKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccessKeysSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AccessKeySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AccessKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccessKeysExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AccessKeyExists(ctx, tx, o.Key)
	if err != nil {
		t.Errorf("Unable to check if AccessKey exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AccessKeyExists to return true, but got false.")
	}
}

func testAccessKeysFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	accessKeyFound, err := FindAccessKey(ctx, tx, o.Key)
	if err != nil {
		t.Error(err)
	}

	if accessKeyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAccessKeysBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AccessKeys().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAccessKeysOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AccessKeys().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAccessKeysAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accessKeyOne := &AccessKey{}
	accessKeyTwo := &AccessKey{}
	if err = randomize.Struct(seed, accessKeyOne, accessKeyDBTypes, false, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}
	if err = randomize.Struct(seed, accessKeyTwo, accessKeyDBTypes, false, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = accessKeyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = accessKeyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AccessKeys().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAccessKeysCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	accessKeyOne := &AccessKey{}
	accessKeyTwo := &AccessKey{}
	if err = randomize.Struct(seed, accessKeyOne, accessKeyDBTypes, false, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}
	if err = randomize.Struct(seed, accessKeyTwo, accessKeyDBTypes, false, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = accessKeyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = accessKeyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AccessKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func accessKeyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AccessKey) error {
	*o = AccessKey{}
	return nil
}

func accessKeyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AccessKey) error {
	*o = AccessKey{}
	return nil
}

func accessKeyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AccessKey) error {
	*o = AccessKey{}
	return nil
}

func accessKeyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AccessKey) error {
	*o = AccessKey{}
	return nil
}

func accessKeyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AccessKey) error {
	*o = AccessKey{}
	return nil
}

func accessKeyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AccessKey) error {
	*o = AccessKey{}
	return nil
}

func accessKeyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AccessKey) error {
	*o = AccessKey{}
	return nil
}

func accessKeyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AccessKey) error {
	*o = AccessKey{}
	return nil
}

func accessKeyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AccessKey) error {
	*o = AccessKey{}
	return nil
}

func testAccessKeysHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AccessKey{}
	o := &AccessKey{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, accessKeyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AccessKey object: %s", err)
	}

	AddAccessKeyHook(boil.BeforeInsertHook, accessKeyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	accessKeyBeforeInsertHooks = []AccessKeyHook{}

	AddAccessKeyHook(boil.AfterInsertHook, accessKeyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	accessKeyAfterInsertHooks = []AccessKeyHook{}

	AddAccessKeyHook(boil.AfterSelectHook, accessKeyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	accessKeyAfterSelectHooks = []AccessKeyHook{}

	AddAccessKeyHook(boil.BeforeUpdateHook, accessKeyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	accessKeyBeforeUpdateHooks = []AccessKeyHook{}

	AddAccessKeyHook(boil.AfterUpdateHook, accessKeyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	accessKeyAfterUpdateHooks = []AccessKeyHook{}

	AddAccessKeyHook(boil.BeforeDeleteHook, accessKeyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	accessKeyBeforeDeleteHooks = []AccessKeyHook{}

	AddAccessKeyHook(boil.AfterDeleteHook, accessKeyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	accessKeyAfterDeleteHooks = []AccessKeyHook{}

	AddAccessKeyHook(boil.BeforeUpsertHook, accessKeyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	accessKeyBeforeUpsertHooks = []AccessKeyHook{}

	AddAccessKeyHook(boil.AfterUpsertHook, accessKeyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	accessKeyAfterUpsertHooks = []AccessKeyHook{}
}

func testAccessKeysInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AccessKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccessKeysInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(accessKeyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AccessKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccessKeyToOneUserUsingUIDUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AccessKey
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, accessKeyDBTypes, false, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UID = foreign.UID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.UIDUser().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UID != foreign.UID {
		t.Errorf("want: %v, got %v", foreign.UID, check.UID)
	}

	slice := AccessKeySlice{&local}
	if err = local.L.LoadUIDUser(ctx, tx, false, (*[]*AccessKey)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UIDUser == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.UIDUser = nil
	if err = local.L.LoadUIDUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UIDUser == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAccessKeyToOneSetOpUserUsingUIDUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AccessKey
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accessKeyDBTypes, false, strmangle.SetComplement(accessKeyPrimaryKeyColumns, accessKeyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUIDUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.UIDUser != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UIDAccessKeys[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UID != x.UID {
			t.Error("foreign key was wrong value", a.UID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UID))
		reflect.Indirect(reflect.ValueOf(&a.UID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UID != x.UID {
			t.Error("foreign key was wrong value", a.UID, x.UID)
		}
	}
}

func testAccessKeysReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
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

func testAccessKeysReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AccessKeySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAccessKeysSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AccessKeys().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	accessKeyDBTypes = map[string]string{`Key`: `character`, `UID`: `bigint`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                = bytes.MinRead
)

func testAccessKeysUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(accessKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(accessKeyAllColumns) == len(accessKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AccessKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAccessKeysSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(accessKeyAllColumns) == len(accessKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AccessKey{}
	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AccessKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, accessKeyDBTypes, true, accessKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(accessKeyAllColumns, accessKeyPrimaryKeyColumns) {
		fields = accessKeyAllColumns
	} else {
		fields = strmangle.SetComplement(
			accessKeyAllColumns,
			accessKeyPrimaryKeyColumns,
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

	slice := AccessKeySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAccessKeysUpsert(t *testing.T) {
	t.Parallel()

	if len(accessKeyAllColumns) == len(accessKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AccessKey{}
	if err = randomize.Struct(seed, &o, accessKeyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AccessKey: %s", err)
	}

	count, err := AccessKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, accessKeyDBTypes, false, accessKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AccessKey struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AccessKey: %s", err)
	}

	count, err = AccessKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
