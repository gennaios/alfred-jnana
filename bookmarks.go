package main

import (
	"strings"

	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

type SearchAllResult struct {
	ID          int64
	Title       string         `db:"title"`
	Section     dbr.NullString `db:"section"`
	Destination string         `db:"destination"`
	FileID      string         `db:"file_id"`
	Path        string         `db:"path"`
	FileName    string         `db:"file_name"`
}

func initDatabase(filepath string) *dbr.Session {
	conn, err := dbr.Open("sqlite3", filepath, nil)
	if err != nil {
		panic(err)
	}
	if conn == nil {
		panic("db nil")
	}

	_, _ = conn.Exec("PRAGMA auto_vacuum = 1")
	_, _ = conn.Exec("PRAGMA foreign_keys = 1")
	_, _ = conn.Exec("PRAGMA ignore_check_constraints = 0")
	_, _ = conn.Exec("PRAGMA journal_mode = WAL")
	_, _ = conn.Exec("PRAGMA synchronous = 0")
	_, _ = conn.Exec("PRAGMA temp_store = 2") // MEMORY

	_, _ = conn.Exec("PRAGMA cache_size = -31250")
	_, _ = conn.Exec("PRAGMA page_size = 8192") // default 4096, match APFS block size?

	sess := conn.NewSession(nil)
	_, err = sess.Begin()
	if err != nil {
		panic(err)
	}

	return sess
}

// Search all bookmarks from FTS5 table, order by rank title, section, & file name
// Return results as slice of struct SearchAllResult, later preped for Alfred script filter
func searchAll(sess *dbr.Session, query string) ([]SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []SearchAllResult

	//	SELECT
	//	bookmarks.id, bookmarks.title, bookmarks.section, bookmarks.destination,
	//	files.file_name, files.path
	//	FROM bookmarks
	//	JOIN files ON files.id = bookmarks.file_id
	//	JOIN bookmarksindex on bookmarks.id = bookmarkindex.rowid
	//	WHERE bookmarksindex MATCH '?' AND rank MATCH 'bm25(5.0, 2.0, 1.0)'

	// NOTE: AND rank MATCH 'bm25(10.0, 5.0)' ORDER BY rank faster than ORDER BY bm25(fts, …)
	_, err := sess.Select("bookmarks.id", "bookmarks.title", "bookmarks.section",
		"bookmarks.destination", "bookmarks.file_id", "files.path", "files.file_name").
		From("bookmarks").
		Join("files", "bookmarks.file_id = files.id").
		Join("bookmarksindex", "bookmarks.id = bookmarksindex.rowid").
		Where("bookmarksindex MATCH ? AND rank MATCH 'bm25(5.0, 2.0, 1.0)'", queryString).
		OrderBy("rank").Limit(100).Load(&results)

	return results, err
}

// Prepare string for SQLite FTS query
// replace '–*' with 'NOT *'
func stringForSQLite(query string) string {
	var queryArray []string

	slc := strings.Split(query, " ")
	for i := range slc {
		term := strings.TrimSpace(slc[i])

		if strings.HasPrefix(term, "-") {
			queryArray = append(queryArray, "NOT "+term[1:]+"*")
		} else {
			queryArray = append(queryArray, term+"*")
		}
	}

	return strings.Join(queryArray[:], " ")
}
