package bookmarks

import (
	"github.com/volatiletech/sqlboiler/v4/queries"
	"jnana/internal/database"
	bookFile "jnana/internal/file"
	"jnana/internal/files"
	"jnana/internal/fulltext"
	"jnana/internal/util"
	"jnana/models"

	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/null/v8"
	"strconv"
)

type SearchAllResult struct {
	ID          int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Section     null.String `boil:"section" json:"section,omitempty" toml:"section" yaml:"section,omitempty"`
	Destination string      `boil:"destination" json:"destination" toml:"destination" yaml:"hash"`
	FileID      string      `boil:"file_id" json:"file_id" toml:"file_id" yaml:"file_id"`
	Path        string      `boil:"path" json:"path" toml:"path" yaml:"path"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
}

// ForFile retrieve existing bookmarks, add new to database if needed and check if updated
func ForFile(db *database.Database, file string, coversCacheDir string) ([]*models.Bookmark, error) {
	var bookmarks []*models.Bookmark
	var err error

	fileRecord, changed, err := files.Get(db, file, true)
	if err != nil {
		return bookmarks, err
	}

	err = queries.Raw(`SELECT id, title, section, destination FROM bookmark
					WHERE file_id = $1`, fileRecord.ID).Bind(db.Ctx, db.Db, &bookmarks)

	// check cover / TODO: move somewhere else?
	//_ = files.CoverForFile(fileRecord, coversCacheDir)

	// file created or changed / or no bookmarks found
	if changed == true || len(bookmarks) == 0 {
		var newBookmarks []*bookFile.Bookmark

		f := bookFile.File{}
		if err = f.Init(file); err != nil {
			return bookmarks, err
		}
		newBookmarks, _ = f.Bookmarks()

		// no bookmarks returned from first, get new
		if len(bookmarks) == 0 {
			// insert new
			bookmarks, err = New(db, fileRecord, newBookmarks)
		} else if bookmarksEqual(bookmarks, newBookmarks) == false {
			// file updated, compare bookmarks
			// update database
			bookmarks, err = Update(db, fileRecord, bookmarks, newBookmarks)
			_ = util.Notification("Bookmarks updated.")
		}
	}

	// run analyze etc upon database close, unsure if faster
	//_, _ = db.conn.Exec("PRAGMA optimize;")
	err = db.Db.Close()
	return bookmarks, err
}

// ForFileFiltered filtered bookmarks for file, uses fileId from elsewhere so there's only one query
func ForFileFiltered(db *database.Database, file string, query string) ([]*SearchAllResult, error) {
	queryString := util.StringForSQLite(query)
	var results []*SearchAllResult

	// only ID needed, no additional fields or checks
	var fileRecord *models.File

	fileRecord, _ = models.Files(models.FileWhere.Path.EQ(file)).One(db.Ctx, db.Db)

	err := queries.Raw(`SELECT
			bookmark.id, bookmark.title, bookmark.section, bookmark.destination
			FROM bookmark
			JOIN bookmark_search on bookmark.id = bookmark_search.rowid `+
		`WHERE bookmark.file_id = `+strconv.FormatInt(fileRecord.ID, 10)+
		` AND bookmark_search MATCH '{title section}: `+*queryString+"' "+
		`ORDER BY 'rank(bookmark_search)'`).Bind(db.Ctx, db.Db, &results)

	// run analyze etc upon database close, unsure if faster
	//_, _ = db.conn.Exec("PRAGMA optimize;")
	_ = db.Db.Close()
	return results, err
}

// SearchAll bookmarks from FTS5 table, order by rank title, section, & file name
// Return results as slice of struct SearchAllResult, later prepped for Alfred script filter
func SearchAll(db *database.Database, query string) ([]*SearchAllResult, error) {
	queryString := util.StringForSQLite(query)
	var results []*SearchAllResult

	// NOTE: AND rank MATCH 'bm25(10.0, 5.0)' ORDER BY rank faster than ORDER BY bm25(fts, …)
	err := queries.Raw(`SELECT
			bookmark.id, bookmark.title, bookmark.section, bookmark.destination,
			bookmark.file_id, file.path, file.name
			FROM bookmark
			JOIN file ON bookmark.file_id = file.id
			JOIN bookmark_search on bookmark.id = bookmark_search.rowid
			WHERE bookmark_search MATCH $1
			AND rank MATCH 'bm25(10.0, 5.0, 2.0, 1.0, 1.0, 1.0, 1.0)'
			ORDER BY rank LIMIT 200`,
		queryString).Bind(db.Ctx, db.Db, &results)

	_ = db.Db.Close()
	return results, err
}

// New insert new bookmarks into database
func New(db *database.Database, file *models.File, bookmarks []*bookFile.Bookmark) ([]*models.Bookmark, error) {
	tx, err := db.Db.BeginTx(db.Ctx, nil)

	// insert new bookmarks
	for i := range bookmarks {
		_, err = queries.Raw(`INSERT INTO
			bookmark
			(file_id, title, section, destination) VALUES ($1, $2, $3, $4)
			`,
			file.ID, bookmarks[i].Title, bookmarks[i].Section, bookmarks[i].Destination).
			Exec(db.Db)
	}

	fulltext.BookmarksCreate(db, *file)

	err = tx.Commit()

	// get newly inserted bookmarks
	var newBookmarks []*models.Bookmark
	err = queries.Raw(`SELECT id, title, section, destination FROM bookmark
		WHERE file_id = $1`, file.ID).BindG(db.Ctx, &newBookmarks)
	return newBookmarks, err
}

// Update bookmarks, delete old first, then call New
func Update(db *database.Database, file *models.File, oldBookmarks []*models.Bookmark, newBookmarks []*bookFile.Bookmark) ([]*models.Bookmark, error) {
	var err error
	var results []*models.Bookmark

	tx, _ := db.Db.BeginTx(db.Ctx, nil)

	fulltext.BookmarksDelete(db, *file)

	if len(oldBookmarks) == len(newBookmarks) {
		// count same, update records
		for i := range oldBookmarks {
			_, err = queries.Raw(`UPDATE bookmark SET
				title = $1, section = $2, destination = $3
				WHERE id = $4`,
				newBookmarks[i].Title,
				newBookmarks[i].Section,
				newBookmarks[i].Destination, oldBookmarks[i].ID).
				Exec(db.Db)
		}
		err = queries.Raw(`SELECT id, title, section, destination FROM bookmark
			WHERE file_id = $1`, file.ID).BindG(db.Ctx, &results)
	} else {
		// count different, delete and insert new
		_, err = queries.Raw(`DELETE from bookmark WHERE file_id = $1`, file.ID).Exec(db.Db)
		results, err = New(db, file, newBookmarks)
	}

	fulltext.BookmarksCreate(db, *file)

	err = tx.Commit()

	return results, err
}

// bookmarksEqual compare bookmarks from database with path, used for update check
func bookmarksEqual(bookmarks []*models.Bookmark, newBookmarks []*bookFile.Bookmark) bool {
	if len(newBookmarks) != len(bookmarks) {
		return false
	} else {
		for i := range newBookmarks {
			if newBookmarks[i].Title != bookmarks[i].Title.String {
				return false
			}
			if newBookmarks[i].Section != bookmarks[i].Section.String {
				return false
			}
			if newBookmarks[i].Destination != bookmarks[i].Destination {
				return false
			}
		}
	}
	return true
}
