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
		db.createTriggers()
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
	_, _ = db.Db.Exec(triggers)
}
