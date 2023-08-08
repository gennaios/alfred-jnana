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

func testBookmarksUpsert(t *testing.T) {
	t.Parallel()
	if len(bookmarkAllColumns) == len(bookmarkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Bookmark{}
	if err = randomize.Struct(seed, &o, bookmarkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Bookmark: %s", err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bookmarkDBTypes, false, bookmarkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Bookmark: %s", err)
	}

	count, err = Bookmarks().Count(ctx, tx)
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

func testBookmarks(t *testing.T) {
	t.Parallel()

	query := Bookmarks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBookmarksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
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

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Bookmarks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookmarkSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BookmarkExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Bookmark exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BookmarkExists to return true, but got false.")
	}
}

func testBookmarksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bookmarkFound, err := FindBookmark(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if bookmarkFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBookmarksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Bookmarks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBookmarksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Bookmarks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBookmarksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bookmarkOne := &Bookmark{}
	bookmarkTwo := &Bookmark{}
	if err = randomize.Struct(seed, bookmarkOne, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}
	if err = randomize.Struct(seed, bookmarkTwo, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookmarkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookmarkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Bookmarks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBookmarksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bookmarkOne := &Bookmark{}
	bookmarkTwo := &Bookmark{}
	if err = randomize.Struct(seed, bookmarkOne, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}
	if err = randomize.Struct(seed, bookmarkTwo, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookmarkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookmarkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bookmarkBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func testBookmarksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Bookmark{}
	o := &Bookmark{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bookmarkDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Bookmark object: %s", err)
	}

	AddBookmarkHook(boil.BeforeInsertHook, bookmarkBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bookmarkBeforeInsertHooks = []BookmarkHook{}

	AddBookmarkHook(boil.AfterInsertHook, bookmarkAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bookmarkAfterInsertHooks = []BookmarkHook{}

	AddBookmarkHook(boil.AfterSelectHook, bookmarkAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bookmarkAfterSelectHooks = []BookmarkHook{}

	AddBookmarkHook(boil.BeforeUpdateHook, bookmarkBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bookmarkBeforeUpdateHooks = []BookmarkHook{}

	AddBookmarkHook(boil.AfterUpdateHook, bookmarkAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bookmarkAfterUpdateHooks = []BookmarkHook{}

	AddBookmarkHook(boil.BeforeDeleteHook, bookmarkBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bookmarkBeforeDeleteHooks = []BookmarkHook{}

	AddBookmarkHook(boil.AfterDeleteHook, bookmarkAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bookmarkAfterDeleteHooks = []BookmarkHook{}

	AddBookmarkHook(boil.BeforeUpsertHook, bookmarkBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bookmarkBeforeUpsertHooks = []BookmarkHook{}

	AddBookmarkHook(boil.AfterUpsertHook, bookmarkAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bookmarkAfterUpsertHooks = []BookmarkHook{}
}

func testBookmarksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookmarksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bookmarkColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookmarkToOneFileUsingFile(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Bookmark
	var foreign File

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FileID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.File().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddFileHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *File) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := BookmarkSlice{&local}
	if err = local.L.LoadFile(ctx, tx, false, (*[]*Bookmark)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.File == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.File = nil
	if err = local.L.LoadFile(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.File == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testBookmarkToOneSetOpFileUsingFile(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bookmark
	var b, c File

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkDBTypes, false, strmangle.SetComplement(bookmarkPrimaryKeyColumns, bookmarkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, fileDBTypes, false, strmangle.SetComplement(filePrimaryKeyColumns, fileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, fileDBTypes, false, strmangle.SetComplement(filePrimaryKeyColumns, fileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*File{&b, &c} {
		err = a.SetFile(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.File != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Bookmarks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FileID != x.ID {
			t.Error("foreign key was wrong value", a.FileID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FileID))
		reflect.Indirect(reflect.ValueOf(&a.FileID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FileID != x.ID {
			t.Error("foreign key was wrong value", a.FileID, x.ID)
		}
	}
}

func testBookmarksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
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

func testBookmarksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookmarkSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookmarksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Bookmarks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bookmarkDBTypes = map[string]string{`ID`: `INTEGER`, `FileID`: `INTEGER`, `Title`: `TEXT`, `Section`: `TEXT`, `Destination`: `TEXT`, `DateAccessed`: `DATETIME`}
	_               = bytes.MinRead
)

func testBookmarksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bookmarkPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bookmarkAllColumns) == len(bookmarkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBookmarksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bookmarkAllColumns) == len(bookmarkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bookmarkAllColumns, bookmarkPrimaryKeyColumns) {
		fields = bookmarkAllColumns
	} else {
		fields = strmangle.SetComplement(
			bookmarkAllColumns,
			bookmarkPrimaryKeyColumns,
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

	slice := BookmarkSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
