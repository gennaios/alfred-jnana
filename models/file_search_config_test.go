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

func testFileSearchConfigsUpsert(t *testing.T) {
	t.Parallel()
	if len(fileSearchConfigAllColumns) == len(fileSearchConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := FileSearchConfig{}
	if err = randomize.Struct(seed, &o, fileSearchConfigDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FileSearchConfig: %s", err)
	}

	count, err := FileSearchConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, fileSearchConfigDBTypes, false, fileSearchConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FileSearchConfig: %s", err)
	}

	count, err = FileSearchConfigs().Count(ctx, tx)
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

func testFileSearchConfigs(t *testing.T) {
	t.Parallel()

	query := FileSearchConfigs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFileSearchConfigsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
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

	count, err := FileSearchConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFileSearchConfigsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := FileSearchConfigs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FileSearchConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFileSearchConfigsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FileSearchConfigSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FileSearchConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFileSearchConfigsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FileSearchConfigExists(ctx, tx, o.K)
	if err != nil {
		t.Errorf("Unable to check if FileSearchConfig exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FileSearchConfigExists to return true, but got false.")
	}
}

func testFileSearchConfigsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	fileSearchConfigFound, err := FindFileSearchConfig(ctx, tx, o.K)
	if err != nil {
		t.Error(err)
	}

	if fileSearchConfigFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFileSearchConfigsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = FileSearchConfigs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFileSearchConfigsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := FileSearchConfigs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFileSearchConfigsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	fileSearchConfigOne := &FileSearchConfig{}
	fileSearchConfigTwo := &FileSearchConfig{}
	if err = randomize.Struct(seed, fileSearchConfigOne, fileSearchConfigDBTypes, false, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}
	if err = randomize.Struct(seed, fileSearchConfigTwo, fileSearchConfigDBTypes, false, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fileSearchConfigOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fileSearchConfigTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FileSearchConfigs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFileSearchConfigsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	fileSearchConfigOne := &FileSearchConfig{}
	fileSearchConfigTwo := &FileSearchConfig{}
	if err = randomize.Struct(seed, fileSearchConfigOne, fileSearchConfigDBTypes, false, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}
	if err = randomize.Struct(seed, fileSearchConfigTwo, fileSearchConfigDBTypes, false, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fileSearchConfigOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fileSearchConfigTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FileSearchConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func fileSearchConfigBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *FileSearchConfig) error {
	*o = FileSearchConfig{}
	return nil
}

func fileSearchConfigAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *FileSearchConfig) error {
	*o = FileSearchConfig{}
	return nil
}

func fileSearchConfigAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *FileSearchConfig) error {
	*o = FileSearchConfig{}
	return nil
}

func fileSearchConfigBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FileSearchConfig) error {
	*o = FileSearchConfig{}
	return nil
}

func fileSearchConfigAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FileSearchConfig) error {
	*o = FileSearchConfig{}
	return nil
}

func fileSearchConfigBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FileSearchConfig) error {
	*o = FileSearchConfig{}
	return nil
}

func fileSearchConfigAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FileSearchConfig) error {
	*o = FileSearchConfig{}
	return nil
}

func fileSearchConfigBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FileSearchConfig) error {
	*o = FileSearchConfig{}
	return nil
}

func fileSearchConfigAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FileSearchConfig) error {
	*o = FileSearchConfig{}
	return nil
}

func testFileSearchConfigsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &FileSearchConfig{}
	o := &FileSearchConfig{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig object: %s", err)
	}

	AddFileSearchConfigHook(boil.BeforeInsertHook, fileSearchConfigBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	fileSearchConfigBeforeInsertHooks = []FileSearchConfigHook{}

	AddFileSearchConfigHook(boil.AfterInsertHook, fileSearchConfigAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	fileSearchConfigAfterInsertHooks = []FileSearchConfigHook{}

	AddFileSearchConfigHook(boil.AfterSelectHook, fileSearchConfigAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	fileSearchConfigAfterSelectHooks = []FileSearchConfigHook{}

	AddFileSearchConfigHook(boil.BeforeUpdateHook, fileSearchConfigBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	fileSearchConfigBeforeUpdateHooks = []FileSearchConfigHook{}

	AddFileSearchConfigHook(boil.AfterUpdateHook, fileSearchConfigAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	fileSearchConfigAfterUpdateHooks = []FileSearchConfigHook{}

	AddFileSearchConfigHook(boil.BeforeDeleteHook, fileSearchConfigBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	fileSearchConfigBeforeDeleteHooks = []FileSearchConfigHook{}

	AddFileSearchConfigHook(boil.AfterDeleteHook, fileSearchConfigAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	fileSearchConfigAfterDeleteHooks = []FileSearchConfigHook{}

	AddFileSearchConfigHook(boil.BeforeUpsertHook, fileSearchConfigBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	fileSearchConfigBeforeUpsertHooks = []FileSearchConfigHook{}

	AddFileSearchConfigHook(boil.AfterUpsertHook, fileSearchConfigAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	fileSearchConfigAfterUpsertHooks = []FileSearchConfigHook{}
}

func testFileSearchConfigsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FileSearchConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFileSearchConfigsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(fileSearchConfigColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := FileSearchConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFileSearchConfigsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
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

func testFileSearchConfigsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FileSearchConfigSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFileSearchConfigsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FileSearchConfigs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	fileSearchConfigDBTypes = map[string]string{`K`: ``, `V`: ``}
	_                       = bytes.MinRead
)

func testFileSearchConfigsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(fileSearchConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(fileSearchConfigAllColumns) == len(fileSearchConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FileSearchConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFileSearchConfigsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(fileSearchConfigAllColumns) == len(fileSearchConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FileSearchConfig{}
	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FileSearchConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fileSearchConfigDBTypes, true, fileSearchConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FileSearchConfig struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(fileSearchConfigAllColumns, fileSearchConfigPrimaryKeyColumns) {
		fields = fileSearchConfigAllColumns
	} else {
		fields = strmangle.SetComplement(
			fileSearchConfigAllColumns,
			fileSearchConfigPrimaryKeyColumns,
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

	slice := FileSearchConfigSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
