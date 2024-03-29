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

func testBookmarkSearchDocsizesUpsert(t *testing.T) {
	t.Parallel()
	if len(bookmarkSearchDocsizeAllColumns) == len(bookmarkSearchDocsizePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, &o, bookmarkSearchDocsizeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookmarkSearchDocsize: %s", err)
	}

	count, err := BookmarkSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bookmarkSearchDocsizeDBTypes, false, bookmarkSearchDocsizePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookmarkSearchDocsize: %s", err)
	}

	count, err = BookmarkSearchDocsizes().Count(ctx, tx)
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

func testBookmarkSearchDocsizes(t *testing.T) {
	t.Parallel()

	query := BookmarkSearchDocsizes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBookmarkSearchDocsizesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
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

	count, err := BookmarkSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarkSearchDocsizesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BookmarkSearchDocsizes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookmarkSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarkSearchDocsizesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookmarkSearchDocsizeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookmarkSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarkSearchDocsizesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BookmarkSearchDocsizeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BookmarkSearchDocsize exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BookmarkSearchDocsizeExists to return true, but got false.")
	}
}

func testBookmarkSearchDocsizesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bookmarkSearchDocsizeFound, err := FindBookmarkSearchDocsize(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if bookmarkSearchDocsizeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBookmarkSearchDocsizesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BookmarkSearchDocsizes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBookmarkSearchDocsizesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BookmarkSearchDocsizes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBookmarkSearchDocsizesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bookmarkSearchDocsizeOne := &BookmarkSearchDocsize{}
	bookmarkSearchDocsizeTwo := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, bookmarkSearchDocsizeOne, bookmarkSearchDocsizeDBTypes, false, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}
	if err = randomize.Struct(seed, bookmarkSearchDocsizeTwo, bookmarkSearchDocsizeDBTypes, false, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookmarkSearchDocsizeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookmarkSearchDocsizeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookmarkSearchDocsizes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBookmarkSearchDocsizesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bookmarkSearchDocsizeOne := &BookmarkSearchDocsize{}
	bookmarkSearchDocsizeTwo := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, bookmarkSearchDocsizeOne, bookmarkSearchDocsizeDBTypes, false, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}
	if err = randomize.Struct(seed, bookmarkSearchDocsizeTwo, bookmarkSearchDocsizeDBTypes, false, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookmarkSearchDocsizeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookmarkSearchDocsizeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookmarkSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bookmarkSearchDocsizeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkSearchDocsize) error {
	*o = BookmarkSearchDocsize{}
	return nil
}

func bookmarkSearchDocsizeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkSearchDocsize) error {
	*o = BookmarkSearchDocsize{}
	return nil
}

func bookmarkSearchDocsizeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkSearchDocsize) error {
	*o = BookmarkSearchDocsize{}
	return nil
}

func bookmarkSearchDocsizeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkSearchDocsize) error {
	*o = BookmarkSearchDocsize{}
	return nil
}

func bookmarkSearchDocsizeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkSearchDocsize) error {
	*o = BookmarkSearchDocsize{}
	return nil
}

func bookmarkSearchDocsizeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkSearchDocsize) error {
	*o = BookmarkSearchDocsize{}
	return nil
}

func bookmarkSearchDocsizeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkSearchDocsize) error {
	*o = BookmarkSearchDocsize{}
	return nil
}

func bookmarkSearchDocsizeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkSearchDocsize) error {
	*o = BookmarkSearchDocsize{}
	return nil
}

func bookmarkSearchDocsizeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkSearchDocsize) error {
	*o = BookmarkSearchDocsize{}
	return nil
}

func testBookmarkSearchDocsizesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BookmarkSearchDocsize{}
	o := &BookmarkSearchDocsize{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize object: %s", err)
	}

	AddBookmarkSearchDocsizeHook(boil.BeforeInsertHook, bookmarkSearchDocsizeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bookmarkSearchDocsizeBeforeInsertHooks = []BookmarkSearchDocsizeHook{}

	AddBookmarkSearchDocsizeHook(boil.AfterInsertHook, bookmarkSearchDocsizeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bookmarkSearchDocsizeAfterInsertHooks = []BookmarkSearchDocsizeHook{}

	AddBookmarkSearchDocsizeHook(boil.AfterSelectHook, bookmarkSearchDocsizeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bookmarkSearchDocsizeAfterSelectHooks = []BookmarkSearchDocsizeHook{}

	AddBookmarkSearchDocsizeHook(boil.BeforeUpdateHook, bookmarkSearchDocsizeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bookmarkSearchDocsizeBeforeUpdateHooks = []BookmarkSearchDocsizeHook{}

	AddBookmarkSearchDocsizeHook(boil.AfterUpdateHook, bookmarkSearchDocsizeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bookmarkSearchDocsizeAfterUpdateHooks = []BookmarkSearchDocsizeHook{}

	AddBookmarkSearchDocsizeHook(boil.BeforeDeleteHook, bookmarkSearchDocsizeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bookmarkSearchDocsizeBeforeDeleteHooks = []BookmarkSearchDocsizeHook{}

	AddBookmarkSearchDocsizeHook(boil.AfterDeleteHook, bookmarkSearchDocsizeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bookmarkSearchDocsizeAfterDeleteHooks = []BookmarkSearchDocsizeHook{}

	AddBookmarkSearchDocsizeHook(boil.BeforeUpsertHook, bookmarkSearchDocsizeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bookmarkSearchDocsizeBeforeUpsertHooks = []BookmarkSearchDocsizeHook{}

	AddBookmarkSearchDocsizeHook(boil.AfterUpsertHook, bookmarkSearchDocsizeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bookmarkSearchDocsizeAfterUpsertHooks = []BookmarkSearchDocsizeHook{}
}

func testBookmarkSearchDocsizesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookmarkSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookmarkSearchDocsizesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bookmarkSearchDocsizeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BookmarkSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookmarkSearchDocsizesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
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

func testBookmarkSearchDocsizesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookmarkSearchDocsizeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookmarkSearchDocsizesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookmarkSearchDocsizes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bookmarkSearchDocsizeDBTypes = map[string]string{`ID`: `INTEGER`, `SZ`: `BLOB`}
	_                            = bytes.MinRead
)

func testBookmarkSearchDocsizesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bookmarkSearchDocsizePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bookmarkSearchDocsizeAllColumns) == len(bookmarkSearchDocsizePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookmarkSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBookmarkSearchDocsizesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bookmarkSearchDocsizeAllColumns) == len(bookmarkSearchDocsizePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkSearchDocsize{}
	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookmarkSearchDocsizes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookmarkSearchDocsizeDBTypes, true, bookmarkSearchDocsizePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookmarkSearchDocsize struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bookmarkSearchDocsizeAllColumns, bookmarkSearchDocsizePrimaryKeyColumns) {
		fields = bookmarkSearchDocsizeAllColumns
	} else {
		fields = strmangle.SetComplement(
			bookmarkSearchDocsizeAllColumns,
			bookmarkSearchDocsizePrimaryKeyColumns,
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

	slice := BookmarkSearchDocsizeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
