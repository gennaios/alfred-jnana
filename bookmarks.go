package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/deckarep/gosx-notifier"
	"github.com/gocraft/dbr"
	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
)

type BookmarkRecord struct {
	ID          int64
	FileId      int64          `db:"file_id"`
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

type Database struct {
	conn *dbr.Connection
	sess *dbr.Session
}

// Get: Look for existing record by file path
func (db *Database) GetBookFromFile(book string) (File, error) {
	var file File
	err := db.sess.Select("*").From("files").Where("path = ?", book).LoadOne(&file)
	return file, err
}

// GetFromHash: look for existing by file hash (sha256)
func (db *Database) GetBookFromHash(hash string) (File, error) {
	var file File
	err := db.sess.Select("*").From("files").Where("hash = ?", hash).LoadOne(&file)
	return file, err
}

func (db *Database) NewBook(file File) error {
	tx, err := db.sess.Begin()
	file.FileName = filepath.Base(file.Path)
	file.FileExtension = filepath.Ext(file.Path)
	file.DateCreated = time.Now().UTC().Format("time.RFC3339")

	stat, err := os.Stat(file.Path)
	if err != nil {
		file.DateModified = stat.ModTime().UTC().Format("time.RFC3339")
	}
	_, err = db.sess.InsertInto("files").
		Columns("id", "path", "file_name", "file_extension", "date_created", "hash").
		Record(&file).
		Exec()
	err = tx.Commit()
	return err
}

func (db *Database) UpdateBook(file File) error {
	tx, err := db.sess.Begin()
	_, err = db.sess.Update("files").
		Set("file_name", filepath.Base(file.Path)).
		Set("path", file.Path).
		Set("hash", file.FileHash).
		Set("file_modified_date", file.DateModified).
		Where("id = ?", file.ID).
		Exec()
	if err != nil {
		fmt.Println("error updating:", err)
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("error updating:", err)
	}
	return err
}

func (db *Database) UpdateBookmarks(file File, bookmarks []Bookmark) error {
	tx, err := db.sess.Begin()
	_, err = db.sess.DeleteFrom("bookmarks").
		Where("file_id = ?", file.ID).
		Exec()
	// insert new bookmarks
	for i := 0; i < len(bookmarks); i++ {
		_, err = db.sess.InsertInto("bookmarks").
			Pair("file_id", file.ID).
			Pair("title", bookmarks[i].Title).
			Pair("section", dbr.NewNullString(bookmarks[i].Section)).
			Pair("destination", bookmarks[i].Destination).
			Exec()
	}
	err = tx.Commit()
	return err
}

func (db *Database) Init(filepath string) {
	conn, err := dbr.Open("sqlite3", filepath, nil)
	if err != nil {
		panic(err)
	}
	if conn == nil {
		panic("db nil")
	}
	db.conn = conn

	_, _ = conn.Exec("PRAGMA auto_vacuum = 1")
	_, _ = conn.Exec("PRAGMA foreign_keys = 1")
	_, _ = conn.Exec("PRAGMA ignore_check_constraints = 0")
	_, _ = conn.Exec("PRAGMA journal_mode = WAL")
	_, _ = conn.Exec("PRAGMA synchronous = 0")
	_, _ = conn.Exec("PRAGMA temp_store = 2") // MEMORY

	_, _ = conn.Exec("PRAGMA cache_size = -31250")
	_, _ = conn.Exec("PRAGMA page_size = 8192") // default 4096, match APFS block size?

	sess := conn.NewSession(nil)
	db.sess = sess
	_, err = sess.Begin()
	if err != nil {
		panic(err)
	}
}

func (db *Database) GetFile(book string) (File, bool, error) {
	file := File{ID: 0, Path: book}
	var err error
	var hash string
	created := true  // check for update if found
	changed := false // return value, to later recheck bookmarks

	// first lookup by file path
	file, err = db.GetBookFromFile(book)

	// not found, possible file moved, look up by file hash
	if err == dbr.ErrNotFound {
		hash, err = fileHash(book)
		file, err = db.GetBookFromHash(hash)
	} else {
		created = false
	}

	// found by hash, verify not dupe
	// TODO move not working
	// TODO: dupe not working, move maybe ok
	if file.ID != 0 && book != file.Path {
		if _, err = os.Stat(file.Path); err != nil {
			if os.IsNotExist(err) {
				// old path doesn't exist, moved, same hash
				_ = notification("File moved: " + file.Path)
				file.Path = book
				err = db.UpdateBook(file)
				return file, false, err
			}
		} else {
			_ = notification("Dupe of: " + file.Path)
		}
	}

	// not found by path or hash, create new record
	if file.ID == 0 {
		file.Path = book
		file.FileHash, err = fileHash(book)
		stat, _ := os.Stat(book)
		modDate := stat.ModTime().UTC()
		file.DateModified = modDate.String()
		err = db.NewBook(file) // TODO: new book
		_ = notification("New Book.")
	}

	// not created, check if different
	if created == false {
		// check book changed against date in database
		stat, _ := os.Stat(book)
		modDate := stat.ModTime().UTC()
		oldDate, _ := time.Parse(time.RFC3339, file.DateModified)
		diff := modDate.Sub(oldDate).Seconds()

		if diff > 1 {
			//date different, check hash value
			changed = true
			file.FileHash, _ = fileHash(book)
			file.DateModified = modDate.String()
			err = db.UpdateBook(file)
		}
	}
	return file, changed, err
}

func (db *Database) BookmarksForFile(file string) ([]BookmarkRecord, error) {
	var bookmarks []BookmarkRecord
	var err error

	fileRecord, changed, _ := db.GetFile(file)
	err = db.sess.Select("id", "title", "section", "destination").From("bookmarks").
		Where("file_id == ?", fileRecord.ID).
		LoadOne(&bookmarks)

	// TODO, file created or changed, compared bookmarks
	if changed == true {
		var newBookmarks []Bookmark

		// PDF get bookmarks from file
		if strings.HasSuffix(file, ".pdf") {
			// TODO: call pdf module
			newBookmarks, _ = bookmarksForPDF(file)
		}
		// TODO: EPUB get bookmarks from file

		// file updated, compare bookmarks
		if bookmarksEqual(bookmarks, newBookmarks) == false {
			// update database
			err = db.UpdateBookmarks(fileRecord, newBookmarks)
			_ = notification("Bookmarks updated.")
			// fetch new bookmarks
		}
	}
	return bookmarks, err
}

// Filtered bookmarks for file
func (db *Database) BookmarksForFileFiltered(file string, query string) ([]SearchAllResult, error) {
	queryString := stringForSQLite(query)
	var results []SearchAllResult

	fileRecord, _, _ := db.GetFile(file)

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
	_, err := db.sess.SelectBySql("select bookmarks.id, bookmarks.title, bookmarks.section, bookmarks.destination from bookmarks JOIN bookmarksindex on bookmarks.id = bookmarksindex.rowid where bookmarks.file_id = " + strconv.FormatInt(fileRecord.ID, 10) + " and bookmarksindex.rowid = bookmarks.id and bookmarksindex match '{title section} : " + queryString + "' ORDER BY 'bm25(bookmarksindex, 5.0, 2.0)';").Load(&results)
	return results, err
}

// Search all bookmarks from FTS5 table, order by rank title, section, & file name
// Return results as slice of struct SearchAllResult, later prepped for Alfred script filter
func (db *Database) searchAll(query string) ([]SearchAllResult, error) {
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
	_, err := db.sess.Select("bookmarks.id", "bookmarks.title", "bookmarks.section",
		"bookmarks.destination", "bookmarks.file_id", "files.path", "files.file_name").
		From("bookmarks").
		Join("files", "bookmarks.file_id = files.id").
		Join("bookmarksindex", "bookmarks.id = bookmarksindex.rowid").
		Where("bookmarksindex MATCH ? AND rank MATCH 'bm25(5.0, 2.0, 1.0)'", queryString).
		OrderBy("rank").Limit(100).Load(&results)
	return results, err
}

// Compare bookmarks from database with file
func bookmarksEqual(bookmarks []BookmarkRecord, newBookmarks []Bookmark) bool {
	var oldBookmarks []Bookmark

	for i := 0; i < len(bookmarks); i++ {
		var bookmark Bookmark
		bookmark.Title = bookmarks[i].Title
		bookmark.Section = bookmarks[i].Section.String
		bookmark.Destination = bookmarks[i].Destination
		oldBookmarks = append(oldBookmarks, bookmark)
	}

	if cmp.Equal(oldBookmarks, newBookmarks) {
		return true
	} else {
		return false
	}
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

func fileHash(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// sha256 hash of open file
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	hashString := hex.EncodeToString(h.Sum(nil))
	return hashString, err
}

func notification(message string) error {
	note := gosxnotifier.NewNotification(message)
	note.Title = "Jnana"
	note.Sound = gosxnotifier.Default

	if err := note.Push(); err != nil {
		return err
	}
	return nil
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
