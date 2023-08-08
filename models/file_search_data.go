// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// FileSearchDatum is an object representing the database table.
type FileSearchDatum struct {
	ID    null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	Block null.Bytes `boil:"block" json:"block,omitempty" toml:"block" yaml:"block,omitempty"`

	R *fileSearchDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L fileSearchDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FileSearchDatumColumns = struct {
	ID    string
	Block string
}{
	ID:    "id",
	Block: "block",
}

var FileSearchDatumTableColumns = struct {
	ID    string
	Block string
}{
	ID:    "file_search_data.id",
	Block: "file_search_data.block",
}

// Generated where

var FileSearchDatumWhere = struct {
	ID    whereHelpernull_Int64
	Block whereHelpernull_Bytes
}{
	ID:    whereHelpernull_Int64{field: "\"file_search_data\".\"id\""},
	Block: whereHelpernull_Bytes{field: "\"file_search_data\".\"block\""},
}

// FileSearchDatumRels is where relationship names are stored.
var FileSearchDatumRels = struct {
}{}

// fileSearchDatumR is where relationships are stored.
type fileSearchDatumR struct {
}

// NewStruct creates a new relationship struct
func (*fileSearchDatumR) NewStruct() *fileSearchDatumR {
	return &fileSearchDatumR{}
}

// fileSearchDatumL is where Load methods for each relationship are stored.
type fileSearchDatumL struct{}

var (
	fileSearchDatumAllColumns            = []string{"id", "block"}
	fileSearchDatumColumnsWithoutDefault = []string{"block"}
	fileSearchDatumColumnsWithDefault    = []string{"id"}
	fileSearchDatumPrimaryKeyColumns     = []string{"id"}
	fileSearchDatumGeneratedColumns      = []string{}
)

type (
	// FileSearchDatumSlice is an alias for a slice of pointers to FileSearchDatum.
	// This should almost always be used instead of []FileSearchDatum.
	FileSearchDatumSlice []*FileSearchDatum
	// FileSearchDatumHook is the signature for custom FileSearchDatum hook methods
	FileSearchDatumHook func(context.Context, boil.ContextExecutor, *FileSearchDatum) error

	fileSearchDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	fileSearchDatumType                 = reflect.TypeOf(&FileSearchDatum{})
	fileSearchDatumMapping              = queries.MakeStructMapping(fileSearchDatumType)
	fileSearchDatumPrimaryKeyMapping, _ = queries.BindMapping(fileSearchDatumType, fileSearchDatumMapping, fileSearchDatumPrimaryKeyColumns)
	fileSearchDatumInsertCacheMut       sync.RWMutex
	fileSearchDatumInsertCache          = make(map[string]insertCache)
	fileSearchDatumUpdateCacheMut       sync.RWMutex
	fileSearchDatumUpdateCache          = make(map[string]updateCache)
	fileSearchDatumUpsertCacheMut       sync.RWMutex
	fileSearchDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var fileSearchDatumAfterSelectHooks []FileSearchDatumHook

var fileSearchDatumBeforeInsertHooks []FileSearchDatumHook
var fileSearchDatumAfterInsertHooks []FileSearchDatumHook

var fileSearchDatumBeforeUpdateHooks []FileSearchDatumHook
var fileSearchDatumAfterUpdateHooks []FileSearchDatumHook

var fileSearchDatumBeforeDeleteHooks []FileSearchDatumHook
var fileSearchDatumAfterDeleteHooks []FileSearchDatumHook

var fileSearchDatumBeforeUpsertHooks []FileSearchDatumHook
var fileSearchDatumAfterUpsertHooks []FileSearchDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FileSearchDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileSearchDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FileSearchDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileSearchDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FileSearchDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileSearchDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FileSearchDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileSearchDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FileSearchDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileSearchDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FileSearchDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileSearchDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FileSearchDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileSearchDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FileSearchDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileSearchDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FileSearchDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileSearchDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFileSearchDatumHook registers your hook function for all future operations.
func AddFileSearchDatumHook(hookPoint boil.HookPoint, fileSearchDatumHook FileSearchDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		fileSearchDatumAfterSelectHooks = append(fileSearchDatumAfterSelectHooks, fileSearchDatumHook)
	case boil.BeforeInsertHook:
		fileSearchDatumBeforeInsertHooks = append(fileSearchDatumBeforeInsertHooks, fileSearchDatumHook)
	case boil.AfterInsertHook:
		fileSearchDatumAfterInsertHooks = append(fileSearchDatumAfterInsertHooks, fileSearchDatumHook)
	case boil.BeforeUpdateHook:
		fileSearchDatumBeforeUpdateHooks = append(fileSearchDatumBeforeUpdateHooks, fileSearchDatumHook)
	case boil.AfterUpdateHook:
		fileSearchDatumAfterUpdateHooks = append(fileSearchDatumAfterUpdateHooks, fileSearchDatumHook)
	case boil.BeforeDeleteHook:
		fileSearchDatumBeforeDeleteHooks = append(fileSearchDatumBeforeDeleteHooks, fileSearchDatumHook)
	case boil.AfterDeleteHook:
		fileSearchDatumAfterDeleteHooks = append(fileSearchDatumAfterDeleteHooks, fileSearchDatumHook)
	case boil.BeforeUpsertHook:
		fileSearchDatumBeforeUpsertHooks = append(fileSearchDatumBeforeUpsertHooks, fileSearchDatumHook)
	case boil.AfterUpsertHook:
		fileSearchDatumAfterUpsertHooks = append(fileSearchDatumAfterUpsertHooks, fileSearchDatumHook)
	}
}

// One returns a single fileSearchDatum record from the query.
func (q fileSearchDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*FileSearchDatum, error) {
	o := &FileSearchDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for file_search_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all FileSearchDatum records from the query.
func (q fileSearchDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (FileSearchDatumSlice, error) {
	var o []*FileSearchDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FileSearchDatum slice")
	}

	if len(fileSearchDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all FileSearchDatum records in the query.
func (q fileSearchDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count file_search_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q fileSearchDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if file_search_data exists")
	}

	return count > 0, nil
}

// FileSearchData retrieves all the records using an executor.
func FileSearchData(mods ...qm.QueryMod) fileSearchDatumQuery {
	mods = append(mods, qm.From("\"file_search_data\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"file_search_data\".*"})
	}

	return fileSearchDatumQuery{q}
}

// FindFileSearchDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFileSearchDatum(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*FileSearchDatum, error) {
	fileSearchDatumObj := &FileSearchDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"file_search_data\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, fileSearchDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from file_search_data")
	}

	if err = fileSearchDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return fileSearchDatumObj, err
	}

	return fileSearchDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *FileSearchDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no file_search_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fileSearchDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	fileSearchDatumInsertCacheMut.RLock()
	cache, cached := fileSearchDatumInsertCache[key]
	fileSearchDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			fileSearchDatumAllColumns,
			fileSearchDatumColumnsWithDefault,
			fileSearchDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(fileSearchDatumType, fileSearchDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(fileSearchDatumType, fileSearchDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"file_search_data\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"file_search_data\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into file_search_data")
	}

	if !cached {
		fileSearchDatumInsertCacheMut.Lock()
		fileSearchDatumInsertCache[key] = cache
		fileSearchDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the FileSearchDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *FileSearchDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	fileSearchDatumUpdateCacheMut.RLock()
	cache, cached := fileSearchDatumUpdateCache[key]
	fileSearchDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			fileSearchDatumAllColumns,
			fileSearchDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update file_search_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"file_search_data\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, fileSearchDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(fileSearchDatumType, fileSearchDatumMapping, append(wl, fileSearchDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update file_search_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for file_search_data")
	}

	if !cached {
		fileSearchDatumUpdateCacheMut.Lock()
		fileSearchDatumUpdateCache[key] = cache
		fileSearchDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q fileSearchDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for file_search_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for file_search_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FileSearchDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fileSearchDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"file_search_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, fileSearchDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in fileSearchDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all fileSearchDatum")
	}
	return rowsAff, nil
}

// Delete deletes a single FileSearchDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FileSearchDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no FileSearchDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), fileSearchDatumPrimaryKeyMapping)
	sql := "DELETE FROM \"file_search_data\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from file_search_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for file_search_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q fileSearchDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no fileSearchDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from file_search_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for file_search_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FileSearchDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(fileSearchDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fileSearchDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"file_search_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, fileSearchDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from fileSearchDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for file_search_data")
	}

	if len(fileSearchDatumAfterDeleteHooks) != 0 {
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
func (o *FileSearchDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFileSearchDatum(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FileSearchDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FileSearchDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fileSearchDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"file_search_data\".* FROM \"file_search_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, fileSearchDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FileSearchDatumSlice")
	}

	*o = slice

	return nil
}

// FileSearchDatumExists checks if the FileSearchDatum row exists.
func FileSearchDatumExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"file_search_data\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if file_search_data exists")
	}

	return exists, nil
}

// Exists checks if the FileSearchDatum row exists.
func (o *FileSearchDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return FileSearchDatumExists(ctx, exec, o.ID)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *FileSearchDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no file_search_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fileSearchDatumColumnsWithDefault, o)

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

	fileSearchDatumUpsertCacheMut.RLock()
	cache, cached := fileSearchDatumUpsertCache[key]
	fileSearchDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			fileSearchDatumAllColumns,
			fileSearchDatumColumnsWithDefault,
			fileSearchDatumColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			fileSearchDatumAllColumns,
			fileSearchDatumPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert file_search_data, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(fileSearchDatumPrimaryKeyColumns))
			copy(conflict, fileSearchDatumPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"file_search_data\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(fileSearchDatumType, fileSearchDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(fileSearchDatumType, fileSearchDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert file_search_data")
	}

	if !cached {
		fileSearchDatumUpsertCacheMut.Lock()
		fileSearchDatumUpsertCache[key] = cache
		fileSearchDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
