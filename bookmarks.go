package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/deckarep/gosx-notifier"
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	conn *dbr.Connection
	sess *dbr.Session
}

type Bookmark struct {
	ID          int64
	FileId      int64          `db:"file_id"`
	Title       string         `db:"title"`
	Section     dbr.NullString `db:"section"`
	Destination string         `db:"destination"`
}

type SearchAllResult struct {
	ID          int64
	Title       string         `db:"title"`
	Section     dbr.NullString `db:"section"`
	Destination string         `db:"destination"`
	FileID      string         `db:"file_id"`
	Path        string         `db:"path"`
	FileName    string         `db:"file_name"`
}

// Init: open SQLite database connection using dbr, create new session
func (db *Database) Init(dbFilePath string) {
	var file string
	// open with PRAGMAs:
	// journal_mode=WAL, locking_mode=EXCLUSIVE, synchronous=0
	if dbFilePath != "memory" {
		file = dbFilePath + "?&_journal_mode=WAL&_locking_mode=EXCLUSIVE&_synchronous=0&_foreign_keys=1"
	} else {
		file = "file::memory:?mode=memory&cache=shared&_foreign_keys=1"
	}

	var err error
	db.conn, err = dbr.Open("sqlite3", file, nil)

	// TODO: return error
	if err != nil {
		panic(err)
	}
	if db.conn == nil {
		panic("db nil")
	}

	_, err = db.conn.Exec("PRAGMA auto_vacuum=2") // unsure if set
	//_, _ = db.conn.Exec("PRAGMA temp_store = 2")  // MEMORY
	//_, _ = db.conn.Exec("PRAGMA cache_size = -31250")

	db.sess = db.conn.NewSession(nil)
	_, err = db.sess.Begin()

	// create tables and triggers if 'files' does not exist
	tables := db.sess.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='files'")
	var result string
	_ = tables.Scan(&result)
	if result != "files" {
		db.createTables()
		db.createTriggers()
	}

	if err != nil {
		panic(err)
	}
}

// Init: open SQLite database connection using dbr, for reading, doesn't create tables etc
func (db *Database) InitForReading(dbFilePath string) {
	db.conn, _ = dbr.Open("sqlite3", dbFilePath, nil)
	db.sess = db.conn.NewSession(nil)
	_, _ = db.sess.Begin()
}

// createTables: create tables files, bookmarks, view bookmarks_view for FTS updates, and FTS5 bookmarksindex
func (db *Database) createTables() {
	var schemaFiles = `
	CREATE TABLE files (
		id INTEGER NOT NULL PRIMARY KEY,
		path TEXT NOT NULL,
	    	file_name TEXT,
	    	file_extension VARCHAR(255) NOT NULL,
	    	file_title TEXT,
	    	file_authors TEXT,
	    	file_subjects TEXT,
	    	file_publisher TEXT,
	    	date_created DATETIME NOT NULL,
	    	date_modified DATETIME,
	    	hash VARCHAR(64) NOT NULL
	)`
	var schemaBookmarks = `
	CREATE TABLE bookmarks (
		id INTEGER NOT NULL PRIMARY KEY,
		file_id INTEGER NOT NULL,
		title TEXT,
		section TEXT,
		destination TEXT NOT NULL,
		FOREIGN KEY (file_id) REFERENCES files (id) ON DELETE CASCADE ON UPDATE CASCADE
	)`
	var schemaView = `
	CREATE VIEW bookmarks_view AS SELECT
		bookmarks.id,
		bookmarks.title,
		bookmarks.section,
		files.file_name,
		files.file_title,
		files.file_authors,
		files.file_subjects,
		files.file_publisher
		FROM bookmarks
		INNER JOIN files ON files.id = bookmarks.file_id
	`
	// prefix: tokenize by length
	var schemaFTS = `
	CREATE VIRTUAL TABLE bookmarksindex USING fts5(
		title,
		section,
		file_name,
		file_title,
		file_authors,
		file_subjects,
		file_publisher,
		content='bookmarks_view',
		content_rowid='id',
		prefix='2 3',
		tokenize='porter unicode61 remove_diacritics 2'
	)`
	_, _ = db.conn.Exec(schemaFiles)
	_, _ = db.conn.Exec(schemaBookmarks)
	_, _ = db.conn.Exec(schemaView)
	_, _ = db.conn.Exec(schemaFTS)
	db.createTriggers()
}

// createTriggers: triggers to update FTS index upon insert, delete, and update
func (db *Database) createTriggers() {
	var triggers = `
	CREATE TRIGGER IF NOT EXISTS update_file_name
		INSTEAD OF UPDATE OF file_name ON bookmarks_view
		BEGIN DELETE FROM bookmarksindex where rowid=old.rowid;
		INSERT INTO bookmarksindex(
		rowid, title, section, file_name, file_title, file_authors, file_subjects, file_publisher)
		VALUES (
		new.id, new.title, new.section, new.file_name, new.file_title, new.file_authors, new.file_subjects, new.file_publisher
		); END;
	CREATE TRIGGER IF NOT EXISTS update_file_title
		INSTEAD OF UPDATE OF file_title ON bookmarks_view
		BEGIN DELETE FROM bookmarksindex where rowid=old.rowid;
		INSERT INTO bookmarksindex(
		rowid, title, section, file_name, file_title, file_authors, file_subjects, file_publisher)
		VALUES (
		new.id, new.title, new.section, new.file_name, new.file_title, new.file_authors, new.file_subjects, new.file_publisher
		); END;
	CREATE TRIGGER IF NOT EXISTS update_file_authors
		INSTEAD OF UPDATE OF file_authors ON bookmarks_view
		BEGIN DELETE FROM bookmarksindex where rowid=old.rowid;
		INSERT INTO bookmarksindex(
		rowid, title, section, file_name, file_title, file_authors, file_subjects, file_publisher)
		VALUES (
		new.id, new.title, new.section, new.file_name, new.file_title, new.file_authors, new.file_subjects, new.file_publisher
		); END;
	CREATE TRIGGER IF NOT EXISTS update_file_subjects
		INSTEAD OF UPDATE OF file_subjects ON bookmarks_view
		BEGIN DELETE FROM bookmarksindex where rowid=old.rowid;
		INSERT INTO bookmarksindex(
		rowid, title, section, file_name, file_title, file_authors, file_subjects, file_publisher)
		VALUES (
		new.id, new.title, new.section, new.file_name, new.file_title, new.file_authors, new.file_subjects, new.file_publisher
		); END;
	CREATE TRIGGER IF NOT EXISTS update_file_publisher
		INSTEAD OF UPDATE OF file_publisher ON bookmarks_view
		BEGIN DELETE FROM bookmarksindex where rowid=old.rowid;
		INSERT INTO bookmarksindex(
		rowid, title, section, file_name, file_title, file_authors, file_subjects, file_publisher)
		VALUES (
		new.id, new.title, new.section, new.file_name, new.file_title, new.file_authors, new.file_subjects, new.file_publisher
		); END;
	`
	_, _ = db.conn.Exec(triggers)
}

// BookmarksForFile: retrieve existing bookmarks, add new to database if needed and check if updated
func (db *Database) BookmarksForFile(file string) ([]*Bookmark, int64, error) {
	var bookmarks []*Bookmark
	var err error

	fileRecord, changed, err := db.GetFile(file, true)
	if err != nil {
		return bookmarks, 0, err
	}

	err = db.sess.SelectBySql(`SELECT id, title, section, destination FROM bookmarks
					WHERE file_id = ?`, fileRecord.ID).
		LoadOne(&bookmarks)

	// path created or changed / or no bookmarks found
	if changed == true || len(bookmarks) == 0 {
		var newBookmarks []*FileBookmark

		f := File{}
		if err = f.Init(file); err != nil {
			return bookmarks, 0, err
		}
		newBookmarks, _ = f.Bookmarks()

		// no bookmarks returned from first, get new
		if len(bookmarks) == 0 {
			// insert new
			bookmarks, err = db.NewBookmarks(fileRecord, newBookmarks)
		} else if bookmarksEqual(bookmarks, newBookmarks) == false {
			// path updated, compare bookmarks
			// update database
			bookmarks, err = db.UpdateBookmarks(fileRecord, newBookmarks)
			_ = notification("Bookmarks updated.")
		}
	}

	// run analyze etc upon database close, unsure if faster
	_, _ = db.conn.Exec("PRAGMA optimize;")
	err = db.conn.Close()
	return bookmarks, fileRecord.ID, err
}

// BookmarksForFileFiltered: filtered bookmarks for file, uses fileId from elsewhere so there's only one query
func (db *Database) BookmarksForFileFiltered(file string, query string) ([]*SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []*SearchAllResult

	// only ID needed, no additional fields or checks
	var fileRecord *DatabaseFile
	_ = db.sess.SelectBySql("SELECT id FROM files WHERE path = ?", file).LoadOne(&fileRecord)

	sql := fmt.Sprintf(
		`SELECT bookmarks.id, bookmarks.title, bookmarks.section, bookmarks.destination
			FROM bookmarks
			JOIN bookmarksindex on bookmarks.id = bookmarksindex.rowid
			WHERE bookmarks.file_id = %s
			AND bookmarksindex MATCH '{title section}: %s'
			ORDER BY 'rank(bookmarksindex)'`,
		strconv.FormatInt(fileRecord.ID, 10), *queryString)
	_, err := db.sess.SelectBySql(sql).Load(&results)

	// run analyze etc upon database close, unsure if faster
	_, _ = db.conn.Exec("PRAGMA optimize;")
	err = db.conn.Close()
	return results, err
}

// searchAll: Search all bookmarks from FTS5 table, order by rank title, section, & path name
// Return results as slice of struct SearchAllResult, later prepped for Alfred script filter
func (db *Database) searchAll(query string) ([]*SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []*SearchAllResult

	// NOTE: AND rank MATCH 'bm25(10.0, 5.0)' ORDER BY rank faster than ORDER BY bm25(fts, …)
	_, err := db.sess.SelectBySql(`SELECT
			bookmarks.id, bookmarks.title, bookmarks.section, bookmarks.destination,
			bookmarks.file_id, files.path, files.file_name
			FROM bookmarks
			JOIN files ON bookmarks.file_id = files.id
			JOIN bookmarksindex on bookmarks.id = bookmarksindex.rowid
			WHERE bookmarksindex MATCH ?
			AND rank MATCH 'bm25(10.0, 5.0, 2.0, 1.0, 1.0, 1.0)'
			ORDER BY rank LIMIT 200`,
		queryString).Load(&results)

	err = db.conn.Close()
	return results, err
}

// NewBookmarks: insert new bookmarks into database
func (db *Database) NewBookmarks(file *DatabaseFile, bookmarks []*FileBookmark) ([]*Bookmark, error) {
	tx, err := db.sess.Begin()

	// insert new bookmarks
	for i := range bookmarks {
		_, err = db.sess.InsertBySql(`INSERT INTO
			bookmarks
			(file_id, title, section, destination) VALUES (?, ?, ?, ?)
			`,
			file.ID, NewNullString(bookmarks[i].Title), NewNullString(bookmarks[i].Section), bookmarks[i].Destination).
			Exec()
	}

	err = tx.Commit()

	// get newly inserted bookmarks
	var newBookmarks []*Bookmark
	err = db.sess.SelectBySql(`SELECT id, title, section, destination FROM bookmarks
		WHERE file_id = ?`, file.ID).LoadOne(&newBookmarks)
	return newBookmarks, err
}

// UpdateBookmarks: update bookmarks, delete old first, then call NewBookmarks
func (db *Database) UpdateBookmarks(file *DatabaseFile, bookmarks []*FileBookmark) ([]*Bookmark, error) {
	tx, err := db.sess.Begin()

	_, err = db.sess.DeleteBySql(`DELETE from bookmarks WHERE file_id = ?`, file.ID).Exec()

	err = tx.Commit()
	newBookmarks, err := db.NewBookmarks(file, bookmarks)
	return newBookmarks, err
}

// bookmarksEqual: compare bookmarks from database with path, used for update check
func bookmarksEqual(bookmarks []*Bookmark, newBookmarks []*FileBookmark) bool {
	if len(newBookmarks) != len(bookmarks) {
		return false
	} else {
		for i := range newBookmarks {
			if newBookmarks[i].Title != bookmarks[i].Title {
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

// stringForSQLite: prepare string for SQLite FTS query
// replace '–*' with 'NOT *'
func stringForSQLite(query string) *string {
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
		} else if strings.Contains(term, ".") || strings.Contains(term, "-") {
			// quote terms containing dot
			queryArray = append(queryArray, "\""+term+"*\"")
		} else {
			// make all terms wildcard
			queryArray = append(queryArray, term+"*")
		}
	}
	s := strings.Join(queryArray, " ")
	return &s
}

// notification: macOS notification using github.com/deckarep/gosx-notifier
func notification(message string) error {
	note := gosxnotifier.NewNotification(message)
	note.Title = "Jnana"
	note.Sound = gosxnotifier.Default

	if err := note.Push(); err != nil {
		return err
	}
	return nil
}

// stringInSlice: Test if string is included in slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
