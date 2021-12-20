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

func testFilesUpsert(t *testing.T) {
	t.Parallel()
	if len(fileAllColumns) == len(filePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := File{}
	if err = randomize.Struct(seed, &o, fileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert File: %s", err)
	}

	count, err := Files().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, fileDBTypes, false, filePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert File: %s", err)
	}

	count, err = Files().Count(ctx, tx)
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

func testFiles(t *testing.T) {
	t.Parallel()

	query := Files()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFilesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
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

	count, err := Files().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFilesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Files().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Files().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFilesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FileSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Files().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFilesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FileExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if File exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FileExists to return true, but got false.")
	}
}

func testFilesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	fileFound, err := FindFile(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if fileFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFilesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Files().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFilesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Files().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFilesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	fileOne := &File{}
	fileTwo := &File{}
	if err = randomize.Struct(seed, fileOne, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}
	if err = randomize.Struct(seed, fileTwo, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fileOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fileTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Files().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFilesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	fileOne := &File{}
	fileTwo := &File{}
	if err = randomize.Struct(seed, fileOne, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}
	if err = randomize.Struct(seed, fileTwo, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fileOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fileTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Files().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func fileBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *File) error {
	*o = File{}
	return nil
}

func fileAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *File) error {
	*o = File{}
	return nil
}

func fileAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *File) error {
	*o = File{}
	return nil
}

func fileBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *File) error {
	*o = File{}
	return nil
}

func fileAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *File) error {
	*o = File{}
	return nil
}

func fileBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *File) error {
	*o = File{}
	return nil
}

func fileAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *File) error {
	*o = File{}
	return nil
}

func fileBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *File) error {
	*o = File{}
	return nil
}

func fileAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *File) error {
	*o = File{}
	return nil
}

func testFilesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &File{}
	o := &File{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, fileDBTypes, false); err != nil {
		t.Errorf("Unable to randomize File object: %s", err)
	}

	AddFileHook(boil.BeforeInsertHook, fileBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	fileBeforeInsertHooks = []FileHook{}

	AddFileHook(boil.AfterInsertHook, fileAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	fileAfterInsertHooks = []FileHook{}

	AddFileHook(boil.AfterSelectHook, fileAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	fileAfterSelectHooks = []FileHook{}

	AddFileHook(boil.BeforeUpdateHook, fileBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	fileBeforeUpdateHooks = []FileHook{}

	AddFileHook(boil.AfterUpdateHook, fileAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	fileAfterUpdateHooks = []FileHook{}

	AddFileHook(boil.BeforeDeleteHook, fileBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	fileBeforeDeleteHooks = []FileHook{}

	AddFileHook(boil.AfterDeleteHook, fileAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	fileAfterDeleteHooks = []FileHook{}

	AddFileHook(boil.BeforeUpsertHook, fileBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	fileBeforeUpsertHooks = []FileHook{}

	AddFileHook(boil.AfterUpsertHook, fileAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	fileAfterUpsertHooks = []FileHook{}
}

func testFilesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Files().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFilesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(fileColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Files().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFileToManyBookmarks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a File
	var b, c Bookmark

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.FileID = a.ID
	c.FileID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Bookmarks().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.FileID == b.FileID {
			bFound = true
		}
		if v.FileID == c.FileID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := FileSlice{&a}
	if err = a.L.LoadBookmarks(ctx, tx, false, (*[]*File)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Bookmarks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Bookmarks = nil
	if err = a.L.LoadBookmarks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Bookmarks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testFileToManyAddOpBookmarks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a File
	var b, c, d, e Bookmark

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fileDBTypes, false, strmangle.SetComplement(filePrimaryKeyColumns, fileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Bookmark{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookmarkDBTypes, false, strmangle.SetComplement(bookmarkPrimaryKeyColumns, bookmarkColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Bookmark{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBookmarks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.FileID {
			t.Error("foreign key was wrong value", a.ID, first.FileID)
		}
		if a.ID != second.FileID {
			t.Error("foreign key was wrong value", a.ID, second.FileID)
		}

		if first.R.File != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.File != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Bookmarks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Bookmarks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Bookmarks().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testFilesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
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

func testFilesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FileSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFilesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Files().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	fileDBTypes = map[string]string{`ID`: `INTEGER`, `Path`: `TEXT`, `Name`: `TEXT`, `Extension`: `VARCHAR(3)`, `Size`: `INTEGER`, `Title`: `TEXT`, `Publisher`: `TEXT`, `PublisherID`: `INTEGER`, `Creator`: `TEXT`, `Subject`: `TEXT`, `Language`: `TEXT`, `Description`: `TEXT`, `DateCreated`: `DATETIME`, `DateModified`: `DATETIME`, `DateAccessed`: `DATETIME`, `Rating`: `INTEGER`, `Hash`: `VARCHAR(64)`}
	_           = bytes.MinRead
)

func testFilesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(filePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(fileAllColumns) == len(filePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Files().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fileDBTypes, true, filePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFilesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(fileAllColumns) == len(filePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &File{}
	if err = randomize.Struct(seed, o, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Files().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fileDBTypes, true, filePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(fileAllColumns, filePrimaryKeyColumns) {
		fields = fileAllColumns
	} else {
		fields = strmangle.SetComplement(
			fileAllColumns,
			filePrimaryKeyColumns,
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

	slice := FileSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
