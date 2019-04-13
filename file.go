package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DatabaseFile struct {
	ID            int64
	Path          string         `db:"path"`
	FileName      string         `db:"file_name"`
	FileExtension string         `db:"file_extension"`
	Title         dbr.NullString `db:"file_title"`
	Authors       dbr.NullString `db:"file_authors"`
	Subjects      dbr.NullString `db:"file_subjects"`
	Publisher     dbr.NullString `db:"file_publisher"`
	DateCreated   string         `db:"date_created"`
	DateModified  string         `db:"date_modified"`
	FileHash      string         `db:"hash"`
}

// AllFiles: get all files, used for update
func (db *Database) AllFiles() ([]*DatabaseFile, error) {
	var files []*DatabaseFile
	_, err := db.sess.Select("*").From("files").Load(&files)
	return files, err
}

func (db *Database) GetFile(book string, check bool) (*DatabaseFile, bool, error) {
	var file *DatabaseFile
	var err error
	var hash string
	changed := false // return value, to later recheck bookmarks

	// first lookup by path path
	file, err = db.GetFileFromPath(book)

	// not found, possible path moved, look up by path hash
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
			// check path exists, or notification will be triggered if not
			if _, err = os.Stat(book); err != nil {
				if !os.IsNotExist(err) {
					_ = notification("Dupe of: " + file.Path)
					check = false
				}
			}
		}
	}

	// not created, check if different
	// NOTE: run each time with path and path filtered bookmarks, try to speed up
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

// GetFileFromPath: Look for existing record by path path
// return columns needed by GetFile, all in case of update
func (db *Database) GetFileFromPath(book string) (*DatabaseFile, error) {
	var file *DatabaseFile
	err := db.sess.Select("*").
		From("files").Where("path = ?", book).LoadOne(&file)
	return file, err
}

// GetFromHash: look for existing by path hash (sha256)
// return columns needed by GetFile, all in case of update
func (db *Database) GetFileFromHash(hash string) (*DatabaseFile, error) {
	var file *DatabaseFile
	err := db.sess.Select("*").
		From("files").Where("hash = ?", hash).LoadOne(&file)
	return file, err
}

// NewFile create new path entry.
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
	err = f.Init(book)

	tx, err := db.sess.Begin()
	if err != nil {
		return &DatabaseFile{}, err
	}
	defer tx.RollbackUnlessCommitted()
	// filepath.Ext returns with dot
	_, err = tx.InsertInto("files").
		Pair("path", book).
		Pair("file_name", filepath.Base(book)).
		Pair("file_extension", filepath.Ext(book)[1:]).
		Pair("file_title", dbr.NewNullString(f.title)).
		Pair("file_authors", dbr.NewNullString(f.authors)).
		Pair("file_subjects", dbr.NewNullString(f.subjects)).
		Pair("file_publisher", dbr.NewNullString(f.publisher)).
		Pair("date_created", time.Now().UTC().Format("2006-01-02 15:04:05")).
		Pair("date_modified", dateModified).
		Pair("hash", hash).
		Exec()
	err = tx.Commit()

	if strings.HasSuffix(book, ".pdf") {
		err = f.pdf.Close()
	} else {
		f.epub.Close()
	}

	file, err := db.GetFileFromPath(book)
	return file, err
}

// UpdateFile: update path on change of path, path name, or date modified
func (db *Database) UpdateFile(file DatabaseFile) error {
	tx, err := db.sess.Begin()
	_, err = db.sess.Update("files").
		Set("path", file.Path).
		Set("file_name", filepath.Base(file.Path)).
		Set("file_title", dbr.NewNullString(file.Title.String)).
		Set("file_authors", dbr.NewNullString(file.Authors.String)).
		Set("file_subjects", dbr.NewNullString(file.Subjects.String)).
		Set("file_publisher", dbr.NewNullString(file.Publisher.String)).
		Set("date_modified", file.DateModified).
		Set("hash", file.FileHash).
		Where("id = ?", file.ID).
		Exec()
	if err != nil {
		fmt.Println("error updating:", err)
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
	err = f.Init(file.Path)
	if err != nil {
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
		if file.Subjects.String != f.subjects && f.subjects != "" {
			file.Subjects.String = f.subjects
			update = true
		}
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

func fileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

// fileHash: create sha256 path hash for later comparison
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
