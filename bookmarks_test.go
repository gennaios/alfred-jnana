package main

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
)

// 	"github.com/stretchr/testify/require"

func assertPragma(t *testing.T, row *sql.Row, expected string, pragma string) {
	var result string
	if err := row.Scan(&result); err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, result, expected, pragma)
}

func TestDatabase_Init(t *testing.T) {
	db := Database{}
	db = initDatabase("memory")
	var row *sql.Row

	//_, err := db.conn.Exec("PRAGMA temp_store=2;")
	//require.NoError(t, err)

	// TODO: running query seems to set to 0
	//row = db.conn.QueryRow("PRAGMA temp_store;")
	//assertPragma(t, row, "2", "temp_store should be 2")

	// TODO: running query seems to set to 0
	//row = db.conn.QueryRow("PRAGMA auto_vacuum;")
	//assertPragma(t, row, "2", "auto_vacuum should be 2")

	// test SQLite PRAGMAs
	row = db.conn.QueryRow("PRAGMA foreign_keys;")
	assertPragma(t, row, "1", "foreign_keys should be 1")

	row = db.conn.QueryRow("PRAGMA journal_mode;")
	assertPragma(t, row, "memory", "journal_mode should be WAL")

	row = db.conn.QueryRow("PRAGMA page_size;")
	assertPragma(t, row, "4096", "page_size should be 4096")

	row = db.conn.QueryRow("PRAGMA synchronous;")
	assertPragma(t, row, "0", "synchronous should be 0")

	// test table creation
	var result string
	filesTable := db.sess.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='files'")
	_ = filesTable.Scan(&result)
	assert.Equal(t, "files", result, "`files` table not created")
	bookmarksTable := db.sess.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='bookmarks'")
	_ = bookmarksTable.Scan(&result)
	assert.Equal(t, "bookmarks", result, "`bookmarks` table not created")
	bookmarksindexTable := db.sess.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='bookmarksindex'")
	_ = bookmarksindexTable.Scan(&result)
	assert.Equal(t, "bookmarksindex", result, "`bookmarksindex` table not created")

	_ = db.conn.Close()
}

func TestDatabase_BookmarksForFile(t *testing.T) {
	file, _ := filepath.Abs("./tests/pdf.pdf")
	if _, err := os.Stat(file); err != nil {
		log.Fatal(err)
	}

	db := Database{}
	db = initDatabase("memory")

	bookmarks, _ := db.BookmarksForFile(file)
	assert.Equal(t, 4, len(bookmarks), "Bookmarks count should be 4")
}

func TestDatabase_BookmarksForFileFiltered(t *testing.T) {
	file, _ := filepath.Abs("./tests/pdf.pdf")
	if _, err := os.Stat(file); err != nil {
		log.Fatal(err)
	}
	db := Database{}
	db = initDatabase("memory")

	bookmarks, _ := db.BookmarksForFileFiltered(file, "links")
	assert.Equal(t, 2, len(bookmarks), "Bookmarks count should be 2")
}
