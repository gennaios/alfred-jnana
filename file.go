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

type File struct {
	ID            int64
	Path          string         `db:"path"`
	FileName      string         `db:"file_name"`
	FileExtension string         `db:"file_extension"`
	Title         dbr.NullString `db:"title"`
	Authors       dbr.NullString `db:"authors"`
	Subjects      dbr.NullString `db:"subjects"`
	DateCreated   string         `db:"date_created"`
	DateModified  string         `db:"file_modified_date"` // TODO: rename -> date_modified
	FileHash      string         `db:"hash"`
}

func (db *Database) GetFile(book string) (File, bool, error) {
	var file File
	var err error
	var hash string
	created := false // check for update if found
	changed := false // return value, to later recheck bookmarks

	// first lookup by file path
	file, err = db.GetFileFromPath(book)

	// not found, possible file moved, look up by file hash
	if err == dbr.ErrNotFound {
		hash, err = fileHash(book)
		file, err = db.GetFileFromHash(hash)
	} else {
		created = false
	}

	// found by hash, verify not dupe
	if file.ID != 0 && book != file.Path {
		if _, err = os.Stat(file.Path); err != nil {
			if os.IsNotExist(err) {
				// old path doesn't exist, moved, same hash
				_ = notification("File moved: " + file.Path)
				file.Path = book
				err = db.UpdateFile(file)
				// hash match, no changes needed
				return file, false, err
			}
		} else {
			_ = notification("Dupe of: " + file.Path)
		}
	}

	// not found by path or hash, create new
	if err == dbr.ErrNotFound {
		file, err = db.NewFile(book)
		if err != nil {
			return file, true, err
		}
		created = true
		_ = notification("New Book, ID: " + string(file.ID))
	}

	// not created, check if different
	if created == false {
		// check book changed against date in database
		stat, _ := os.Stat(book)
		modDate := stat.ModTime().UTC()
		oldDate, _ := time.Parse(time.RFC3339, file.DateModified)
		if err != nil {
			fmt.Println("date error:", err)
		}
		diff := modDate.Sub(oldDate).Seconds()

		if diff > 1 {
			//date different, check hash value
			changed = true
			file.FileHash, _ = fileHash(book)
			file.DateModified = modDate.String()
			err = db.UpdateFile(file)
		}
	}
	return file, changed, err
}

// Get: Look for existing record by file path
func (db *Database) GetFileFromPath(book string) (File, error) {
	var file File
	err := db.sess.Select("*").From("files").Where("path = ?", book).LoadOne(&file)
	return file, err
}

// GetFromHash: look for existing by file hash (sha256)
func (db *Database) GetFileFromHash(hash string) (File, error) {
	var file File
	err := db.sess.Select("*").From("files").Where("hash = ?", hash).LoadOne(&file)
	return file, err
}

// NewFile create new file entry.
// File struct comes in with only path.
func (db *Database) NewFile(book string) (File, error) {
	// path, name, extension, created, modified, hash
	stat, err := os.Stat(book)
	if err != nil {
		return File{}, err
	}
	hash, _ := fileHash(book)
	// format string for insert, strange set then get by format doesn't work
	// TODO: format not compatible with alfred-gnosis
	dateModified := stat.ModTime().UTC().Format(time.RFC3339)
	dateModified = strings.Replace(dateModified, " +0000 UTC", "", 1)

	tx, err := db.sess.Begin()
	if err != nil {
		return File{}, err
	}
	defer tx.RollbackUnlessCommitted()
	// filepath.Ext returns with dot
	_, err = tx.InsertInto("files").
		Pair("path", book).
		Pair("file_name", filepath.Base(book)).
		Pair("file_extension", filepath.Ext(book)[1:]).
		Pair("date_created", time.Now().UTC().Format(time.RFC3339)).
		Pair("file_modified_date", dateModified).
		Pair("hash", hash).
		Exec()
	err = tx.Commit()
	file, err := db.GetFileFromPath(book)
	return file, err
}

func (db *Database) UpdateFile(file File) error {
	// format string for insert, strange set then get by format doesn't work
	dateModified := strings.Replace(file.DateModified, " +0000 UTC", "", 1)

	tx, err := db.sess.Begin()
	_, err = db.sess.Update("files").
		Set("file_name", filepath.Base(file.Path)).
		Set("path", file.Path).
		Set("hash", file.FileHash).
		Set("file_modified_date", dateModified).
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
