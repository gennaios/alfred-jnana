// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Publisher is an object representing the database table.
type Publisher struct {
	ID   null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	Name string     `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *publisherR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L publisherL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PublisherColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

var PublisherTableColumns = struct {
	ID   string
	Name string
}{
	ID:   "publisher.id",
	Name: "publisher.name",
}

// Generated where

var PublisherWhere = struct {
	ID   whereHelpernull_Int64
	Name whereHelperstring
}{
	ID:   whereHelpernull_Int64{field: "\"publisher\".\"id\""},
	Name: whereHelperstring{field: "\"publisher\".\"name\""},
}

// PublisherRels is where relationship names are stored.
var PublisherRels = struct {
}{}

// publisherR is where relationships are stored.
type publisherR struct {
}

// NewStruct creates a new relationship struct
func (*publisherR) NewStruct() *publisherR {
	return &publisherR{}
}

// publisherL is where Load methods for each relationship are stored.
type publisherL struct{}

var (
	publisherAllColumns            = []string{"id", "name"}
	publisherColumnsWithoutDefault = []string{}
	publisherColumnsWithDefault    = []string{"id", "name"}
	publisherPrimaryKeyColumns     = []string{"id"}
)

type (
	// PublisherSlice is an alias for a slice of pointers to Publisher.
	// This should almost always be used instead of []Publisher.
	PublisherSlice []*Publisher
	// PublisherHook is the signature for custom Publisher hook methods
	PublisherHook func(context.Context, boil.ContextExecutor, *Publisher) error

	publisherQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	publisherType                 = reflect.TypeOf(&Publisher{})
	publisherMapping              = queries.MakeStructMapping(publisherType)
	publisherPrimaryKeyMapping, _ = queries.BindMapping(publisherType, publisherMapping, publisherPrimaryKeyColumns)
	publisherInsertCacheMut       sync.RWMutex
	publisherInsertCache          = make(map[string]insertCache)
	publisherUpdateCacheMut       sync.RWMutex
	publisherUpdateCache          = make(map[string]updateCache)
	publisherUpsertCacheMut       sync.RWMutex
	publisherUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var publisherBeforeInsertHooks []PublisherHook
var publisherBeforeUpdateHooks []PublisherHook
var publisherBeforeDeleteHooks []PublisherHook
var publisherBeforeUpsertHooks []PublisherHook

var publisherAfterInsertHooks []PublisherHook
var publisherAfterSelectHooks []PublisherHook
var publisherAfterUpdateHooks []PublisherHook
var publisherAfterDeleteHooks []PublisherHook
var publisherAfterUpsertHooks []PublisherHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Publisher) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Publisher) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Publisher) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Publisher) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Publisher) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Publisher) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Publisher) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Publisher) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Publisher) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPublisherHook registers your hook function for all future operations.
func AddPublisherHook(hookPoint boil.HookPoint, publisherHook PublisherHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		publisherBeforeInsertHooks = append(publisherBeforeInsertHooks, publisherHook)
	case boil.BeforeUpdateHook:
		publisherBeforeUpdateHooks = append(publisherBeforeUpdateHooks, publisherHook)
	case boil.BeforeDeleteHook:
		publisherBeforeDeleteHooks = append(publisherBeforeDeleteHooks, publisherHook)
	case boil.BeforeUpsertHook:
		publisherBeforeUpsertHooks = append(publisherBeforeUpsertHooks, publisherHook)
	case boil.AfterInsertHook:
		publisherAfterInsertHooks = append(publisherAfterInsertHooks, publisherHook)
	case boil.AfterSelectHook:
		publisherAfterSelectHooks = append(publisherAfterSelectHooks, publisherHook)
	case boil.AfterUpdateHook:
		publisherAfterUpdateHooks = append(publisherAfterUpdateHooks, publisherHook)
	case boil.AfterDeleteHook:
		publisherAfterDeleteHooks = append(publisherAfterDeleteHooks, publisherHook)
	case boil.AfterUpsertHook:
		publisherAfterUpsertHooks = append(publisherAfterUpsertHooks, publisherHook)
	}
}

// One returns a single publisher record from the query.
func (q publisherQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Publisher, error) {
	o := &Publisher{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for publisher")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Publisher records from the query.
func (q publisherQuery) All(ctx context.Context, exec boil.ContextExecutor) (PublisherSlice, error) {
	var o []*Publisher

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Publisher slice")
	}

	if len(publisherAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Publisher records in the query.
func (q publisherQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count publisher rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q publisherQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if publisher exists")
	}

	return count > 0, nil
}

// Publishers retrieves all the records using an executor.
func Publishers(mods ...qm.QueryMod) publisherQuery {
	mods = append(mods, qm.From("\"publisher\""))
	return publisherQuery{NewQuery(mods...)}
}

// FindPublisher retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPublisher(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*Publisher, error) {
	publisherObj := &Publisher{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"publisher\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, publisherObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from publisher")
	}

	if err = publisherObj.doAfterSelectHooks(ctx, exec); err != nil {
		return publisherObj, err
	}

	return publisherObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Publisher) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publisher provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(publisherColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	publisherInsertCacheMut.RLock()
	cache, cached := publisherInsertCache[key]
	publisherInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			publisherAllColumns,
			publisherColumnsWithDefault,
			publisherColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(publisherType, publisherMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(publisherType, publisherMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"publisher\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"publisher\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into publisher")
	}

	if !cached {
		publisherInsertCacheMut.Lock()
		publisherInsertCache[key] = cache
		publisherInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Publisher.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Publisher) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	publisherUpdateCacheMut.RLock()
	cache, cached := publisherUpdateCache[key]
	publisherUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			publisherAllColumns,
			publisherPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update publisher, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"publisher\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, publisherPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(publisherType, publisherMapping, append(wl, publisherPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update publisher row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for publisher")
	}

	if !cached {
		publisherUpdateCacheMut.Lock()
		publisherUpdateCache[key] = cache
		publisherUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q publisherQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for publisher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for publisher")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PublisherSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"publisher\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, publisherPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in publisher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all publisher")
	}
	return rowsAff, nil
}

// Delete deletes a single Publisher record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Publisher) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Publisher provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), publisherPrimaryKeyMapping)
	sql := "DELETE FROM \"publisher\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from publisher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for publisher")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q publisherQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no publisherQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publisher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publisher")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PublisherSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(publisherBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"publisher\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, publisherPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publisher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publisher")
	}

	if len(publisherAfterDeleteHooks) != 0 {
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
func (o *Publisher) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPublisher(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PublisherSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PublisherSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"publisher\".* FROM \"publisher\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, publisherPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PublisherSlice")
	}

	*o = slice

	return nil
}

// PublisherExists checks if the Publisher row exists.
func PublisherExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"publisher\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if publisher exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Publisher) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publisher provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(publisherColumnsWithDefault, o)

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

	publisherUpsertCacheMut.RLock()
	cache, cached := publisherUpsertCache[key]
	publisherUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			publisherAllColumns,
			publisherColumnsWithDefault,
			publisherColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			publisherAllColumns,
			publisherPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert publisher, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(publisherPrimaryKeyColumns))
			copy(conflict, publisherPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"publisher\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(publisherType, publisherMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(publisherType, publisherMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert publisher")
	}

	if !cached {
		publisherUpsertCacheMut.Lock()
		publisherUpsertCache[key] = cache
		publisherUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}