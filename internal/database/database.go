package database

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

type Database struct {
	Db  *sql.DB
	Ctx context.Context
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
	db.Db, err = sql.Open("sqlite3", file)
	boil.SetDB(db.Db)
	db.Ctx = context.Background()

	// TODO: return error
	if err != nil {
		panic(err)
	}

	_, _ = queries.Raw("PRAGMA auto_vacuum=2").Exec(db.Db) // unsure if set
	//_, _ = queries.Raw(db.Db, "PRAGMA temp_store = 2")  // MEMORY
	//_, _ = queries.Raw(db.Db, "PRAGMA cache_size = -31250")

	type TableName struct {
		Name string `boil:"name"`
	}
	var result TableName

	// create tables and triggers if 'file' does not exist
	_ = queries.Raw("SELECT name FROM sqlite_master WHERE type='table' AND name='file'").Bind(db.Ctx, db.Db, &result)
	if result.Name != "file" {
		db.createTables()
	}

	if err != nil {
		panic(err)
	}
}

// InitForReading open SQLite database connection, for reading, doesn't create tables etc
func (db *Database) InitForReading(dbFilePath string) {
	var err error
	db.Db, err = sql.Open("sqlite3", dbFilePath+"?&mode=ro&_journal_mode=WAL&cache=shared")

	if err != nil {
		panic(err)
	}

	boil.SetDB(db.Db)
	db.Ctx = context.Background()
}

// createTables: create tables file, bookmark, view bookmark_view for FTS updates, and FTS5 bookmark_search
func (db *Database) createTables() {
	var schemaFiles = `
	CREATE TABLE IF NOT EXISTS file (
		id INTEGER NOT NULL PRIMARY KEY,
		path TEXT NOT NULL,
	    	name TEXT,
	    	format INTEGER NOT NULL,
	    	size INTEGER NOT NULL,
	    	title TEXT,
	    	creator TEXT,
	    	subject TEXT,
	    	publisher TEXT,
	    	publisher_id INTEGER,
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
		file.id as file_id,
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
	_, _ = db.Db.Exec(schemaFiles)
	_, _ = db.Db.Exec(schemaBookmarks)
	_, _ = db.Db.Exec(schemaView)
	_, _ = db.Db.Exec(fileFTS)
	_, _ = db.Db.Exec(bookmarkFTS)
}
