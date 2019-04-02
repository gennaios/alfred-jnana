package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

type FileRecord struct {
	ID            int64
	Path          string `db:"path"`
	FileName      string `db:"file_name"`
	FileExtension string `db:"file_extension"`
	Title         string `db:"title"`
	Authors       string `db:"authors"`
	Subjects      string `db:"subjects"`
	DateCreated   string `db:"date_created"`
	DateModified  string `db:"file_modified_date"` // TODO: rename -> date_modified
	FileHash      string `db:"hash"`
}

type BookmarkRecord struct {
	ID          int64
	FileId      int64
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

func databaseRecordForFile(dbFile, file string) FileRecord {
	sess := initDatabase(dbFile)
	var fileRecord FileRecord

	// look for existing
	err := sess.Select("id", "file_name", "file_modified_date", "hash").
		From("files").Where("path == ?", file).
		LoadOne(&fileRecord)
	if err != nil {
		// TODO: error to Alfred
		panic(err)
	}

	// TODO: if none, check by hash
	if fileRecord.ID == 0 {
		fmt.Println("Doesn't exist: ", file)
	}

	// TODO: if none by hash, create

	return fileRecord
}

func forFile(dbFile string, file string) ([]BookmarkRecord, error) {
	var bookmarks []BookmarkRecord
	var err error

	fileRecord := databaseRecordForFile(dbFile, file)

	// TODO, record created if none, always not 0
	if fileRecord.ID != 0 {
		sess := initDatabase(dbFile)

		// look for existing
		err = sess.Select("id", "title", "section", "destination").From("bookmarks").
			Where("file_id == ?", fileRecord.ID).
			LoadOne(&bookmarks)
	}

	if fileRecord.ID == 0 {
		// get from file and import
		if strings.HasSuffix(file, ".pdf") {
			// TODO: call pdf module
			bookmarks, _ := bookmarksForPDF(file)
			print("Bookmarks: ", len(bookmarks), bookmarks[0].Title, bookmarks[0].Section, bookmarks[0].Destination)
		}
	}

	return bookmarks, err
}

// Filtered bookmarks for file
func forFileFiltered(dbFile string, file string, query string) ([]SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []SearchAllResult

	sess := initDatabase(dbFile)
	fileRecord := databaseRecordForFile(dbFile, file)

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
	_, err := sess.SelectBySql("select bookmarks.id, bookmarks.title, bookmarks.section, bookmarks.destination from bookmarks JOIN bookmarksindex on bookmarks.id = bookmarksindex.rowid where bookmarks.file_id = " + strconv.FormatInt(fileRecord.ID, 10) + " and bookmarksindex.rowid = bookmarks.id and bookmarksindex match '{title section} : " + queryString + "' ORDER BY 'bm25(bookmarksindex, 5.0, 2.0)';").Load(&results)
	return results, err
}

// Search all bookmarks from FTS5 table, order by rank title, section, & file name
// Return results as slice of struct SearchAllResult, later prepped for Alfred script filter
func searchAll(sess *dbr.Session, query string) ([]SearchAllResult, error) {
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

// Test if string is included in slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
