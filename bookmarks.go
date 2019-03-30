package main

import (
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
	"strings"
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

	sess := conn.NewSession(nil)
	_, err = sess.Begin()
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	return sess
}

func searchAll(sess *dbr.Session, query string) ([]SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []SearchAllResult

	//err := db.Select(&results, `SELECT
	//	bookmarks.id, bookmarks.title, bookmarks.section, bookmarks.destination,
	//	files.file_name, files.path
	//	FROM bookmarks
	//	JOIN files ON files.id = bookmarks.file_id
	//	JOIN bookmarkindex on bookmarks.id = bookmarkindex.rowid
	//	WHERE bookmarkindex MATCH '?' ORDER BY rank LIMIT 100`, queryString)

	_, err := sess.Select("bookmarks.id", "bookmarks.title", "bookmarks.section",
		"bookmarks.destination", "bookmarks.file_id", "files.path", "files.file_name").
		From("bookmarks").
		Join("files", "bookmarks.file_id = files.id").
		Join("bookmarkindex", "bookmarks.id = bookmarkindex.rowid").
		Where("bookmarkindex MATCH ?", queryString).
		OrderBy("rank").Limit(100).Load(&results)

	return results, err
}

// Prepare string for SQLite FTS query
// replace '–*' with 'NOT *'
func stringForSQLite(query string) string {
	var queryArray []string
	query = strings.TrimSpace(query)

	slc := strings.Split(query, " ")
	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])

		if strings.HasPrefix(slc[i], "-") {
			queryArray = append(queryArray, "NOT "+slc[i][1:]+"*")
		} else {
			queryArray = append(queryArray, slc[i]+"*")
		}
	}

	return strings.Join(queryArray[:], " ")
}
