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

// BookmarkSearchDatum is an object representing the database table.
type BookmarkSearchDatum struct {
	ID    null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	Block null.Bytes `boil:"block" json:"block,omitempty" toml:"block" yaml:"block,omitempty"`

	R *bookmarkSearchDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookmarkSearchDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookmarkSearchDatumColumns = struct {
	ID    string
	Block string
}{
	ID:    "id",
	Block: "block",
}

var BookmarkSearchDatumTableColumns = struct {
	ID    string
	Block string
}{
	ID:    "bookmark_search_data.id",
	Block: "bookmark_search_data.block",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Bytes struct{ field string }

func (w whereHelpernull_Bytes) EQ(x null.Bytes) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bytes) NEQ(x null.Bytes) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bytes) LT(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bytes) LTE(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bytes) GT(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bytes) GTE(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Bytes) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bytes) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var BookmarkSearchDatumWhere = struct {
	ID    whereHelpernull_Int64
	Block whereHelpernull_Bytes
}{
	ID:    whereHelpernull_Int64{field: "\"bookmark_search_data\".\"id\""},
	Block: whereHelpernull_Bytes{field: "\"bookmark_search_data\".\"block\""},
}

// BookmarkSearchDatumRels is where relationship names are stored.
var BookmarkSearchDatumRels = struct {
}{}

// bookmarkSearchDatumR is where relationships are stored.
type bookmarkSearchDatumR struct {
}

// NewStruct creates a new relationship struct
func (*bookmarkSearchDatumR) NewStruct() *bookmarkSearchDatumR {
	return &bookmarkSearchDatumR{}
}

// bookmarkSearchDatumL is where Load methods for each relationship are stored.
type bookmarkSearchDatumL struct{}

var (
	bookmarkSearchDatumAllColumns            = []string{"id", "block"}
	bookmarkSearchDatumColumnsWithoutDefault = []string{"block"}
	bookmarkSearchDatumColumnsWithDefault    = []string{"id"}
	bookmarkSearchDatumPrimaryKeyColumns     = []string{"id"}
	bookmarkSearchDatumGeneratedColumns      = []string{}
)

type (
	// BookmarkSearchDatumSlice is an alias for a slice of pointers to BookmarkSearchDatum.
	// This should almost always be used instead of []BookmarkSearchDatum.
	BookmarkSearchDatumSlice []*BookmarkSearchDatum
	// BookmarkSearchDatumHook is the signature for custom BookmarkSearchDatum hook methods
	BookmarkSearchDatumHook func(context.Context, boil.ContextExecutor, *BookmarkSearchDatum) error

	bookmarkSearchDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookmarkSearchDatumType                 = reflect.TypeOf(&BookmarkSearchDatum{})
	bookmarkSearchDatumMapping              = queries.MakeStructMapping(bookmarkSearchDatumType)
	bookmarkSearchDatumPrimaryKeyMapping, _ = queries.BindMapping(bookmarkSearchDatumType, bookmarkSearchDatumMapping, bookmarkSearchDatumPrimaryKeyColumns)
	bookmarkSearchDatumInsertCacheMut       sync.RWMutex
	bookmarkSearchDatumInsertCache          = make(map[string]insertCache)
	bookmarkSearchDatumUpdateCacheMut       sync.RWMutex
	bookmarkSearchDatumUpdateCache          = make(map[string]updateCache)
	bookmarkSearchDatumUpsertCacheMut       sync.RWMutex
	bookmarkSearchDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookmarkSearchDatumAfterSelectHooks []BookmarkSearchDatumHook

var bookmarkSearchDatumBeforeInsertHooks []BookmarkSearchDatumHook
var bookmarkSearchDatumAfterInsertHooks []BookmarkSearchDatumHook

var bookmarkSearchDatumBeforeUpdateHooks []BookmarkSearchDatumHook
var bookmarkSearchDatumAfterUpdateHooks []BookmarkSearchDatumHook

var bookmarkSearchDatumBeforeDeleteHooks []BookmarkSearchDatumHook
var bookmarkSearchDatumAfterDeleteHooks []BookmarkSearchDatumHook

var bookmarkSearchDatumBeforeUpsertHooks []BookmarkSearchDatumHook
var bookmarkSearchDatumAfterUpsertHooks []BookmarkSearchDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BookmarkSearchDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BookmarkSearchDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BookmarkSearchDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BookmarkSearchDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BookmarkSearchDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BookmarkSearchDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BookmarkSearchDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BookmarkSearchDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BookmarkSearchDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkSearchDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookmarkSearchDatumHook registers your hook function for all future operations.
func AddBookmarkSearchDatumHook(hookPoint boil.HookPoint, bookmarkSearchDatumHook BookmarkSearchDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bookmarkSearchDatumAfterSelectHooks = append(bookmarkSearchDatumAfterSelectHooks, bookmarkSearchDatumHook)
	case boil.BeforeInsertHook:
		bookmarkSearchDatumBeforeInsertHooks = append(bookmarkSearchDatumBeforeInsertHooks, bookmarkSearchDatumHook)
	case boil.AfterInsertHook:
		bookmarkSearchDatumAfterInsertHooks = append(bookmarkSearchDatumAfterInsertHooks, bookmarkSearchDatumHook)
	case boil.BeforeUpdateHook:
		bookmarkSearchDatumBeforeUpdateHooks = append(bookmarkSearchDatumBeforeUpdateHooks, bookmarkSearchDatumHook)
	case boil.AfterUpdateHook:
		bookmarkSearchDatumAfterUpdateHooks = append(bookmarkSearchDatumAfterUpdateHooks, bookmarkSearchDatumHook)
	case boil.BeforeDeleteHook:
		bookmarkSearchDatumBeforeDeleteHooks = append(bookmarkSearchDatumBeforeDeleteHooks, bookmarkSearchDatumHook)
	case boil.AfterDeleteHook:
		bookmarkSearchDatumAfterDeleteHooks = append(bookmarkSearchDatumAfterDeleteHooks, bookmarkSearchDatumHook)
	case boil.BeforeUpsertHook:
		bookmarkSearchDatumBeforeUpsertHooks = append(bookmarkSearchDatumBeforeUpsertHooks, bookmarkSearchDatumHook)
	case boil.AfterUpsertHook:
		bookmarkSearchDatumAfterUpsertHooks = append(bookmarkSearchDatumAfterUpsertHooks, bookmarkSearchDatumHook)
	}
}

// One returns a single bookmarkSearchDatum record from the query.
func (q bookmarkSearchDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BookmarkSearchDatum, error) {
	o := &BookmarkSearchDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bookmark_search_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BookmarkSearchDatum records from the query.
func (q bookmarkSearchDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookmarkSearchDatumSlice, error) {
	var o []*BookmarkSearchDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BookmarkSearchDatum slice")
	}

	if len(bookmarkSearchDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BookmarkSearchDatum records in the query.
func (q bookmarkSearchDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bookmark_search_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bookmarkSearchDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bookmark_search_data exists")
	}

	return count > 0, nil
}

// BookmarkSearchData retrieves all the records using an executor.
func BookmarkSearchData(mods ...qm.QueryMod) bookmarkSearchDatumQuery {
	mods = append(mods, qm.From("\"bookmark_search_data\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"bookmark_search_data\".*"})
	}

	return bookmarkSearchDatumQuery{q}
}

// FindBookmarkSearchDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBookmarkSearchDatum(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*BookmarkSearchDatum, error) {
	bookmarkSearchDatumObj := &BookmarkSearchDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"bookmark_search_data\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bookmarkSearchDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bookmark_search_data")
	}

	if err = bookmarkSearchDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bookmarkSearchDatumObj, err
	}

	return bookmarkSearchDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BookmarkSearchDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bookmark_search_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookmarkSearchDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookmarkSearchDatumInsertCacheMut.RLock()
	cache, cached := bookmarkSearchDatumInsertCache[key]
	bookmarkSearchDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookmarkSearchDatumAllColumns,
			bookmarkSearchDatumColumnsWithDefault,
			bookmarkSearchDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookmarkSearchDatumType, bookmarkSearchDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookmarkSearchDatumType, bookmarkSearchDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"bookmark_search_data\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"bookmark_search_data\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into bookmark_search_data")
	}

	if !cached {
		bookmarkSearchDatumInsertCacheMut.Lock()
		bookmarkSearchDatumInsertCache[key] = cache
		bookmarkSearchDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BookmarkSearchDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BookmarkSearchDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookmarkSearchDatumUpdateCacheMut.RLock()
	cache, cached := bookmarkSearchDatumUpdateCache[key]
	bookmarkSearchDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookmarkSearchDatumAllColumns,
			bookmarkSearchDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bookmark_search_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"bookmark_search_data\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, bookmarkSearchDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookmarkSearchDatumType, bookmarkSearchDatumMapping, append(wl, bookmarkSearchDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update bookmark_search_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bookmark_search_data")
	}

	if !cached {
		bookmarkSearchDatumUpdateCacheMut.Lock()
		bookmarkSearchDatumUpdateCache[key] = cache
		bookmarkSearchDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bookmarkSearchDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bookmark_search_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bookmark_search_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookmarkSearchDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkSearchDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"bookmark_search_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkSearchDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bookmarkSearchDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bookmarkSearchDatum")
	}
	return rowsAff, nil
}

// Delete deletes a single BookmarkSearchDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BookmarkSearchDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BookmarkSearchDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookmarkSearchDatumPrimaryKeyMapping)
	sql := "DELETE FROM \"bookmark_search_data\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bookmark_search_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bookmark_search_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bookmarkSearchDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bookmarkSearchDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bookmark_search_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bookmark_search_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookmarkSearchDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookmarkSearchDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkSearchDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"bookmark_search_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkSearchDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bookmarkSearchDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bookmark_search_data")
	}

	if len(bookmarkSearchDatumAfterDeleteHooks) != 0 {
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
func (o *BookmarkSearchDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBookmarkSearchDatum(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookmarkSearchDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookmarkSearchDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkSearchDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"bookmark_search_data\".* FROM \"bookmark_search_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkSearchDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BookmarkSearchDatumSlice")
	}

	*o = slice

	return nil
}

// BookmarkSearchDatumExists checks if the BookmarkSearchDatum row exists.
func BookmarkSearchDatumExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"bookmark_search_data\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bookmark_search_data exists")
	}

	return exists, nil
}

// Exists checks if the BookmarkSearchDatum row exists.
func (o *BookmarkSearchDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BookmarkSearchDatumExists(ctx, exec, o.ID)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BookmarkSearchDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bookmark_search_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookmarkSearchDatumColumnsWithDefault, o)

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

	bookmarkSearchDatumUpsertCacheMut.RLock()
	cache, cached := bookmarkSearchDatumUpsertCache[key]
	bookmarkSearchDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookmarkSearchDatumAllColumns,
			bookmarkSearchDatumColumnsWithDefault,
			bookmarkSearchDatumColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bookmarkSearchDatumAllColumns,
			bookmarkSearchDatumPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert bookmark_search_data, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bookmarkSearchDatumPrimaryKeyColumns))
			copy(conflict, bookmarkSearchDatumPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"bookmark_search_data\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bookmarkSearchDatumType, bookmarkSearchDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookmarkSearchDatumType, bookmarkSearchDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert bookmark_search_data")
	}

	if !cached {
		bookmarkSearchDatumUpsertCacheMut.Lock()
		bookmarkSearchDatumUpsertCache[key] = cache
		bookmarkSearchDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
