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

// BookmarkSearchDocsize is an object representing the database table.
type BookmarkSearchDocsize struct {
	ID null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	SZ null.Bytes `boil:"sz" json:"sz,omitempty" toml:"sz" yaml:"sz,omitempty"`

	R *bookmarkSearchDocsizeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookmarkSearchDocsizeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookmarkSearchDocsizeColumns = struct {
	ID string
	SZ string
}{
	ID: "id",
	SZ: "sz",
}

var BookmarkSearchDocsizeTableColumns = struct {
	ID string
	SZ string
}{
	ID: "bookmark_search_docsize.id",
	SZ: "bookmark_search_docsize.sz",
}

// Generated where

var BookmarkSearchDocsizeWhere = struct {
	ID whereHelpernull_Int64
	SZ whereHelpernull_Bytes
}{
	ID: whereHelpernull_Int64{field: "\"bookmark_search_docsize\".\"id\""},
	SZ: whereHelpernull_Bytes{field: "\"bookmark_search_docsize\".\"sz\""},
}

// BookmarkSearchDocsizeRels is where relationship names are stored.
var BookmarkSearchDocsizeRels = struct {
}{}

// bookmarkSearchDocsizeR is where relationships are stored.
type bookmarkSearchDocsizeR struct {
}

// NewStruct creates a new relationship struct
func (*bookmarkSearchDocsizeR) NewStruct() *bookmarkSearchDocsizeR {
	return &bookmarkSearchDocsizeR{}
}

// bookmarkSearchDocsizeL is where Load methods for each relationship are stored.
type bookmarkSearchDocsizeL struct{}

var (
	bookmarkSearchDocsizeAllColumns            = []string{"id", "sz"}
	bookmarkSearchDocsizeColumnsWithoutDefault = []string{"sz"}
	bookmarkSearchDocsizeColumnsWithDefault    = []string{"id"}
	bookmarkSearchDocsizePrimaryKeyColumns     = []string{"id"}
	bookmarkSearchDocsizeGeneratedColumns      = []string{}
)

type (
	// BookmarkSearchDocsizeSlice is an alias for a slice of pointers to BookmarkSearchDocsize.
	// This should almost always be used instead of []BookmarkSearchDocsize.
	BookmarkSearchDocsizeSlice []*BookmarkSearchDocsize
	// BookmarkSearchDocsizeHook is the signature for custom BookmarkSearchDocsize hook methods
	BookmarkSearchDocsizeHook func(context.Context, boil.ContextExecutor, *BookmarkSearchDocsize) error

	bookmarkSearchDocsizeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookmarkSearchDocsizeType                 = reflect.TypeOf(&BookmarkSearchDocsize{})
	bookmarkSearchDocsizeMapping              = queries.MakeStructMapping(bookmarkSearchDocsizeType)
	bookmarkSearchDocsizePrimaryKeyMapping, _ = queries.BindMapping(bookmarkSearchDocsizeType, bookmarkSearchDocsizeMapping, bookmarkSearchDocsizePrimaryKeyColumns)
	bookmarkSearchDocsizeInsertCacheMut       sync.RWMutex
	bookmarkSearchDocsizeInsertCache          = make(map[string]insertCache)
	bookmarkSearchDocsizeUpdateCacheMut       sync.RWMutex
	bookmarkSearchDocsizeUpdateCache          = make(map[string]updateCache)
	bookmarkSearchDocsizeUpsertCacheMut       sync.RWMutex
	bookmarkSearchDocsizeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookmarkSearchDocsizeAfterSelectHooks []BookmarkSearchDocsizeHook

var bookmarkSearchDocsizeBeforeInsertHooks []BookmarkSearchDocsizeHook
var bookmarkSearchDocsizeAfterInsertHooks []BookmarkSearchDocsizeHook

var bookmarkSearchDocsizeBeforeUpdateHooks []BookmarkSearchDocsizeHook
var bookmarkSearchDocsizeAfterUpdateHooks []BookmarkSearchDocsizeHook

var bookmarkSearchDocsizeBeforeDeleteHooks []BookmarkSearchDocsizeHook
var bookmarkSearchDocsizeAfterDeleteHooks []BookmarkSearchDocsizeHook

var bookmarkSearchDocsizeBeforeUpsertHooks []BookmarkSearchDocsizeHook
var bookmarkSearchDocsizeAfterUpsertHooks []BookmarkSearchDocsizeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BookmarkSearchDocsize) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDocsizeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BookmarkSearchDocsize) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDocsizeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BookmarkSearchDocsize) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDocsizeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BookmarkSearchDocsize) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDocsizeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BookmarkSearchDocsize) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDocsizeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BookmarkSearchDocsize) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDocsizeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BookmarkSearchDocsize) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDocsizeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BookmarkSearchDocsize) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDocsizeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BookmarkSearchDocsize) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDocsizeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookmarkSearchDocsizeHook registers your hook function for all future operations.
func AddBookmarkSearchDocsizeHook(hookPoint boil.HookPoint, bookmarkSearchDocsizeHook BookmarkSearchDocsizeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bookmarkSearchDocsizeAfterSelectHooks = append(bookmarkSearchDocsizeAfterSelectHooks, bookmarkSearchDocsizeHook)
	case boil.BeforeInsertHook:
		bookmarkSearchDocsizeBeforeInsertHooks = append(bookmarkSearchDocsizeBeforeInsertHooks, bookmarkSearchDocsizeHook)
	case boil.AfterInsertHook:
		bookmarkSearchDocsizeAfterInsertHooks = append(bookmarkSearchDocsizeAfterInsertHooks, bookmarkSearchDocsizeHook)
	case boil.BeforeUpdateHook:
		bookmarkSearchDocsizeBeforeUpdateHooks = append(bookmarkSearchDocsizeBeforeUpdateHooks, bookmarkSearchDocsizeHook)
	case boil.AfterUpdateHook:
		bookmarkSearchDocsizeAfterUpdateHooks = append(bookmarkSearchDocsizeAfterUpdateHooks, bookmarkSearchDocsizeHook)
	case boil.BeforeDeleteHook:
		bookmarkSearchDocsizeBeforeDeleteHooks = append(bookmarkSearchDocsizeBeforeDeleteHooks, bookmarkSearchDocsizeHook)
	case boil.AfterDeleteHook:
		bookmarkSearchDocsizeAfterDeleteHooks = append(bookmarkSearchDocsizeAfterDeleteHooks, bookmarkSearchDocsizeHook)
	case boil.BeforeUpsertHook:
		bookmarkSearchDocsizeBeforeUpsertHooks = append(bookmarkSearchDocsizeBeforeUpsertHooks, bookmarkSearchDocsizeHook)
	case boil.AfterUpsertHook:
		bookmarkSearchDocsizeAfterUpsertHooks = append(bookmarkSearchDocsizeAfterUpsertHooks, bookmarkSearchDocsizeHook)
	}
}

// One returns a single bookmarkSearchDocsize record from the query.
func (q bookmarkSearchDocsizeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BookmarkSearchDocsize, error) {
	o := &BookmarkSearchDocsize{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bookmark_search_docsize")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BookmarkSearchDocsize records from the query.
func (q bookmarkSearchDocsizeQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookmarkSearchDocsizeSlice, error) {
	var o []*BookmarkSearchDocsize

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BookmarkSearchDocsize slice")
	}

	if len(bookmarkSearchDocsizeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BookmarkSearchDocsize records in the query.
func (q bookmarkSearchDocsizeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bookmark_search_docsize rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bookmarkSearchDocsizeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bookmark_search_docsize exists")
	}

	return count > 0, nil
}

// BookmarkSearchDocsizes retrieves all the records using an executor.
func BookmarkSearchDocsizes(mods ...qm.QueryMod) bookmarkSearchDocsizeQuery {
	mods = append(mods, qm.From("\"bookmark_search_docsize\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"bookmark_search_docsize\".*"})
	}

	return bookmarkSearchDocsizeQuery{q}
}

// FindBookmarkSearchDocsize retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBookmarkSearchDocsize(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*BookmarkSearchDocsize, error) {
	bookmarkSearchDocsizeObj := &BookmarkSearchDocsize{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"bookmark_search_docsize\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bookmarkSearchDocsizeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bookmark_search_docsize")
	}

	if err = bookmarkSearchDocsizeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bookmarkSearchDocsizeObj, err
	}

	return bookmarkSearchDocsizeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BookmarkSearchDocsize) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bookmark_search_docsize provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookmarkSearchDocsizeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookmarkSearchDocsizeInsertCacheMut.RLock()
	cache, cached := bookmarkSearchDocsizeInsertCache[key]
	bookmarkSearchDocsizeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookmarkSearchDocsizeAllColumns,
			bookmarkSearchDocsizeColumnsWithDefault,
			bookmarkSearchDocsizeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookmarkSearchDocsizeType, bookmarkSearchDocsizeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookmarkSearchDocsizeType, bookmarkSearchDocsizeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"bookmark_search_docsize\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"bookmark_search_docsize\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into bookmark_search_docsize")
	}

	if !cached {
		bookmarkSearchDocsizeInsertCacheMut.Lock()
		bookmarkSearchDocsizeInsertCache[key] = cache
		bookmarkSearchDocsizeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BookmarkSearchDocsize.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BookmarkSearchDocsize) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookmarkSearchDocsizeUpdateCacheMut.RLock()
	cache, cached := bookmarkSearchDocsizeUpdateCache[key]
	bookmarkSearchDocsizeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookmarkSearchDocsizeAllColumns,
			bookmarkSearchDocsizePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bookmark_search_docsize, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"bookmark_search_docsize\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, bookmarkSearchDocsizePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookmarkSearchDocsizeType, bookmarkSearchDocsizeMapping, append(wl, bookmarkSearchDocsizePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update bookmark_search_docsize row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bookmark_search_docsize")
	}

	if !cached {
		bookmarkSearchDocsizeUpdateCacheMut.Lock()
		bookmarkSearchDocsizeUpdateCache[key] = cache
		bookmarkSearchDocsizeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bookmarkSearchDocsizeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bookmark_search_docsize")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bookmark_search_docsize")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookmarkSearchDocsizeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkSearchDocsizePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"bookmark_search_docsize\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkSearchDocsizePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bookmarkSearchDocsize slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bookmarkSearchDocsize")
	}
	return rowsAff, nil
}

// Delete deletes a single BookmarkSearchDocsize record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BookmarkSearchDocsize) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BookmarkSearchDocsize provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookmarkSearchDocsizePrimaryKeyMapping)
	sql := "DELETE FROM \"bookmark_search_docsize\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bookmark_search_docsize")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bookmark_search_docsize")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bookmarkSearchDocsizeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bookmarkSearchDocsizeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bookmark_search_docsize")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bookmark_search_docsize")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookmarkSearchDocsizeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookmarkSearchDocsizeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkSearchDocsizePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"bookmark_search_docsize\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkSearchDocsizePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bookmarkSearchDocsize slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bookmark_search_docsize")
	}

	if len(bookmarkSearchDocsizeAfterDeleteHooks) != 0 {
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
func (o *BookmarkSearchDocsize) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBookmarkSearchDocsize(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookmarkSearchDocsizeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookmarkSearchDocsizeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkSearchDocsizePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"bookmark_search_docsize\".* FROM \"bookmark_search_docsize\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkSearchDocsizePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BookmarkSearchDocsizeSlice")
	}

	*o = slice

	return nil
}

// BookmarkSearchDocsizeExists checks if the BookmarkSearchDocsize row exists.
func BookmarkSearchDocsizeExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"bookmark_search_docsize\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bookmark_search_docsize exists")
	}

	return exists, nil
}

// Exists checks if the BookmarkSearchDocsize row exists.
func (o *BookmarkSearchDocsize) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BookmarkSearchDocsizeExists(ctx, exec, o.ID)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BookmarkSearchDocsize) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bookmark_search_docsize provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookmarkSearchDocsizeColumnsWithDefault, o)

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

	bookmarkSearchDocsizeUpsertCacheMut.RLock()
	cache, cached := bookmarkSearchDocsizeUpsertCache[key]
	bookmarkSearchDocsizeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookmarkSearchDocsizeAllColumns,
			bookmarkSearchDocsizeColumnsWithDefault,
			bookmarkSearchDocsizeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bookmarkSearchDocsizeAllColumns,
			bookmarkSearchDocsizePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert bookmark_search_docsize, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bookmarkSearchDocsizePrimaryKeyColumns))
			copy(conflict, bookmarkSearchDocsizePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"bookmark_search_docsize\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bookmarkSearchDocsizeType, bookmarkSearchDocsizeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookmarkSearchDocsizeType, bookmarkSearchDocsizeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert bookmark_search_docsize")
	}

	if !cached {
		bookmarkSearchDocsizeUpsertCacheMut.Lock()
		bookmarkSearchDocsizeUpsertCache[key] = cache
		bookmarkSearchDocsizeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
