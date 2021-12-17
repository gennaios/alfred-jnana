package main

import (
	"strconv"
	"strings"

	"github.com/deckarep/gosx-notifier"
	"github.com/gocraft/dbr/v2"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	conn *dbr.Connection
	sess *dbr.Session
}

type Bookmark struct {
	ID          int64
	FileId      int64          `db:"file_id"`
	Title       dbr.NullString `db:"title"`
	Section     dbr.NullString `db:"section"`
	Destination string         `db:"destination"`
}

type SearchAllResult struct {
	ID          int64
	Title       dbr.NullString `db:"title"`
	Section     dbr.NullString `db:"section"`
	Destination string         `db:"destination"`
	FileID      string         `db:"file_id"`
	Path        string         `db:"path"`
	Name        string         `db:"name"`
}

// Init: open SQLite database connection using dbr, create new session
func (db *Database) Init(dbFilePath string) {
	var file string
	// open with PRAGMAs:
	// journal_mode=WAL, synchronous=0
	if dbFilePath != "memory" {
		file = dbFilePath + "?&_journal_mode=WAL&_synchronous=0&_foreign_keys=1"
	} else {
		file = "file::memory:?mode=memory&cache=shared&_foreign_keys=1&_synchronous=0"
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

	// create tables and triggers if 'file' does not exist
	tables := db.sess.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='file'")
	var result string
	_ = tables.Scan(&result)
	if result != "file" {
		db.createTables()
		db.createTriggers()
	}

	if err != nil {
		panic(err)
	}
}

// Init: open SQLite database connection using dbr, for reading, doesn't create tables etc
func (db *Database) InitForReading(dbFilePath string) {
	db.conn, _ = dbr.Open("sqlite3", dbFilePath+"?&mode=ro&_journal_mode=WAL&cache=shared", nil)
	db.sess = db.conn.NewSession(nil)
	_, _ = db.sess.Begin()
}

// createTables: create tables file, bookmark, view bookmark_view for FTS updates, and FTS5 bookmark_search
func (db *Database) createTables() {
	var schemaFiles = `
	CREATE TABLE IF NOT EXISTS file (
		id INTEGER NOT NULL PRIMARY KEY,
		path TEXT NOT NULL,
	    	name TEXT,
	    	extension VARCHAR(255) NOT NULL,
	    	size INTEGER NOT NULL,
	    	title TEXT,
	    	creator TEXT,
	    	subject TEXT,
	    	publisher TEXT,
	    	language TEXT,
	    	description TEXT,
	    	date_created DATETIME NOT NULL,
	    	date_modified DATETIME,
	    	date_accessed DATETIME,
	    	rating INTEGER,
	    	hash VARCHAR(64) NOT NULL
	)`
	var schemaBookmarks = `
	CREATE TABLE IF NOT EXISTS bookmark (
		id INTEGER NOT NULL PRIMARY KEY,
		file_id INTEGER NOT NULL,
		title TEXT,
		section TEXT,
		destination TEXT NOT NULL,
		FOREIGN KEY (file_id) REFERENCES file (id) ON DELETE CASCADE ON UPDATE CASCADE
	)`
	var schemaView = `
	CREATE VIEW IF NOT EXISTS bookmark_view AS SELECT
		bookmark.id,
		bookmark.title,
		bookmark.section,
		file.name,
		file.title as file_title,
		file.creator,
		file.subject,
		file.publisher
		FROM bookmark
		INNER JOIN file ON file.id = bookmark.file_id
	`
	// prefix: tokenize by length
	var fileFTS = `
	CREATE VIRTUAL TABLE file_search USING fts5(
		name,
		title,
		creator,
		subject,
		publisher,
		description,
		content='file',
		content_rowid='id',
		prefix='3',
		tokenize='porter unicode61 remove_diacritics 2'
	)`
	var bookmarkFTS = `
	CREATE VIRTUAL TABLE bookmark_search USING fts5(
		title,
		section,
		name,
		file_title,
		creator,
		subject,
		publisher,
		content='bookmark_view',
		content_rowid='id',
		prefix='3',
		tokenize='porter unicode61 remove_diacritics 2'
	)`
	_, _ = db.conn.Exec(schemaFiles)
	_, _ = db.conn.Exec(schemaBookmarks)
	_, _ = db.conn.Exec(schemaView)
	_, _ = db.conn.Exec(fileFTS)
	_, _ = db.conn.Exec(bookmarkFTS)
	db.createTriggers()
}

// createTriggers: triggers to update FTS index upon insert, delete, and update
func (db *Database) createTriggers() {
	var triggers = `
	CREATE TRIGGER IF NOT EXISTS file_delete
		AFTER DELETE ON file
		BEGIN DELETE FROM file_search where rowid=old.id;
		END;
	CREATE TRIGGER IF NOT EXISTS file_insert
		AFTER INSERT ON file
		BEGIN INSERT INTO file_search(rowid, name, title, creator, subject, publisher, description)
		VALUES (new.id, new.name, new.title, new.creator, new.subject, new.publisher, new.description);
		END;
	CREATE TRIGGER IF NOT EXISTS bookmark_delete
		AFTER DELETE ON bookmark
		BEGIN DELETE FROM bookmark_search where rowid=old.id;
		END;
	CREATE TRIGGER IF NOT EXISTS bookmark_insert
		AFTER INSERT ON bookmark
		BEGIN INSERT INTO bookmark_search(rowid, title, section, name, file_title, creator, subject, publisher)
		VALUES (new.id, new.title, new.section, (SELECT name FROM file WHERE id = new.file_id), (SELECT title FROM file WHERE id = new.file_id), (SELECT creator FROM file WHERE id = new.file_id), (SELECT subject FROM file WHERE id = new.file_id), (SELECT publisher FROM file WHERE id = new.file_id));
		END;
	CREATE TRIGGER IF NOT EXISTS bookmark_update
		AFTER UPDATE ON bookmark
		BEGIN INSERT INTO bookmark_search(rowid, title, section, name, file_title, creator, subject, publisher)
		VALUES (new.id, new.title, new.section, (SELECT name FROM file WHERE id = new.file_id), (SELECT title FROM file WHERE id = new.file_id), (SELECT creator FROM file WHERE id = new.file_id), (SELECT subject FROM file WHERE id = new.file_id), (SELECT publisher FROM file WHERE id = new.file_id));
		END;
	CREATE TRIGGER IF NOT EXISTS update_file_name
		INSTEAD OF UPDATE OF name ON bookmark_view
		BEGIN DELETE FROM bookmark_search where rowid=old.rowid;
		INSERT INTO bookmark_search(
		rowid, title, section, name, file_title, creator, subject, publisher)
		VALUES (
		new.id, new.title, new.section, new.name, new.file_title, new.creator, new.subject, new.publisher
		); END;
	CREATE TRIGGER IF NOT EXISTS update_title
		INSTEAD OF UPDATE OF title ON bookmark_view
		BEGIN DELETE FROM bookmark_search where rowid=old.rowid;
		INSERT INTO bookmark_search(
		rowid, title, section, name, file_title, creator, subject, publisher)
		VALUES (
		new.id, new.title, new.section, new.name, new.file_title, new.creator, new.subject, new.publisher
		); END;
	CREATE TRIGGER IF NOT EXISTS update_creator
		INSTEAD OF UPDATE OF creator ON bookmark_view
		BEGIN DELETE FROM bookmark_search where rowid=old.rowid;
		INSERT INTO bookmark_search(
		rowid, title, section, name, file_title, creator, subject, publisher)
		VALUES (
		new.id, new.title, new.section, new.name, new.file_title, new.creator, new.subject, new.publisher
		); END;
	CREATE TRIGGER IF NOT EXISTS update_subject
		INSTEAD OF UPDATE OF subject ON bookmark_view
		BEGIN DELETE FROM bookmark_search where rowid=old.rowid;
		INSERT INTO bookmark_search(
		rowid, title, section, name, file_title, creator, subject, publisher)
		VALUES (
		new.id, new.title, new.section, new.name, new.file_title, new.creator, new.subject, new.publisher
		); END;
	CREATE TRIGGER IF NOT EXISTS update_publisher
		INSTEAD OF UPDATE OF publisher ON bookmark_view
		BEGIN DELETE FROM bookmark_search where rowid=old.rowid;
		INSERT INTO bookmark_search(
		rowid, title, section, name, file_title, creator, subject, publisher)
		VALUES (
		new.id, new.title, new.section, new.name, new.file_title, new.creator, new.subject, new.publisher
		); END;
	CREATE TRIGGER IF NOT EXISTS update_file_name
		AFTER UPDATE OF name ON file
		BEGIN DELETE FROM file_search where rowid=old.id;
		INSERT INTO file_search(
		rowid, name, title, creator, subject, publisher, description)
		VALUES (
		new.id, new.name, new.title, new.creator, new.subject, new.publisher, new.description
		); END;
	CREATE TRIGGER IF NOT EXISTS update_file_title
		AFTER UPDATE OF title ON file
		BEGIN DELETE FROM file_search where rowid=old.id;
		INSERT INTO file_search(
		rowid, name, title, creator, subject, publisher, description)
		VALUES (
		new.id, new.name, new.title, new.creator, new.subject, new.publisher, new.description
		); END;
	CREATE TRIGGER IF NOT EXISTS update_file_creator
		AFTER UPDATE OF creator ON file
		BEGIN DELETE FROM file_search where rowid=old.id;
		INSERT INTO file_search(
		rowid, name, title, creator, subject, publisher, description)
		VALUES (
		new.id, new.name, new.title, new.creator, new.subject, new.publisher, new.description
		); END;
	CREATE TRIGGER IF NOT EXISTS update_file_subject
		AFTER UPDATE OF subject ON file
		BEGIN DELETE FROM file_search where rowid=old.id;
		INSERT INTO file_search(
		rowid, name, title, creator, subject, publisher, description)
		VALUES (
		new.id, new.name, new.title, new.creator, new.subject, new.publisher, new.description
		); END;
	CREATE TRIGGER IF NOT EXISTS update_file_publisher
		AFTER UPDATE OF publisher ON file
		BEGIN DELETE FROM file_search where rowid=old.id;
		INSERT INTO file_search(
		rowid, name, title, creator, subject, publisher, description)
		VALUES (
		new.id, new.name, new.title, new.creator, new.subject, new.publisher, new.description
		); END;
	CREATE TRIGGER IF NOT EXISTS update_file_description
		AFTER UPDATE OF description ON file
		BEGIN DELETE FROM file_search where rowid=old.id;
		INSERT INTO file_search(
		rowid, name, title, creator, subject, publisher, description)
		VALUES (
		new.id, new.name, new.title, new.creator, new.subject, new.publisher, new.description
		); END;
	`
	_, _ = db.conn.Exec(triggers)
}

// BookmarksForFile: retrieve existing bookmarks, add new to database if needed and check if updated
func (db *Database) BookmarksForFile(file string, coversCacheDir string) ([]*Bookmark, error) {
	var bookmarks []*Bookmark
	var err error

	fileRecord, changed, err := db.GetFile(file, true)
	if err != nil {
		return bookmarks, err
	}

	err = db.sess.SelectBySql(`SELECT id, title, section, destination FROM bookmark
					WHERE file_id = ?`, fileRecord.ID).
		LoadOne(&bookmarks)

	// check cover / TODO: move somewhere else?
	//_ = db.CoverForFile(fileRecord, coversCacheDir)

	// file created or changed / or no bookmarks found
	if changed == true || len(bookmarks) == 0 {
		var newBookmarks []*FileBookmark

		f := File{}
		if err = f.Init(file); err != nil {
			return bookmarks, err
		}
		newBookmarks, _ = f.Bookmarks()

		// no bookmarks returned from first, get new
		if len(bookmarks) == 0 {
			// insert new
			bookmarks, err = db.NewBookmarks(fileRecord, newBookmarks)
		} else if bookmarksEqual(bookmarks, newBookmarks) == false {
			// file updated, compare bookmarks
			// update database
			bookmarks, err = db.UpdateBookmarks(fileRecord, bookmarks, newBookmarks)
			_ = notification("Bookmarks updated.")
		}
	}

	// run analyze etc upon database close, unsure if faster
	//_, _ = db.conn.Exec("PRAGMA optimize;")
	err = db.conn.Close()
	return bookmarks, err
}

// BookmarksForFileFiltered: filtered bookmarks for file, uses fileId from elsewhere so there's only one query
func (db *Database) BookmarksForFileFiltered(file string, query string) ([]*SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []*SearchAllResult

	// only ID needed, no additional fields or checks
	var fileRecord *DatabaseFile
	_ = db.sess.SelectBySql("SELECT id FROM file WHERE path = ?", file).LoadOne(&fileRecord)

	_, err := db.sess.SelectBySql(`SELECT
			bookmark.id, bookmark.title, bookmark.section, bookmark.destination
			FROM bookmark
			JOIN bookmark_search on bookmark.id = bookmark_search.rowid
			WHERE bookmark.file_id = ` + strconv.FormatInt(fileRecord.ID, 10) +
		` AND bookmark_search MATCH '{title section}: ` + *queryString +
		`' ORDER BY 'rank(bookmark_search)'`).Load(&results)

	// run analyze etc upon database close, unsure if faster
	//_, _ = db.conn.Exec("PRAGMA optimize;")
	err = db.conn.Close()
	return results, err
}

// searchAll: Search all bookmarks from FTS5 table, order by rank title, section, & file name
// Return results as slice of struct SearchAllResult, later prepped for Alfred script filter
func (db *Database) searchAll(query string) ([]*SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []*SearchAllResult

	// NOTE: AND rank MATCH 'bm25(10.0, 5.0)' ORDER BY rank faster than ORDER BY bm25(fts, …)
	_, err := db.sess.SelectBySql(`SELECT
			bookmark.id, bookmark.title, bookmark.section, bookmark.destination,
			bookmark.file_id, file.path, file.name
			FROM bookmark
			JOIN file ON bookmark.file_id = file.id
			JOIN bookmark_search on bookmark.id = bookmark_search.rowid
			WHERE bookmark_search MATCH ?
			AND rank MATCH 'bm25(10.0, 5.0, 2.0, 1.0, 1.0, 1.0, 1.0)'
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
			bookmark
			(file_id, title, section, destination) VALUES (?, ?, ?, ?)
			`,
			file.ID, NewNullString(bookmarks[i].Title), NewNullString(bookmarks[i].Section), bookmarks[i].Destination).
			Exec()
	}

	err = tx.Commit()

	// get newly inserted bookmarks
	var newBookmarks []*Bookmark
	err = db.sess.SelectBySql(`SELECT id, title, section, destination FROM bookmark
		WHERE file_id = ?`, file.ID).LoadOne(&newBookmarks)
	return newBookmarks, err
}

// UpdateBookmarks: update bookmarks, delete old first, then call NewBookmarks
func (db *Database) UpdateBookmarks(file *DatabaseFile, oldBookmarks []*Bookmark, newBookmarks []*FileBookmark) ([]*Bookmark, error) {
	var err error
	var results []*Bookmark

	tx, _ := db.sess.Begin()
	if len(oldBookmarks) == len(newBookmarks) {
		// count same, update records
		for i := range oldBookmarks {
			_, err = db.sess.UpdateBySql(`UPDATE bookmark SET
				title = ?, section = ?, destination = ?
				WHERE id = ?`,
				NewNullString(newBookmarks[i].Title),
				NewNullString(newBookmarks[i].Section),
				newBookmarks[i].Destination, oldBookmarks[i].ID).
				Exec()
		}
		err = db.sess.SelectBySql(`SELECT id, title, section, destination FROM bookmark
			WHERE file_id = ?`, file.ID).LoadOne(&results)
	} else {
		// count different, delete and insert new
		_, err = db.sess.DeleteBySql(`DELETE from bookmark WHERE file_id = ?`, file.ID).Exec()
		results, err = db.NewBookmarks(file, newBookmarks)
	}
	err = tx.Commit()

	return results, err
}

// bookmarksEqual: compare bookmarks from database with path, used for update check
func bookmarksEqual(bookmarks []*Bookmark, newBookmarks []*FileBookmark) bool {
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

// stringForSQLite: prepare string for SQLite FTS query
// make all terms wildcard
func stringForSQLite(query string) *string {
	var querySlice []string
	queryOperators := []string{"and", "or", "not", "AND", "OR", "NOT"}

	slc := strings.Split(query, " ")
	for i := range slc {
		term := slc[i]
		if strings.HasPrefix(term, "-") {
			// exclude terms beginning with '-', change to 'NOT [term]'
			querySlice = append(querySlice, "NOT "+term[1:]+"*")
		} else if stringInSlice(term, queryOperators) {
			// auto capitalize operators 'and', 'or', 'not'
			querySlice = append(querySlice, strings.ToUpper(term))
		} else if strings.Contains(term, ".") || strings.Contains(term, "-") {
			// quote terms containing dot
			querySlice = append(querySlice, "\""+term+"*\"")
		} else {
			querySlice = append(querySlice, term+"*")
		}
	}
	s := strings.Join(querySlice, " ")
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
	for i := range list {
		if list[i] == a {
			return true
		}
	}
	return false
}
