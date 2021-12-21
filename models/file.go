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

// File is an object representing the database table.
type File struct {
	ID           int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Path         string      `boil:"path" json:"path" toml:"path" yaml:"path"`
	Name         string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Extension    string      `boil:"extension" json:"extension" toml:"extension" yaml:"extension"`
	Size         int64       `boil:"size" json:"size" toml:"size" yaml:"size"`
	Title        null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Publisher    null.String `boil:"publisher" json:"publisher,omitempty" toml:"publisher" yaml:"publisher,omitempty"`
	PublisherID  null.Int64  `boil:"publisher_id" json:"publisher_id,omitempty" toml:"publisher_id" yaml:"publisher_id,omitempty"`
	Creator      null.String `boil:"creator" json:"creator,omitempty" toml:"creator" yaml:"creator,omitempty"`
	Subject      null.String `boil:"subject" json:"subject,omitempty" toml:"subject" yaml:"subject,omitempty"`
	Language     null.String `boil:"language" json:"language,omitempty" toml:"language" yaml:"language,omitempty"`
	Description  null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	DateCreated  time.Time   `boil:"date_created" json:"date_created" toml:"date_created" yaml:"date_created"`
	DateModified time.Time   `boil:"date_modified" json:"date_modified" toml:"date_modified" yaml:"date_modified"`
	DateAccessed null.Time   `boil:"date_accessed" json:"date_accessed,omitempty" toml:"date_accessed" yaml:"date_accessed,omitempty"`
	Rating       null.Int64  `boil:"rating" json:"rating,omitempty" toml:"rating" yaml:"rating,omitempty"`
	Hash         string      `boil:"hash" json:"hash" toml:"hash" yaml:"hash"`

	R *fileR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L fileL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FileColumns = struct {
	ID           string
	Path         string
	Name         string
	Extension    string
	Size         string
	Title        string
	Publisher    string
	PublisherID  string
	Creator      string
	Subject      string
	Language     string
	Description  string
	DateCreated  string
	DateModified string
	DateAccessed string
	Rating       string
	Hash         string
}{
	ID:           "id",
	Path:         "path",
	Name:         "name",
	Extension:    "extension",
	Size:         "size",
	Title:        "title",
	Publisher:    "publisher",
	PublisherID:  "publisher_id",
	Creator:      "creator",
	Subject:      "subject",
	Language:     "language",
	Description:  "description",
	DateCreated:  "date_created",
	DateModified: "date_modified",
	DateAccessed: "date_accessed",
	Rating:       "rating",
	Hash:         "hash",
}

var FileTableColumns = struct {
	ID           string
	Path         string
	Name         string
	Extension    string
	Size         string
	Title        string
	Publisher    string
	PublisherID  string
	Creator      string
	Subject      string
	Language     string
	Description  string
	DateCreated  string
	DateModified string
	DateAccessed string
	Rating       string
	Hash         string
}{
	ID:           "file.id",
	Path:         "file.path",
	Name:         "file.name",
	Extension:    "file.extension",
	Size:         "file.size",
	Title:        "file.title",
	Publisher:    "file.publisher",
	PublisherID:  "file.publisher_id",
	Creator:      "file.creator",
	Subject:      "file.subject",
	Language:     "file.language",
	Description:  "file.description",
	DateCreated:  "file.date_created",
	DateModified: "file.date_modified",
	DateAccessed: "file.date_accessed",
	Rating:       "file.rating",
	Hash:         "file.hash",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var FileWhere = struct {
	ID           whereHelperint64
	Path         whereHelperstring
	Name         whereHelperstring
	Extension    whereHelperstring
	Size         whereHelperint64
	Title        whereHelpernull_String
	Publisher    whereHelpernull_String
	PublisherID  whereHelpernull_Int64
	Creator      whereHelpernull_String
	Subject      whereHelpernull_String
	Language     whereHelpernull_String
	Description  whereHelpernull_String
	DateCreated  whereHelpertime_Time
	DateModified whereHelpertime_Time
	DateAccessed whereHelpernull_Time
	Rating       whereHelpernull_Int64
	Hash         whereHelperstring
}{
	ID:           whereHelperint64{field: "\"file\".\"id\""},
	Path:         whereHelperstring{field: "\"file\".\"path\""},
	Name:         whereHelperstring{field: "\"file\".\"name\""},
	Extension:    whereHelperstring{field: "\"file\".\"extension\""},
	Size:         whereHelperint64{field: "\"file\".\"size\""},
	Title:        whereHelpernull_String{field: "\"file\".\"title\""},
	Publisher:    whereHelpernull_String{field: "\"file\".\"publisher\""},
	PublisherID:  whereHelpernull_Int64{field: "\"file\".\"publisher_id\""},
	Creator:      whereHelpernull_String{field: "\"file\".\"creator\""},
	Subject:      whereHelpernull_String{field: "\"file\".\"subject\""},
	Language:     whereHelpernull_String{field: "\"file\".\"language\""},
	Description:  whereHelpernull_String{field: "\"file\".\"description\""},
	DateCreated:  whereHelpertime_Time{field: "\"file\".\"date_created\""},
	DateModified: whereHelpertime_Time{field: "\"file\".\"date_modified\""},
	DateAccessed: whereHelpernull_Time{field: "\"file\".\"date_accessed\""},
	Rating:       whereHelpernull_Int64{field: "\"file\".\"rating\""},
	Hash:         whereHelperstring{field: "\"file\".\"hash\""},
}

// FileRels is where relationship names are stored.
var FileRels = struct {
	Bookmarks string
}{
	Bookmarks: "Bookmarks",
}

// fileR is where relationships are stored.
type fileR struct {
	Bookmarks BookmarkSlice `boil:"Bookmarks" json:"Bookmarks" toml:"Bookmarks" yaml:"Bookmarks"`
}

// NewStruct creates a new relationship struct
func (*fileR) NewStruct() *fileR {
	return &fileR{}
}

// fileL is where Load methods for each relationship are stored.
type fileL struct{}

var (
	fileAllColumns            = []string{"id", "path", "name", "extension", "size", "title", "publisher", "publisher_id", "creator", "subject", "language", "description", "date_created", "date_modified", "date_accessed", "rating", "hash"}
	fileColumnsWithoutDefault = []string{"path", "name", "extension", "size", "title", "publisher", "publisher_id", "creator", "subject", "language", "description", "date_created", "date_modified", "date_accessed", "rating", "hash"}
	fileColumnsWithDefault    = []string{"id"}
	filePrimaryKeyColumns     = []string{"id"}
)

type (
	// FileSlice is an alias for a slice of pointers to File.
	// This should almost always be used instead of []File.
	FileSlice []*File
	// FileHook is the signature for custom File hook methods
	FileHook func(context.Context, boil.ContextExecutor, *File) error

	fileQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	fileType                 = reflect.TypeOf(&File{})
	fileMapping              = queries.MakeStructMapping(fileType)
	filePrimaryKeyMapping, _ = queries.BindMapping(fileType, fileMapping, filePrimaryKeyColumns)
	fileInsertCacheMut       sync.RWMutex
	fileInsertCache          = make(map[string]insertCache)
	fileUpdateCacheMut       sync.RWMutex
	fileUpdateCache          = make(map[string]updateCache)
	fileUpsertCacheMut       sync.RWMutex
	fileUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var fileBeforeInsertHooks []FileHook
var fileBeforeUpdateHooks []FileHook
var fileBeforeDeleteHooks []FileHook
var fileBeforeUpsertHooks []FileHook

var fileAfterInsertHooks []FileHook
var fileAfterSelectHooks []FileHook
var fileAfterUpdateHooks []FileHook
var fileAfterDeleteHooks []FileHook
var fileAfterUpsertHooks []FileHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *File) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *File) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *File) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *File) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *File) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *File) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *File) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *File) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *File) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFileHook registers your hook function for all future operations.
func AddFileHook(hookPoint boil.HookPoint, fileHook FileHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		fileBeforeInsertHooks = append(fileBeforeInsertHooks, fileHook)
	case boil.BeforeUpdateHook:
		fileBeforeUpdateHooks = append(fileBeforeUpdateHooks, fileHook)
	case boil.BeforeDeleteHook:
		fileBeforeDeleteHooks = append(fileBeforeDeleteHooks, fileHook)
	case boil.BeforeUpsertHook:
		fileBeforeUpsertHooks = append(fileBeforeUpsertHooks, fileHook)
	case boil.AfterInsertHook:
		fileAfterInsertHooks = append(fileAfterInsertHooks, fileHook)
	case boil.AfterSelectHook:
		fileAfterSelectHooks = append(fileAfterSelectHooks, fileHook)
	case boil.AfterUpdateHook:
		fileAfterUpdateHooks = append(fileAfterUpdateHooks, fileHook)
	case boil.AfterDeleteHook:
		fileAfterDeleteHooks = append(fileAfterDeleteHooks, fileHook)
	case boil.AfterUpsertHook:
		fileAfterUpsertHooks = append(fileAfterUpsertHooks, fileHook)
	}
}

// One returns a single file record from the query.
func (q fileQuery) One(ctx context.Context, exec boil.ContextExecutor) (*File, error) {
	o := &File{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for file")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all File records from the query.
func (q fileQuery) All(ctx context.Context, exec boil.ContextExecutor) (FileSlice, error) {
	var o []*File

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to File slice")
	}

	if len(fileAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all File records in the query.
func (q fileQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count file rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q fileQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if file exists")
	}

	return count > 0, nil
}

// Bookmarks retrieves all the bookmark's Bookmarks with an executor.
func (o *File) Bookmarks(mods ...qm.QueryMod) bookmarkQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"bookmark\".\"file_id\"=?", o.ID),
	)

	query := Bookmarks(queryMods...)
	queries.SetFrom(query.Query, "\"bookmark\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"bookmark\".*"})
	}

	return query
}

// LoadBookmarks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (fileL) LoadBookmarks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFile interface{}, mods queries.Applicator) error {
	var slice []*File
	var object *File

	if singular {
		object = maybeFile.(*File)
	} else {
		slice = *maybeFile.(*[]*File)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &fileR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &fileR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`bookmark`),
		qm.WhereIn(`bookmark.file_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load bookmark")
	}

	var resultSlice []*Bookmark
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice bookmark")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on bookmark")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bookmark")
	}

	if len(bookmarkAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Bookmarks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &bookmarkR{}
			}
			foreign.R.File = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.FileID {
				local.R.Bookmarks = append(local.R.Bookmarks, foreign)
				if foreign.R == nil {
					foreign.R = &bookmarkR{}
				}
				foreign.R.File = local
				break
			}
		}
	}

	return nil
}

// AddBookmarks adds the given related objects to the existing relationships
// of the file, optionally inserting them as new records.
// Appends related to o.R.Bookmarks.
// Sets related.R.File appropriately.
func (o *File) AddBookmarks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Bookmark) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FileID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"bookmark\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"file_id"}),
				strmangle.WhereClause("\"", "\"", 0, bookmarkPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FileID = o.ID
		}
	}

	if o.R == nil {
		o.R = &fileR{
			Bookmarks: related,
		}
	} else {
		o.R.Bookmarks = append(o.R.Bookmarks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &bookmarkR{
				File: o,
			}
		} else {
			rel.R.File = o
		}
	}
	return nil
}

// Files retrieves all the records using an executor.
func Files(mods ...qm.QueryMod) fileQuery {
	mods = append(mods, qm.From("\"file\""))
	return fileQuery{NewQuery(mods...)}
}

// FindFile retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFile(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*File, error) {
	fileObj := &File{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"file\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, fileObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from file")
	}

	if err = fileObj.doAfterSelectHooks(ctx, exec); err != nil {
		return fileObj, err
	}

	return fileObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *File) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no file provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fileColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	fileInsertCacheMut.RLock()
	cache, cached := fileInsertCache[key]
	fileInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			fileAllColumns,
			fileColumnsWithDefault,
			fileColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(fileType, fileMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(fileType, fileMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"file\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"file\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into file")
	}

	if !cached {
		fileInsertCacheMut.Lock()
		fileInsertCache[key] = cache
		fileInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the File.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *File) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	fileUpdateCacheMut.RLock()
	cache, cached := fileUpdateCache[key]
	fileUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			fileAllColumns,
			filePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update file, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"file\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, filePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(fileType, fileMapping, append(wl, filePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update file row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for file")
	}

	if !cached {
		fileUpdateCacheMut.Lock()
		fileUpdateCache[key] = cache
		fileUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q fileQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for file")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for file")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FileSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"file\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in file slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all file")
	}
	return rowsAff, nil
}

// Delete deletes a single File record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *File) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no File provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), filePrimaryKeyMapping)
	sql := "DELETE FROM \"file\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from file")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for file")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q fileQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no fileQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from file")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for file")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FileSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(fileBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"file\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from file slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for file")
	}

	if len(fileAfterDeleteHooks) != 0 {
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
func (o *File) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFile(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FileSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FileSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"file\".* FROM \"file\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FileSlice")
	}

	*o = slice

	return nil
}

// FileExists checks if the File row exists.
func FileExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"file\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if file exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *File) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no file provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fileColumnsWithDefault, o)

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

	fileUpsertCacheMut.RLock()
	cache, cached := fileUpsertCache[key]
	fileUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			fileAllColumns,
			fileColumnsWithDefault,
			fileColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			fileAllColumns,
			filePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert file, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(filePrimaryKeyColumns))
			copy(conflict, filePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"file\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(fileType, fileMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(fileType, fileMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert file")
	}

	if !cached {
		fileUpsertCacheMut.Lock()
		fileUpsertCache[key] = cache
		fileUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}