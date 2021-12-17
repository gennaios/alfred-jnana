package main

import (
	. "jnana/internal"

	"fmt"
	"github.com/campoy/unique"
	"github.com/djherbis/times"
	"github.com/gocraft/dbr/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type DatabaseFile struct {
	ID           int64
	Path         string         `db:"path"`
	Name         string         `db:"name"`
	Extension    string         `db:"extension"`
	Size         int64          `db:"size"`
	Title        dbr.NullString `db:"title"`
	Creator      dbr.NullString `db:"creator"`
	Subject      dbr.NullString `db:"subject"`
	Publisher    dbr.NullString `db:"publisher"`
	Language     dbr.NullString `db:"language"`
	Description  dbr.NullString `db:"description"`
	DateCreated  string         `db:"date_created"`
	DateModified string         `db:"date_modified"`
	DateAccessed dbr.NullString `db:"date_accessed"`
	Rating       dbr.NullInt64  `db:"rating"`
	FileHash     string         `db:"hash"`
}

// AllFiles get all files, used for update
func (db *Database) AllFiles() ([]*DatabaseFile, error) {
	var files []*DatabaseFile
	_, err := db.sess.Select("*").From("file").Load(&files)
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
		hash, err = FileHash(book)
		file, err = db.GetFileFromHash(hash)
	}

	// not found by path or hash, create new
	if err == dbr.ErrNotFound {
		file, err = db.NewFile(book)
		if err != nil {
			return file, true, err
		}
	} else if file != nil {
		if file.ID != 0 && book != file.Path {
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
	}

	// not created, check if different
	// NOTE: run each time with file and file filtered bookmarks, try to speed up
	if check == true && file != nil {
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
			file.FileHash, _ = FileHash(book)
			file.DateModified = modDate.Format("2006-01-02 15:04:05")
			err = db.UpdateFile(*file)
		}
	}
	return file, changed, err
}

// GetFileFromPath Look for existing record by file path
// return columns needed by GetFile, all in case of update
func (db *Database) GetFileFromPath(book string) (*DatabaseFile, error) {
	var file *DatabaseFile
	err := db.sess.SelectBySql("SELECT * FROM file WHERE path = ?", book).LoadOne(&file)
	return file, err
}

// GetFileFromHash look for existing by file hash (sha256)
// return columns needed by GetFile, all in case of update
func (db *Database) GetFileFromHash(hash string) (*DatabaseFile, error) {
	var file *DatabaseFile
	err := db.sess.SelectBySql("SELECT * FROM file WHERE hash = ?", hash).LoadOne(&file)
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
	hash, _ := FileHash(book)

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
	_, err = tx.InsertInto("file").
		Pair("path", book).
		Pair("name", filepath.Base(book)).
		Pair("extension", strings.ToLower(filepath.Ext(book)[1:])).
		Pair("size", stat.Size()).
		Pair("title", NewNullString(f.title)).
		Pair("creator", NewNullString(f.creator)).
		Pair("subject", NewNullString(f.subject)).
		Pair("publisher", NewNullString(f.publisher)).
		Pair("date_created", dateCreated).
		Pair("date_modified", dateModified).
		Pair("date_accessed", NewNullString(t.AccessTime().UTC().Format("2006-01-02 15:04:05"))).
		Pair("hash", hash).
		Exec()

	if err != nil {
		// TODO: workaround PDF metadata issue
		_, err = tx.InsertInto("file").
			Pair("path", book).
			Pair("name", filepath.Base(book)).
			Pair("size", stat.Size()).
			Pair("extension", filepath.Ext(book)[1:]).
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

// RecentFiles list of recently opened files
func (db *Database) RecentFiles() ([]*DatabaseFile, error) {
	var files []*DatabaseFile
	_, err := db.sess.Select("*").From("file").OrderDesc("date_accessed").Limit(50).Load(&files)
	return files, err
}

// SearchFiles Search all files from FTS5 table,
// order by rank: name, title, creator, subject, publisher, description
// Return results as slice of struct DatabaseFile, later prepped for Alfred script filter
func (db *Database) SearchFiles(query string) ([]*DatabaseFile, error) {
	queryString := stringForSQLite(query)
	var results []*DatabaseFile

	// NOTE: AND rank MATCH 'bm25(…)' ORDER BY rank faster than ORDER BY bm25(fts, …)
	_, err := db.sess.SelectBySql(`SELECT
			file.id, file.path, file.name, file.extension, file.title, file.subject
			FROM file
			JOIN file_search on file.id = file_search.rowid
			WHERE file_search MATCH ?
			AND rank MATCH 'bm25(10.0, 2.0, 2.0, 2.0, 2.0, 2.0)'
			ORDER BY rank LIMIT 200`,
		queryString).Load(&results)

	err = db.conn.Close()
	return results, err
}

// UpdateFile update file on change of path, file name, or date modified
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

	_, err = db.sess.UpdateBySql(`UPDATE file SET
			path = ?, name = ?, size = ?,
			title = ?, creator = ?, subject = ?, publisher = ?,
			date_modified = ?, date_accessed = ?, hash = ?
			WHERE id = ?`,
		file.Path, filepath.Base(file.Path), stat.Size(),
		NewNullString(file.Title.String), NewNullString(file.Creator.String), NewNullString(file.Subject.String), NewNullString(file.Publisher.String),
		file.DateModified, NewNullString(file.DateAccessed.String), file.FileHash,
		file.ID).Exec()

	if err != nil {
		// TODO: PDF metadata "unrecognized token", workaround don't update metadata
		_, err = db.sess.UpdateBySql(`UPDATE file SET
			path = ?, name = ?, size = ?, date_modified = ?, date_accessed = ?, hash = ?
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

// UpdateDateAccessed update last opened
func (db *Database) UpdateDateAccessed(file *DatabaseFile) {
	currentTime := time.Now().UTC().Format("2006-01-02 15:04:05.000")
	_, _ = db.sess.UpdateBySql(`UPDATE file SET date_accessed = ? WHERE id = ?`,
		NewNullString(currentTime), file.ID).Exec()
}

// UpdateMetadata check for updates to metadata
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
	if file.Creator.String != f.creator && f.creator != "" {
		file.Creator.String = f.creator
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

// UpdateSubject set subject/keywords for file
func (db *Database) UpdateSubject(file *DatabaseFile, subject string) error {
	var err error

	if subject == "" {
		return err
	}

	terms := strings.Split(strings.ToLower(subject), ",")
	s := trimMetadata(terms)

	less := lessString(&s)
	unique.Slice(&s, less)

	newSubject := strings.Join(s, ", ")

	if newSubject != file.Subject.String {
		file.Subject.String = newSubject
		err = db.UpdateFile(*file)
	}

	return err
}

func lessString(v interface{}) func(i, j int) bool {
	s := *v.(*[]string)
	return func(i, j int) bool { return s[i] < s[j] }
}

func NewNullString(s string) dbr.NullString {
	if len(s) == 0 {
		return dbr.NullString{}
	}
	return dbr.NewNullString(s)
}
