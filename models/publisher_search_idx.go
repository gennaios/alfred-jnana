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

// PublisherSearchIdx is an object representing the database table.
type PublisherSearchIdx struct {
	Segid string      `boil:"segid" json:"segid" toml:"segid" yaml:"segid"`
	Term  string      `boil:"term" json:"term" toml:"term" yaml:"term"`
	Pgno  null.String `boil:"pgno" json:"pgno,omitempty" toml:"pgno" yaml:"pgno,omitempty"`

	R *publisherSearchIdxR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L publisherSearchIdxL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PublisherSearchIdxColumns = struct {
	Segid string
	Term  string
	Pgno  string
}{
	Segid: "segid",
	Term:  "term",
	Pgno:  "pgno",
}

var PublisherSearchIdxTableColumns = struct {
	Segid string
	Term  string
	Pgno  string
}{
	Segid: "publisher_search_idx.segid",
	Term:  "publisher_search_idx.term",
	Pgno:  "publisher_search_idx.pgno",
}

// Generated where

var PublisherSearchIdxWhere = struct {
	Segid whereHelperstring
	Term  whereHelperstring
	Pgno  whereHelpernull_String
}{
	Segid: whereHelperstring{field: "\"publisher_search_idx\".\"segid\""},
	Term:  whereHelperstring{field: "\"publisher_search_idx\".\"term\""},
	Pgno:  whereHelpernull_String{field: "\"publisher_search_idx\".\"pgno\""},
}

// PublisherSearchIdxRels is where relationship names are stored.
var PublisherSearchIdxRels = struct {
}{}

// publisherSearchIdxR is where relationships are stored.
type publisherSearchIdxR struct {
}

// NewStruct creates a new relationship struct
func (*publisherSearchIdxR) NewStruct() *publisherSearchIdxR {
	return &publisherSearchIdxR{}
}

// publisherSearchIdxL is where Load methods for each relationship are stored.
type publisherSearchIdxL struct{}

var (
	publisherSearchIdxAllColumns            = []string{"segid", "term", "pgno"}
	publisherSearchIdxColumnsWithoutDefault = []string{"segid", "term", "pgno"}
	publisherSearchIdxColumnsWithDefault    = []string{}
	publisherSearchIdxPrimaryKeyColumns     = []string{"segid", "term"}
	publisherSearchIdxGeneratedColumns      = []string{}
)

type (
	// PublisherSearchIdxSlice is an alias for a slice of pointers to PublisherSearchIdx.
	// This should almost always be used instead of []PublisherSearchIdx.
	PublisherSearchIdxSlice []*PublisherSearchIdx
	// PublisherSearchIdxHook is the signature for custom PublisherSearchIdx hook methods
	PublisherSearchIdxHook func(context.Context, boil.ContextExecutor, *PublisherSearchIdx) error

	publisherSearchIdxQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	publisherSearchIdxType                 = reflect.TypeOf(&PublisherSearchIdx{})
	publisherSearchIdxMapping              = queries.MakeStructMapping(publisherSearchIdxType)
	publisherSearchIdxPrimaryKeyMapping, _ = queries.BindMapping(publisherSearchIdxType, publisherSearchIdxMapping, publisherSearchIdxPrimaryKeyColumns)
	publisherSearchIdxInsertCacheMut       sync.RWMutex
	publisherSearchIdxInsertCache          = make(map[string]insertCache)
	publisherSearchIdxUpdateCacheMut       sync.RWMutex
	publisherSearchIdxUpdateCache          = make(map[string]updateCache)
	publisherSearchIdxUpsertCacheMut       sync.RWMutex
	publisherSearchIdxUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var publisherSearchIdxAfterSelectHooks []PublisherSearchIdxHook

var publisherSearchIdxBeforeInsertHooks []PublisherSearchIdxHook
var publisherSearchIdxAfterInsertHooks []PublisherSearchIdxHook

var publisherSearchIdxBeforeUpdateHooks []PublisherSearchIdxHook
var publisherSearchIdxAfterUpdateHooks []PublisherSearchIdxHook

var publisherSearchIdxBeforeDeleteHooks []PublisherSearchIdxHook
var publisherSearchIdxAfterDeleteHooks []PublisherSearchIdxHook

var publisherSearchIdxBeforeUpsertHooks []PublisherSearchIdxHook
var publisherSearchIdxAfterUpsertHooks []PublisherSearchIdxHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PublisherSearchIdx) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherSearchIdxAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PublisherSearchIdx) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherSearchIdxBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PublisherSearchIdx) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherSearchIdxAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PublisherSearchIdx) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherSearchIdxBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PublisherSearchIdx) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherSearchIdxAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PublisherSearchIdx) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherSearchIdxBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PublisherSearchIdx) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherSearchIdxAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PublisherSearchIdx) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherSearchIdxBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PublisherSearchIdx) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherSearchIdxAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPublisherSearchIdxHook registers your hook function for all future operations.
func AddPublisherSearchIdxHook(hookPoint boil.HookPoint, publisherSearchIdxHook PublisherSearchIdxHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		publisherSearchIdxAfterSelectHooks = append(publisherSearchIdxAfterSelectHooks, publisherSearchIdxHook)
	case boil.BeforeInsertHook:
		publisherSearchIdxBeforeInsertHooks = append(publisherSearchIdxBeforeInsertHooks, publisherSearchIdxHook)
	case boil.AfterInsertHook:
		publisherSearchIdxAfterInsertHooks = append(publisherSearchIdxAfterInsertHooks, publisherSearchIdxHook)
	case boil.BeforeUpdateHook:
		publisherSearchIdxBeforeUpdateHooks = append(publisherSearchIdxBeforeUpdateHooks, publisherSearchIdxHook)
	case boil.AfterUpdateHook:
		publisherSearchIdxAfterUpdateHooks = append(publisherSearchIdxAfterUpdateHooks, publisherSearchIdxHook)
	case boil.BeforeDeleteHook:
		publisherSearchIdxBeforeDeleteHooks = append(publisherSearchIdxBeforeDeleteHooks, publisherSearchIdxHook)
	case boil.AfterDeleteHook:
		publisherSearchIdxAfterDeleteHooks = append(publisherSearchIdxAfterDeleteHooks, publisherSearchIdxHook)
	case boil.BeforeUpsertHook:
		publisherSearchIdxBeforeUpsertHooks = append(publisherSearchIdxBeforeUpsertHooks, publisherSearchIdxHook)
	case boil.AfterUpsertHook:
		publisherSearchIdxAfterUpsertHooks = append(publisherSearchIdxAfterUpsertHooks, publisherSearchIdxHook)
	}
}

// One returns a single publisherSearchIdx record from the query.
func (q publisherSearchIdxQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PublisherSearchIdx, error) {
	o := &PublisherSearchIdx{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for publisher_search_idx")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PublisherSearchIdx records from the query.
func (q publisherSearchIdxQuery) All(ctx context.Context, exec boil.ContextExecutor) (PublisherSearchIdxSlice, error) {
	var o []*PublisherSearchIdx

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PublisherSearchIdx slice")
	}

	if len(publisherSearchIdxAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PublisherSearchIdx records in the query.
func (q publisherSearchIdxQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count publisher_search_idx rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q publisherSearchIdxQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if publisher_search_idx exists")
	}

	return count > 0, nil
}

// PublisherSearchIdxes retrieves all the records using an executor.
func PublisherSearchIdxes(mods ...qm.QueryMod) publisherSearchIdxQuery {
	mods = append(mods, qm.From("\"publisher_search_idx\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"publisher_search_idx\".*"})
	}

	return publisherSearchIdxQuery{q}
}

// FindPublisherSearchIdx retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPublisherSearchIdx(ctx context.Context, exec boil.ContextExecutor, segid string, term string, selectCols ...string) (*PublisherSearchIdx, error) {
	publisherSearchIdxObj := &PublisherSearchIdx{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"publisher_search_idx\" where \"segid\"=? AND \"term\"=?", sel,
	)

	q := queries.Raw(query, segid, term)

	err := q.Bind(ctx, exec, publisherSearchIdxObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from publisher_search_idx")
	}

	if err = publisherSearchIdxObj.doAfterSelectHooks(ctx, exec); err != nil {
		return publisherSearchIdxObj, err
	}

	return publisherSearchIdxObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PublisherSearchIdx) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publisher_search_idx provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(publisherSearchIdxColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	publisherSearchIdxInsertCacheMut.RLock()
	cache, cached := publisherSearchIdxInsertCache[key]
	publisherSearchIdxInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			publisherSearchIdxAllColumns,
			publisherSearchIdxColumnsWithDefault,
			publisherSearchIdxColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(publisherSearchIdxType, publisherSearchIdxMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(publisherSearchIdxType, publisherSearchIdxMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"publisher_search_idx\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"publisher_search_idx\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into publisher_search_idx")
	}

	if !cached {
		publisherSearchIdxInsertCacheMut.Lock()
		publisherSearchIdxInsertCache[key] = cache
		publisherSearchIdxInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PublisherSearchIdx.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PublisherSearchIdx) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	publisherSearchIdxUpdateCacheMut.RLock()
	cache, cached := publisherSearchIdxUpdateCache[key]
	publisherSearchIdxUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			publisherSearchIdxAllColumns,
			publisherSearchIdxPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update publisher_search_idx, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"publisher_search_idx\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, publisherSearchIdxPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(publisherSearchIdxType, publisherSearchIdxMapping, append(wl, publisherSearchIdxPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update publisher_search_idx row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for publisher_search_idx")
	}

	if !cached {
		publisherSearchIdxUpdateCacheMut.Lock()
		publisherSearchIdxUpdateCache[key] = cache
		publisherSearchIdxUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q publisherSearchIdxQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for publisher_search_idx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for publisher_search_idx")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PublisherSearchIdxSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherSearchIdxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"publisher_search_idx\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, publisherSearchIdxPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in publisherSearchIdx slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all publisherSearchIdx")
	}
	return rowsAff, nil
}

// Delete deletes a single PublisherSearchIdx record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PublisherSearchIdx) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PublisherSearchIdx provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), publisherSearchIdxPrimaryKeyMapping)
	sql := "DELETE FROM \"publisher_search_idx\" WHERE \"segid\"=? AND \"term\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from publisher_search_idx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for publisher_search_idx")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q publisherSearchIdxQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no publisherSearchIdxQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publisher_search_idx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publisher_search_idx")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PublisherSearchIdxSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(publisherSearchIdxBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherSearchIdxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"publisher_search_idx\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, publisherSearchIdxPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publisherSearchIdx slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publisher_search_idx")
	}

	if len(publisherSearchIdxAfterDeleteHooks) != 0 {
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
func (o *PublisherSearchIdx) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPublisherSearchIdx(ctx, exec, o.Segid, o.Term)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PublisherSearchIdxSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PublisherSearchIdxSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherSearchIdxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"publisher_search_idx\".* FROM \"publisher_search_idx\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, publisherSearchIdxPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PublisherSearchIdxSlice")
	}

	*o = slice

	return nil
}

// PublisherSearchIdxExists checks if the PublisherSearchIdx row exists.
func PublisherSearchIdxExists(ctx context.Context, exec boil.ContextExecutor, segid string, term string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"publisher_search_idx\" where \"segid\"=? AND \"term\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, segid, term)
	}
	row := exec.QueryRowContext(ctx, sql, segid, term)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if publisher_search_idx exists")
	}

	return exists, nil
}

// Exists checks if the PublisherSearchIdx row exists.
func (o *PublisherSearchIdx) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PublisherSearchIdxExists(ctx, exec, o.Segid, o.Term)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PublisherSearchIdx) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publisher_search_idx provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(publisherSearchIdxColumnsWithDefault, o)

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

	publisherSearchIdxUpsertCacheMut.RLock()
	cache, cached := publisherSearchIdxUpsertCache[key]
	publisherSearchIdxUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			publisherSearchIdxAllColumns,
			publisherSearchIdxColumnsWithDefault,
			publisherSearchIdxColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			publisherSearchIdxAllColumns,
			publisherSearchIdxPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert publisher_search_idx, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(publisherSearchIdxPrimaryKeyColumns))
			copy(conflict, publisherSearchIdxPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"publisher_search_idx\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(publisherSearchIdxType, publisherSearchIdxMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(publisherSearchIdxType, publisherSearchIdxMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert publisher_search_idx")
	}

	if !cached {
		publisherSearchIdxUpsertCacheMut.Lock()
		publisherSearchIdxUpsertCache[key] = cache
		publisherSearchIdxUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
