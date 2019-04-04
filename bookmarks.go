package main

import (
	"github.com/deckarep/gosx-notifier"
	"github.com/gocraft/dbr"
	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
)

type Database struct {
	conn *dbr.Connection
	sess *dbr.Session
}

type BookmarkRecord struct {
	ID          int64
	FileId      int64          `db:"file_id"`
	Title       string         `db:"title"`
	Section     dbr.NullString `db:"section"`
	Destination string         `db:"destination"`
}

// TODO: combine with above
type SearchAllResult struct {
	ID          int64
	Title       string         `db:"title"`
	Section     dbr.NullString `db:"section"`
	Destination string         `db:"destination"`
	FileID      string         `db:"file_id"`
	Path        string         `db:"path"`
	FileName    string         `db:"file_name"`
}

func (db *Database) Init(filepath string) {
	conn, err := dbr.Open("sqlite3", filepath, nil)
	if err != nil {
		panic(err)
	}
	if conn == nil {
		panic("db nil")
	}
	db.conn = conn

	_, _ = conn.Exec("PRAGMA auto_vacuum = 1")
	_, _ = conn.Exec("PRAGMA foreign_keys = 1")
	_, _ = conn.Exec("PRAGMA ignore_check_constraints = 0")
	_, _ = conn.Exec("PRAGMA journal_mode = WAL")
	_, _ = conn.Exec("PRAGMA synchronous = 0")
	_, _ = conn.Exec("PRAGMA temp_store = 2") // MEMORY

	_, _ = conn.Exec("PRAGMA cache_size = -31250")
	_, _ = conn.Exec("PRAGMA page_size = 8192") // default 4096, match APFS block size?

	sess := conn.NewSession(nil)
	db.sess = sess
	_, err = sess.Begin()
	if err != nil {
		panic(err)
	}
}

func (db *Database) BookmarksForFile(file string) ([]BookmarkRecord, error) {
	var bookmarks []BookmarkRecord
	var err error

	fileRecord, changed, _ := db.GetFile(file)
	err = db.sess.Select("id", "title", "section", "destination").From("bookmarks").
		Where("file_id == ?", fileRecord.ID).
		LoadOne(&bookmarks)

	// TODO, file created or changed, compared bookmarks
	if changed == true {
		var newBookmarks []Bookmark

		// PDF get bookmarks from file
		if strings.HasSuffix(file, ".pdf") {
			// TODO: call pdf module
			newBookmarks, _ = bookmarksForPDF(file)
		}
		// TODO: EPUB get bookmarks from file

		// file updated, compare bookmarks
		if bookmarksEqual(bookmarks, newBookmarks) == false {
			// update database
			err = db.UpdateBookmarks(fileRecord, newBookmarks)
			_ = notification("Bookmarks updated.")
			// fetch new bookmarks
		}
	}
	return bookmarks, err
}

// Filtered bookmarks for file
func (db *Database) BookmarksForFileFiltered(file string, query string) ([]SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []SearchAllResult

	fileRecord, _, _ := db.GetFile(file)

	/*
		select bookmarks.id, bookmarks.title, bookmarks.section, bookmarks.destination
		from bookmarks
		JOIN bookmarksindex on bookmarks.id = bookmarksindex.rowid
		where bookmarks.file_id = ? and bookmarksindex.rowid = bookmarks.id
		and bookmarksindex match '{title section} : …' ORDER BY 'bm25(bookmarksindex, 5.0, 2.0)';
	*/

	//_, err := sess.Select("bookmarks.id", "bookmarks.title", "bookmarks.section",
	//	"bookmarks.destination").
	//	From("bookmarks").
	//	Join("bookmarksindex", "bookmarks.id = bookmarksindex.rowid").
	//	Where("bookmarks.file_id = ?", fileRecord.ID).
	//	Where("bookmarksindex MATCH ? ORDER BY 'bm25(bookmarksindex, 5.0, 2.0)'", queryString).
	//	Load(&results)

	// TODO: limit column matches not working
	_, err := db.sess.SelectBySql("select bookmarks.id, bookmarks.title, bookmarks.section, bookmarks.destination from bookmarks JOIN bookmarksindex on bookmarks.id = bookmarksindex.rowid where bookmarks.file_id = " + strconv.FormatInt(fileRecord.ID, 10) + " and bookmarksindex.rowid = bookmarks.id and bookmarksindex match '{title section} : " + queryString + "' ORDER BY 'bm25(bookmarksindex, 5.0, 2.0)';").Load(&results)
	return results, err
}

// Search all bookmarks from FTS5 table, order by rank title, section, & file name
// Return results as slice of struct SearchAllResult, later prepped for Alfred script filter
func (db *Database) searchAll(query string) ([]SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []SearchAllResult

	//SELECT
	//bookmarks.id, bookmarks.title, bookmarks.section, bookmarks.destination,
	//bookmarks.file_id, files.path, files.file_name
	//FROM bookmarks
	//JOIN files ON bookmarks.file_id = files.id
	//JOIN bookmarksindex on bookmarks.id = bookmarksindex.rowid
	//WHERE bookmarksindex MATCH '?' AND rank MATCH 'bm25(5.0, 2.0, 1.0)'

	// NOTE: AND rank MATCH 'bm25(10.0, 5.0)' ORDER BY rank faster than ORDER BY bm25(fts, …)
	_, err := db.sess.Select("bookmarks.id", "bookmarks.title", "bookmarks.section",
		"bookmarks.destination", "bookmarks.file_id", "files.path", "files.file_name").
		From("bookmarks").
		Join("files", "bookmarks.file_id = files.id").
		Join("bookmarksindex", "bookmarks.id = bookmarksindex.rowid").
		Where("bookmarksindex MATCH ? AND rank MATCH 'bm25(5.0, 2.0, 1.0)'", queryString).
		OrderBy("rank").Limit(100).Load(&results)
	return results, err
}

func (db *Database) UpdateBookmarks(file File, bookmarks []Bookmark) error {
	tx, err := db.sess.Begin()
	_, err = db.sess.DeleteFrom("bookmarks").
		Where("file_id = ?", file.ID).
		Exec()
	// insert new bookmarks
	for i := 0; i < len(bookmarks); i++ {
		_, err = db.sess.InsertInto("bookmarks").
			Pair("file_id", file.ID).
			Pair("title", bookmarks[i].Title).
			Pair("section", dbr.NewNullString(bookmarks[i].Section)).
			Pair("destination", bookmarks[i].Destination).
			Exec()
	}
	err = tx.Commit()
	return err
}

// Compare bookmarks from database with file
func bookmarksEqual(bookmarks []BookmarkRecord, newBookmarks []Bookmark) bool {
	var oldBookmarks []Bookmark

	for i := 0; i < len(bookmarks); i++ {
		var bookmark Bookmark
		bookmark.Title = bookmarks[i].Title
		bookmark.Section = bookmarks[i].Section.String
		bookmark.Destination = bookmarks[i].Destination
		oldBookmarks = append(oldBookmarks, bookmark)
	}

	if cmp.Equal(oldBookmarks, newBookmarks) {
		return true
	} else {
		return false
	}
}

// Prepare string for SQLite FTS query
// replace '–*' with 'NOT *'
func stringForSQLite(query string) string {
	var queryArray []string
	queryOperators := []string{"AND", "OR", "NOT", "and", "or", "not"}

	slc := strings.Split(query, " ")
	for i := range slc {
		term := slc[i]

		if strings.HasPrefix(term, "-") {
			// exclude terms beginning with '-', change to 'NOT [term]'
			queryArray = append(queryArray, "NOT "+term[1:]+"*")
		} else if stringInSlice(term, queryOperators) {
			// auto capitalize operators 'and', 'or', 'not'
			queryArray = append(queryArray, strings.ToUpper(term))
		} else {
			queryArray = append(queryArray, term+"*")
		}
	}
	return strings.TrimSpace(strings.Join(queryArray[:], " "))
}

func notification(message string) error {
	note := gosxnotifier.NewNotification(message)
	note.Title = "Jnana"
	note.Sound = gosxnotifier.Default

	if err := note.Push(); err != nil {
		return err
	}
	return nil
}

// Test if string is included in slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
