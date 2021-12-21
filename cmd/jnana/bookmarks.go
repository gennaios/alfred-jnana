package main

import (
	"jnana/models"

	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"

	"database/sql"
	"fmt"
	"github.com/deckarep/gosx-notifier"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/null/v8"
	"strconv"
	"strings"
)

type Database struct {
	db  *sql.DB
	ctx context.Context
}

type SearchAllResult struct {
	ID          int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Section     null.String `boil:"section" json:"section,omitempty" toml:"section" yaml:"section,omitempty"`
	Destination string      `boil:"destination" json:"destination" toml:"destination" yaml:"hash"`
	FileID      string      `boil:"file_id" json:"file_id" toml:"file_id" yaml:"file_id"`
	Path        string      `boil:"path" json:"path" toml:"path" yaml:"path"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
}

// Init open SQLite database connection, create new session
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
	db.db, err = sql.Open("sqlite3", file)
	boil.SetDB(db.db)
	db.ctx = context.Background()

	// TODO: return error
	if err != nil {
		panic(err)
	}

	_, _ = queries.Raw("PRAGMA auto_vacuum=2").Exec(db.db) // unsure if set
	//_, _ = queries.Raw(db.db, "PRAGMA temp_store = 2")  // MEMORY
	//_, _ = queries.Raw(db.db, "PRAGMA cache_size = -31250")

	type TableName struct {
		Name string `boil:"name"`
	}
	var result TableName

	// create tables and triggers if 'file' does not exist
	_ = queries.Raw("SELECT name FROM sqlite_master WHERE type='table' AND name='file'").Bind(db.ctx, db.db, &result)
	if result.Name != "file" {
		db.createTables()
		db.createTriggers()
	}

	if err != nil {
		panic(err)
	}
}

// InitForReading open SQLite database connection, for reading, doesn't create tables etc
func (db *Database) InitForReading(dbFilePath string) {
	var err error
	db.db, err = sql.Open("sqlite3", dbFilePath+"?&mode=ro&_journal_mode=WAL&cache=shared")

	if err != nil {
		panic(err)
	}

	boil.SetDB(db.db)
	db.ctx = context.Background()
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
	_, _ = db.db.Exec(schemaFiles)
	_, _ = db.db.Exec(schemaBookmarks)
	_, _ = db.db.Exec(schemaView)
	_, _ = db.db.Exec(fileFTS)
	_, _ = db.db.Exec(bookmarkFTS)
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
	_, _ = db.db.Exec(triggers)
}

// BookmarksForFile retrieve existing bookmarks, add new to database if needed and check if updated
func (db *Database) BookmarksForFile(file string, coversCacheDir string) ([]*models.Bookmark, error) {
	var bookmarks []*models.Bookmark
	var err error

	fileRecord, changed, err := db.GetFile(file, true)
	if err != nil {
		return bookmarks, err
	}

	err = queries.Raw(`SELECT id, title, section, destination FROM bookmark
					WHERE file_id = $1`, fileRecord.ID).Bind(db.ctx, db.db, &bookmarks)

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
	err = db.db.Close()
	return bookmarks, err
}

// BookmarksForFileFiltered filtered bookmarks for file, uses fileId from elsewhere so there's only one query
func (db *Database) BookmarksForFileFiltered(file string, query string) ([]*SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []*SearchAllResult

	// only ID needed, no additional fields or checks
	var fileRecord *models.File

	fileRecord, _ = models.Files(models.FileWhere.Path.EQ(file)).One(db.ctx, db.db)
	fmt.Println("record: " + fileRecord.Path)

	err := queries.Raw(`SELECT
			bookmark.id, bookmark.title, bookmark.section, bookmark.destination
			FROM bookmark
			JOIN bookmark_search on bookmark.id = bookmark_search.rowid `+
		`WHERE bookmark.file_id = `+strconv.FormatInt(fileRecord.ID, 10)+
		` AND bookmark_search MATCH '{title section}: `+*queryString+"' "+
		`ORDER BY 'rank(bookmark_search)'`).Bind(db.ctx, db.db, &results)

	// run analyze etc upon database close, unsure if faster
	//_, _ = db.conn.Exec("PRAGMA optimize;")
	_ = db.db.Close()
	return results, err
}

// searchAll: Search all bookmarks from FTS5 table, order by rank title, section, & file name
// Return results as slice of struct SearchAllResult, later prepped for Alfred script filter
func (db *Database) searchAll(query string) ([]*SearchAllResult, error) {
	queryString := stringForSQLite(query)
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
		queryString).Bind(db.ctx, db.db, &results)

	_ = db.db.Close()
	return results, err
}

// NewBookmarks insert new bookmarks into database
func (db *Database) NewBookmarks(file *models.File, bookmarks []*FileBookmark) ([]*models.Bookmark, error) {
	tx, err := db.db.BeginTx(db.ctx, nil)

	// insert new bookmarks
	for i := range bookmarks {
		_, err = queries.Raw(`INSERT INTO
			bookmark
			(file_id, title, section, destination) VALUES ($1, $2, $3, $4)
			`,
			file.ID, bookmarks[i].Title, bookmarks[i].Section, bookmarks[i].Destination).
			Exec(db.db)
	}

	err = tx.Commit()

	// get newly inserted bookmarks
	var newBookmarks []*models.Bookmark
	err = queries.Raw(`SELECT id, title, section, destination FROM bookmark
		WHERE file_id = $1`, file.ID).BindG(db.ctx, &newBookmarks)
	return newBookmarks, err
}

// UpdateBookmarks update bookmarks, delete old first, then call NewBookmarks
func (db *Database) UpdateBookmarks(file *models.File, oldBookmarks []*models.Bookmark, newBookmarks []*FileBookmark) ([]*models.Bookmark, error) {
	var err error
	var results []*models.Bookmark

	tx, _ := db.db.BeginTx(db.ctx, nil)
	if len(oldBookmarks) == len(newBookmarks) {
		// count same, update records
		for i := range oldBookmarks {
			_, err = queries.Raw(`UPDATE bookmark SET
				title = $1, section = $2, destination = $3
				WHERE id = $4`,
				newBookmarks[i].Title,
				newBookmarks[i].Section,
				newBookmarks[i].Destination, oldBookmarks[i].ID).
				Exec(db.db)
		}
		err = queries.Raw(`SELECT id, title, section, destination FROM bookmark
			WHERE file_id = $1`, file.ID).BindG(db.ctx, &results)
	} else {
		// count different, delete and insert new
		_, err = queries.Raw(`DELETE from bookmark WHERE file_id = $1`, file.ID).Exec(db.db)
		results, err = db.NewBookmarks(file, newBookmarks)
	}
	err = tx.Commit()

	return results, err
}

// bookmarksEqual compare bookmarks from database with path, used for update check
func bookmarksEqual(bookmarks []*models.Bookmark, newBookmarks []*FileBookmark) bool {
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
