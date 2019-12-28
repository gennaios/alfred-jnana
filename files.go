package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/campoy/unique"
	"github.com/djherbis/times"
	"github.com/gocraft/dbr/v2"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type DatabaseFile struct {
	ID            int64
	Path          string         `db:"path"`
	FileName      string         `db:"file_name"`
	FileExtension string         `db:"file_extension"`
	FileSize      int64          `db:"file_size"`
	Title         dbr.NullString `db:"file_title"`
	Authors       dbr.NullString `db:"file_authors"`
	Subjects      dbr.NullString `db:"file_subjects"`
	Publisher     dbr.NullString `db:"file_publisher"`
	Language      dbr.NullString `db:"language"`
	Description   dbr.NullString `db:"description"`
	DateCreated   string         `db:"date_created"`
	DateModified  string         `db:"date_modified"`
	DateAccessed  dbr.NullString `db:"date_accessed"`
	Rating        dbr.NullInt64  `db:"rating"`
	FileHash      string         `db:"hash"`
}

// AllFiles: get all files, used for update
func (db *Database) AllFiles() ([]*DatabaseFile, error) {
	var files []*DatabaseFile
	_, err := db.sess.Select("*").From("files").Load(&files)
	return files, err
}

// CoverForFile creates thumbnails using ImageMagick
// 160x160 - 80x80 2x
// 7.x: magick convert -resize 160x160 -background transparent -colorspace srgb -depth 8 -gravity center -extent 160x160 -strip  …
func (db *Database) CoverForFile(fileRecord *DatabaseFile, coversCacheDir string) bool {
	var err error
	coverPath := filepath.Join(coversCacheDir, strconv.FormatInt(fileRecord.ID, 10)+".png")

	_, err = os.Stat(coverPath)
	if !os.IsNotExist(err) {
		return true
	} else {
		_ = notification("File ID: " + strconv.FormatInt(fileRecord.ID, 10) + ".png")
	}

	// create thumbnail

	// check for new file
	_, errFile := os.Stat(coverPath)
	if !os.IsNotExist(errFile) {
		return true
	} else {
		return false
	}
}

func (db *Database) GetFile(book string, check bool) (*DatabaseFile, bool, error) {
	var file *DatabaseFile
	var err error
	var hash string
	changed := false // return value, to later recheck bookmarks

	// first lookup by file path
	file, err = db.GetFileFromPath(book)

	// not found, possible file moved, look up by file hash
	if err == dbr.ErrNotFound {
		hash, err = fileHash(book)
		file, err = db.GetFileFromHash(hash)
	}

	// not found by path or hash, create new
	if err == dbr.ErrNotFound {
		file, err = db.NewFile(book)
		if err != nil {
			return file, true, err
		}
	} else if file.ID != 0 && book != file.Path {
		// found by hash, verify not dupe
		if _, err = os.Stat(file.Path); err != nil {
			if os.IsNotExist(err) {
				// old path doesn't exist, moved, same hash
				_ = notification("File moved: " + file.Path)
				file.Path = book
				err = db.UpdateFile(*file)
				// hash match, no changes needed
				return file, false, err
			}
		} else {
			// check file exists, or notification will be triggered if not
			if _, err = os.Stat(book); err != nil {
				if !os.IsNotExist(err) {
					_ = notification("Dupe of: " + file.Path)
					check = false
				}
			}
		}
	}

	// not created, check if different
	// NOTE: run each time with file and file filtered bookmarks, try to speed up
	if check == true {
		// check book changed against date in database
		stat, err := os.Stat(book)
		if err != nil {
			return file, false, err
		}
		modDate := stat.ModTime().UTC().Truncate(time.Second)
		oldDate, err := time.Parse(time.RFC3339, file.DateModified)
		if err != nil {
			fmt.Println("date error:", err)
		}

		if modDate.After(oldDate) {
			//date different, check hash value
			changed = true
			file.FileHash, _ = fileHash(book)
			file.DateModified = modDate.Format("2006-01-02 15:04:05")
			err = db.UpdateFile(*file)
		}
	}
	return file, changed, err
}

// GetFileFromPath: Look for existing record by file path
// return columns needed by GetFile, all in case of update
func (db *Database) GetFileFromPath(book string) (*DatabaseFile, error) {
	var file *DatabaseFile
	err := db.sess.SelectBySql("SELECT * FROM files WHERE path = ?", book).LoadOne(&file)
	return file, err
}

// GetFromHash: look for existing by file hash (sha256)
// return columns needed by GetFile, all in case of update
func (db *Database) GetFileFromHash(hash string) (*DatabaseFile, error) {
	var file *DatabaseFile
	err := db.sess.SelectBySql("SELECT * FROM files WHERE hash = ?", hash).LoadOne(&file)
	return file, err
}

// NewFile create new file entry.
// DatabaseFile struct comes in with only path.
// Required fields: path, name, extension, created, modified, hash
func (db *Database) NewFile(book string) (*DatabaseFile, error) {
	stat, err := os.Stat(book)
	if err != nil {
		return &DatabaseFile{}, err
	}
	// format string for insert, strange set then get by format doesn't work
	dateModified := stat.ModTime().UTC().Format("2006-01-02 15:04:05")
	hash, _ := fileHash(book)

	f := File{}
	if err = f.Init(book); err != nil {
		return &DatabaseFile{}, err
	}

	t, err := times.Stat(book)
	if err != nil {
		log.Fatal(err.Error())
	}
	dateCreated := ""
	if t.HasBirthTime() {
		dateCreated = t.BirthTime().UTC().Format("2006-01-02 15:04:05")
	}

	tx, err := db.sess.Begin()
	if err != nil {
		return &DatabaseFile{}, err
	}
	defer tx.RollbackUnlessCommitted()

	// filepath.Ext returns with dot
	_, err = tx.InsertInto("files").
		Pair("path", book).
		Pair("file_name", filepath.Base(book)).
		Pair("file_extension", strings.ToLower(filepath.Ext(book)[1:])).
		Pair("file_size", stat.Size()).
		Pair("file_title", NewNullString(f.title)).
		Pair("file_authors", NewNullString(f.authors)).
		Pair("file_subjects", NewNullString(f.subjects)).
		Pair("file_publisher", NewNullString(f.publisher)).
		Pair("date_created", dateCreated).
		Pair("date_modified", dateModified).
		Pair("date_accessed", NewNullString(t.AccessTime().UTC().Format("2006-01-02 15:04:05"))).
		Pair("hash", hash).
		Exec()

	if err != nil {
		// TODO: workaround PDF metadata issue
		_, err = tx.InsertInto("files").
			Pair("path", book).
			Pair("file_name", filepath.Base(book)).
			Pair("file_size", stat.Size()).
			Pair("file_extension", filepath.Ext(book)[1:]).
			Pair("date_created", dateCreated).
			Pair("date_modified", dateModified).
			Pair("date_accessed", NewNullString(t.AccessTime().UTC().Format("2006-01-02 15:04:05"))).
			Pair("hash", hash).
			Exec()
	}

	err = tx.Commit()

	//err = f.file.Close() // TODO: invalid memory address or nil pointer, maybe closed in File?
	if strings.HasSuffix(book, ".epub") {
		f.epub.Close()
	}

	file, err := db.GetFileFromPath(book)
	return file, err
}

// RecentFiles: list of recently opened files
func (db *Database) RecentFiles() ([]*DatabaseFile, error) {
	var files []*DatabaseFile
	_, err := db.sess.Select("*").From("files").OrderDesc("date_accessed").Limit(50).Load(&files)
	return files, err
}

// SearchFiles: Search all files from FTS5 table,
// order by rank: file_name, file_title, file_authors, file_subjects, file_publisher, description
// Return results as slice of struct DatabaseFile, later prepped for Alfred script filter
func (db *Database) SearchFiles(query string) ([]*DatabaseFile, error) {
	queryString := stringForSQLite(query)
	var results []*DatabaseFile

	// NOTE: AND rank MATCH 'bm25(…)' ORDER BY rank faster than ORDER BY bm25(fts, …)
	_, err := db.sess.SelectBySql(`SELECT
			files.id, files.path, files.file_name, files.file_extension, files.file_title, files.file_subjects
			FROM files
			JOIN filesindex on files.id = filesindex.rowid
			WHERE filesindex MATCH ?
			AND rank MATCH 'bm25(10.0, 2.0, 2.0, 2.0, 2.0, 2.0)'
			ORDER BY rank LIMIT 200`,
		queryString).Load(&results)

	err = db.conn.Close()
	return results, err
}

// UpdateFile: update file on change of path, file name, or date modified
func (db *Database) UpdateFile(file DatabaseFile) error {
	stat, err := os.Stat(file.Path)
	if err != nil {
		return err
	}
	t, err := times.Stat(file.Path)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.DateAccessed.String = t.AccessTime().UTC().Format("2006-01-02 15:04:05")

	tx, err := db.sess.Begin()

	_, err = db.sess.UpdateBySql(`UPDATE files SET
			path = ?, file_name = ?, file_size = ?,
			file_title = ?, file_authors = ?, file_subjects = ?, file_publisher = ?,
			date_modified = ?, date_accessed = ?, hash = ?
			WHERE id = ?`,
		file.Path, filepath.Base(file.Path), stat.Size(),
		NewNullString(file.Title.String), NewNullString(file.Authors.String), NewNullString(file.Subjects.String), NewNullString(file.Publisher.String),
		file.DateModified, NewNullString(file.DateAccessed.String), file.FileHash,
		file.ID).Exec()

	if err != nil {
		// TODO: PDF metadata "unrecognized token", workaround don't update metadata
		_, err = db.sess.UpdateBySql(`UPDATE files SET
			path = ?, file_name = ?, file_size = ?, date_modified = ?, date_accessed = ?, hash = ?
			WHERE id = ?`,
			file.Path, filepath.Base(file.Path), stat.Size(), file.DateModified, NewNullString(file.DateAccessed.String),
			file.FileHash, file.ID).Exec()
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error committing:", err)
	}
	return err
}

// UpdateMetadata: check for updates to metadata
func (db *Database) UpdateMetadata(file *DatabaseFile) (bool, error) {
	var err error
	if _, err = os.Stat(file.Path); err != nil {
		return false, err
	}
	update := false

	f := File{}
	if err = f.Init(file.Path); err != nil {
		return false, err
	}

	if file.Title.String != f.title && f.title != "" {
		file.Title.String = f.title
		update = true
	}
	if file.Authors.String != f.authors && f.authors != "" {
		file.Authors.String = f.authors
		update = true
	}
	if strings.HasSuffix(file.Path, "epub") {
		if file.Publisher.String != f.publisher && f.publisher != "" {
			file.Publisher.String = f.publisher
			update = true
		}
	}

	if update == true {
		err = db.UpdateFile(*file)
		if err == nil {
			return true, err
		}
	}

	return false, err
}

// UpdateSubjects: set subjects/keywords for file
func (db *Database) UpdateSubjects(file *DatabaseFile, subjects string) error {
	var err error

	if subjects == "" {
		return err
	}

	terms := strings.Split(strings.ToLower(subjects), ",")
	s := trimMetadata(terms)

	less := lessString(&s)
	unique.Slice(&s, less)

	newSubjects := strings.Join(s, ", ")

	if newSubjects != file.Subjects.String {
		file.Subjects.String = newSubjects
		err = db.UpdateFile(*file)
	}

	return err
}

func lessString(v interface{}) func(i, j int) bool {
	s := *v.(*[]string)
	return func(i, j int) bool { return s[i] < s[j] }
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

// fileHash: create sha256 file hash for later comparison
func fileHash(file string) (string, error) {
	if _, err := os.Stat(file); err != nil {
		return "", err
	}
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	hashString := hex.EncodeToString(h.Sum(nil))
	return hashString, err
}

func NewNullString(s string) dbr.NullString {
	if len(s) == 0 {
		return dbr.NullString{}
	}
	return dbr.NewNullString(s)
}
