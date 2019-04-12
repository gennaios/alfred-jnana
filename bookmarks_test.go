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

func TestDatabase_Init(t *testing.T) {
	db := initDatabase()
	var row *sql.Row

	//_, err := db.conn.Exec("PRAGMA temp_store=2;")
	//require.NoError(t, err)

	// TODO: running query seems to set to 0
	//row = db.conn.QueryRow("PRAGMA temp_store;")
	//assertPragma(t, row, "2", "temp_store should be 2")

	// TODO: running query seems to set to 0
	//row = db.conn.QueryRow("PRAGMA auto_vacuum;")
	//assertPragma(t, row, "2", "auto_vacuum should be 2")

	row = db.conn.QueryRow("PRAGMA foreign_keys;")
	assertPragma(t, row, "1", "foreign_keys should be 1")

	row = db.conn.QueryRow("PRAGMA journal_mode;")
	assertPragma(t, row, "wal", "journal_mode should be WAL")

	row = db.conn.QueryRow("PRAGMA page_size;")
	assertPragma(t, row, "4096", "page_size should be 4096")

	row = db.conn.QueryRow("PRAGMA synchronous;")
	assertPragma(t, row, "0", "synchronous should be 0")

	_ = db.conn.Close()
}

//func TestInitMemory(t *testing.T) {
//	db := Database{}
//	db.Init("file::memory:?mode=memory&cache=shared")
//
//}

func TestDatabase_BookmarksForFile(t *testing.T) {
	file, _ := filepath.Abs("./tests/pdf.pdf")
	if _, err := os.Stat(file); err != nil {
		log.Fatal(err)
	}

	db := initDatabase()
	bookmarks, _ := db.BookmarksForFile(file)
	assert.Equal(t, len(bookmarks), 4, "Bookmarks count should be 4")
}

func TestDatabase_BookmarksForFileFiltered(t *testing.T) {
	file, _ := filepath.Abs("./tests/pdf.pdf")
	if _, err := os.Stat(file); err != nil {
		log.Fatal(err)
	}

	db := initDatabase()
	bookmarks, _ := db.BookmarksForFileFiltered(file, "links")
	assert.Equal(t, len(bookmarks), 2, "Bookmarks count should be 2")
}
