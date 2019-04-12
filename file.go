package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
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
	Title         dbr.NullString `db:"file_title"`
	Authors       dbr.NullString `db:"file_authors"`
	Subjects      dbr.NullString `db:"file_subjects"`
	Publisher     dbr.NullString `db:"file_publisher"`
	DateCreated   string         `db:"date_created"`
	DateModified  string         `db:"date_modified"`
	FileHash      string         `db:"hash"`
}

// AllFiles: get all files, used for update
func (db *Database) AllFiles() ([]*File, error) {
	var files []*File
	_, err := db.sess.Select("*").From("files").Load(&files)
	return files, err
}

func (db *Database) GetFile(book string) (*File, bool, error) {
	var file *File
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
				err = db.UpdateFile(*file)
				// hash match, no changes needed
				return file, false, err
			}
		} else {
			// check file exists, or notification will be triggered if not
			if _, err = os.Stat(book); err != nil {
				if !os.IsNotExist(err) {
					_ = notification("Dupe of: " + file.Path)
				}
			}
		}
	}

	// not found by path or hash, create new
	if err == dbr.ErrNotFound {
		file, err = db.NewFile(book)
		if err != nil {
			return file, true, err
		}
		created = true
	}

	// not created, check if different
	if created == false {
		// check book changed against date in database
		stat, err := os.Stat(book)
		if err != nil {
			return file, false, err
		}
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
			file.DateModified = modDate.Format("2006-01-02 15:04:05")
			err = db.UpdateFile(*file)
		}
	}
	return file, changed, err
}

// GetFileFromPath: Look for existing record by file path
// return columns needed by GetFile
func (db *Database) GetFileFromPath(book string) (*File, error) {
	var file *File
	err := db.sess.Select("id", "path", "date_modified").
		From("files").Where("path = ?", book).LoadOne(&file)
	return file, err
}

// GetFromHash: look for existing by file hash (sha256)
// return columns needed by GetFile
func (db *Database) GetFileFromHash(hash string) (*File, error) {
	var file *File
	err := db.sess.Select("id", "path", "date_modified").
		From("files").Where("hash = ?", hash).LoadOne(&file)
	return file, err
}

// NewFile create new file entry.
// File struct comes in with only path.
// Required fields: path, name, extension, created, modified, hash
func (db *Database) NewFile(book string) (*File, error) {
	stat, err := os.Stat(book)
	if err != nil {
		return &File{}, err
	}
	// format string for insert, strange set then get by format doesn't work
	dateModified := stat.ModTime().UTC().Format("2006-01-02 15:04:05")
	hash, _ := fileHash(book)

	// TODO: get file metadata

	tx, err := db.sess.Begin()
	if err != nil {
		return &File{}, err
	}
	defer tx.RollbackUnlessCommitted()
	// filepath.Ext returns with dot
	_, err = tx.InsertInto("files").
		Pair("path", book).
		Pair("file_name", filepath.Base(book)).
		Pair("file_extension", filepath.Ext(book)[1:]).
		Pair("date_created", time.Now().UTC().Format("2006-01-02 15:04:05")).
		Pair("date_modified", dateModified).
		Pair("hash", hash).
		Exec()
	err = tx.Commit()
	file, err := db.GetFileFromPath(book)
	return file, err
}

// UpdateFile: update file on change of path, file name, or date modified
func (db *Database) UpdateFile(file File) error {
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

// UpdateFileCheck: check for updates to metadata
func (db *Database) UpdateFileCheck(file *File) (bool, error) {
	var err error
	if _, err = os.Stat(file.Path); err != nil {
		return false, err
	}
	update := false

	if strings.HasSuffix(file.Path, "pdf") {
		metadata := MetadataForPDF(file.Path)

		if file.Title.String != metadata.Title && metadata.Title != "" {
			file.Title.String = metadata.Title
			update = true
		}
		if file.Authors.String != metadata.Authors && metadata.Authors != "" {
			file.Authors.String = metadata.Authors
			update = true
		}
	} else {
		epub := Epub{}
		if err = epub.Init(file.Path); err != nil {
			log.Println("Error: ", file.FileName)
			log.Print(err)
		}

		if file.Title.String != epub.title && epub.title != "" {
			//fmt.Println("Title:", epub.title)
			file.Title.String = epub.title
			update = true
		}
		if file.Authors.String != epub.authors && epub.authors != "" {
			//fmt.Println("Authors:", epub.authors)
			file.Authors.String = epub.authors
			update = true
		}
		if file.Subjects.String != epub.subjects && epub.subjects != "" {
			//fmt.Println("Subjects:", epub.subjects)
			file.Subjects.String = epub.subjects
			update = true
		}
		if file.Publisher.String != epub.publisher && epub.publisher != "" {
			//fmt.Println("Publisher:", epub.publisher)
			file.Publisher.String = epub.publisher
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
