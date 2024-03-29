// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testPublisherSearchDocsizesUpsert(t *testing.T) {
	t.Parallel()
	if len(publisherSearchDocsizeAllColumns) == len(publisherSearchDocsizePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PublisherSearchDocsize{}
	if err = randomize.Struct(seed, &o, publisherSearchDocsizeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PublisherSearchDocsize: %s", err)
	}

	count, err := PublisherSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, publisherSearchDocsizeDBTypes, false, publisherSearchDocsizePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PublisherSearchDocsize: %s", err)
	}

	count, err = PublisherSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPublisherSearchDocsizes(t *testing.T) {
	t.Parallel()

	query := PublisherSearchDocsizes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPublisherSearchDocsizesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
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

	count, err := PublisherSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPublisherSearchDocsizesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PublisherSearchDocsizes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PublisherSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPublisherSearchDocsizesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PublisherSearchDocsizeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PublisherSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPublisherSearchDocsizesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PublisherSearchDocsizeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PublisherSearchDocsize exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PublisherSearchDocsizeExists to return true, but got false.")
	}
}

func testPublisherSearchDocsizesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	publisherSearchDocsizeFound, err := FindPublisherSearchDocsize(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if publisherSearchDocsizeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPublisherSearchDocsizesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PublisherSearchDocsizes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPublisherSearchDocsizesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PublisherSearchDocsizes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPublisherSearchDocsizesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	publisherSearchDocsizeOne := &PublisherSearchDocsize{}
	publisherSearchDocsizeTwo := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, publisherSearchDocsizeOne, publisherSearchDocsizeDBTypes, false, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}
	if err = randomize.Struct(seed, publisherSearchDocsizeTwo, publisherSearchDocsizeDBTypes, false, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = publisherSearchDocsizeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = publisherSearchDocsizeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PublisherSearchDocsizes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPublisherSearchDocsizesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	publisherSearchDocsizeOne := &PublisherSearchDocsize{}
	publisherSearchDocsizeTwo := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, publisherSearchDocsizeOne, publisherSearchDocsizeDBTypes, false, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}
	if err = randomize.Struct(seed, publisherSearchDocsizeTwo, publisherSearchDocsizeDBTypes, false, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = publisherSearchDocsizeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = publisherSearchDocsizeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PublisherSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func publisherSearchDocsizeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PublisherSearchDocsize) error {
	*o = PublisherSearchDocsize{}
	return nil
}

func publisherSearchDocsizeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PublisherSearchDocsize) error {
	*o = PublisherSearchDocsize{}
	return nil
}

func publisherSearchDocsizeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PublisherSearchDocsize) error {
	*o = PublisherSearchDocsize{}
	return nil
}

func publisherSearchDocsizeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PublisherSearchDocsize) error {
	*o = PublisherSearchDocsize{}
	return nil
}

func publisherSearchDocsizeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PublisherSearchDocsize) error {
	*o = PublisherSearchDocsize{}
	return nil
}

func publisherSearchDocsizeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PublisherSearchDocsize) error {
	*o = PublisherSearchDocsize{}
	return nil
}

func publisherSearchDocsizeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PublisherSearchDocsize) error {
	*o = PublisherSearchDocsize{}
	return nil
}

func publisherSearchDocsizeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PublisherSearchDocsize) error {
	*o = PublisherSearchDocsize{}
	return nil
}

func publisherSearchDocsizeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PublisherSearchDocsize) error {
	*o = PublisherSearchDocsize{}
	return nil
}

func testPublisherSearchDocsizesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PublisherSearchDocsize{}
	o := &PublisherSearchDocsize{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize object: %s", err)
	}

	AddPublisherSearchDocsizeHook(boil.BeforeInsertHook, publisherSearchDocsizeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	publisherSearchDocsizeBeforeInsertHooks = []PublisherSearchDocsizeHook{}

	AddPublisherSearchDocsizeHook(boil.AfterInsertHook, publisherSearchDocsizeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	publisherSearchDocsizeAfterInsertHooks = []PublisherSearchDocsizeHook{}

	AddPublisherSearchDocsizeHook(boil.AfterSelectHook, publisherSearchDocsizeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	publisherSearchDocsizeAfterSelectHooks = []PublisherSearchDocsizeHook{}

	AddPublisherSearchDocsizeHook(boil.BeforeUpdateHook, publisherSearchDocsizeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	publisherSearchDocsizeBeforeUpdateHooks = []PublisherSearchDocsizeHook{}

	AddPublisherSearchDocsizeHook(boil.AfterUpdateHook, publisherSearchDocsizeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	publisherSearchDocsizeAfterUpdateHooks = []PublisherSearchDocsizeHook{}

	AddPublisherSearchDocsizeHook(boil.BeforeDeleteHook, publisherSearchDocsizeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	publisherSearchDocsizeBeforeDeleteHooks = []PublisherSearchDocsizeHook{}

	AddPublisherSearchDocsizeHook(boil.AfterDeleteHook, publisherSearchDocsizeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	publisherSearchDocsizeAfterDeleteHooks = []PublisherSearchDocsizeHook{}

	AddPublisherSearchDocsizeHook(boil.BeforeUpsertHook, publisherSearchDocsizeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	publisherSearchDocsizeBeforeUpsertHooks = []PublisherSearchDocsizeHook{}

	AddPublisherSearchDocsizeHook(boil.AfterUpsertHook, publisherSearchDocsizeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	publisherSearchDocsizeAfterUpsertHooks = []PublisherSearchDocsizeHook{}
}

func testPublisherSearchDocsizesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PublisherSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPublisherSearchDocsizesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(publisherSearchDocsizeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PublisherSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPublisherSearchDocsizesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
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

func testPublisherSearchDocsizesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PublisherSearchDocsizeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPublisherSearchDocsizesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PublisherSearchDocsizes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	publisherSearchDocsizeDBTypes = map[string]string{`ID`: `INTEGER`, `SZ`: `BLOB`}
	_                             = bytes.MinRead
)

func testPublisherSearchDocsizesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(publisherSearchDocsizePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(publisherSearchDocsizeAllColumns) == len(publisherSearchDocsizePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PublisherSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPublisherSearchDocsizesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(publisherSearchDocsizeAllColumns) == len(publisherSearchDocsizePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PublisherSearchDocsize{}
	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PublisherSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, publisherSearchDocsizeDBTypes, true, publisherSearchDocsizePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PublisherSearchDocsize struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(publisherSearchDocsizeAllColumns, publisherSearchDocsizePrimaryKeyColumns) {
		fields = publisherSearchDocsizeAllColumns
	} else {
		fields = strmangle.SetComplement(
			publisherSearchDocsizeAllColumns,
			publisherSearchDocsizePrimaryKeyColumns,
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

	slice := PublisherSearchDocsizeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
