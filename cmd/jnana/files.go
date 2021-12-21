package main

import (
	. "jnana/internal"
	"jnana/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"fmt"
	"github.com/campoy/unique"
	"github.com/djherbis/times"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// AllFiles get all files, used for update
func (db *Database) AllFiles() ([]*models.File, error) {
	var files []*models.File
	var err error

	files, err = models.Files(qm.SQL("SELECT * FROM file")).All(db.ctx, db.db)
	return files, err
}

// CoverForFile creates thumbnails using ImageMagick
// 160x160 - 80x80 2x
// 7.x: magick convert -resize 160x160 -background transparent -colorspace srgb -depth 8 -gravity center -extent 160x160 -strip  …
func (db *Database) CoverForFile(fileRecord *models.File, coversCacheDir string) bool {
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

func (db *Database) GetFile(book string, check bool) (*models.File, bool, error) {
	var file *models.File
	var err error
	var hash string
	changed := false // return value, to later recheck bookmarks

	// first lookup by file path
	file, err = db.GetFileFromPath(book)

	// not found, possible file moved, look up by file hash
	if err != nil {
		hash, err = FileHash(book)
		file, err = db.GetFileFromHash(hash)
	}

	// not found by path or hash, create new
	if err != nil {
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

		if err != nil {
			fmt.Println("date error:", err)
		}

		if modDate.After(file.DateModified) {
			//date different, check hash value
			changed = true
			file.Hash, _ = FileHash(book)
			file.DateModified = modDate
			err = db.UpdateFile(*file)
		}
	}
	return file, changed, err
}

// GetFileFromPath Look for existing record by file path
// return columns needed by GetFile, all in case of update
func (db *Database) GetFileFromPath(book string) (*models.File, error) {
	file, err := models.Files(qm.Where("path = ?", book)).One(db.ctx, db.db)
	return file, err
}

// GetFileFromHash look for existing by file hash (sha256)
// return columns needed by GetFile, all in case of update
func (db *Database) GetFileFromHash(hash string) (*models.File, error) {
	file, err := models.Files(qm.Where("hash = ?", hash)).One(db.ctx, db.db)
	return file, err
}

// NewFile create new file entry.
// models.File struct comes in with only path.
// Required fields: path, name, extension, created, modified, hash
func (db *Database) NewFile(book string) (*models.File, error) {
	stat, err := os.Stat(book)
	if err != nil {
		return &models.File{}, err
	}
	// format string for insert, strange set then get by format doesn't work
	dateModified := stat.ModTime().UTC()
	hash, _ := FileHash(book)

	f := File{}
	if err = f.Init(book); err != nil {
		return &models.File{}, err
	}

	t, err := times.Stat(book)
	if err != nil {
		log.Fatal(err.Error())
	}
	var dateCreated time.Time
	if t.HasBirthTime() {
		dateCreated = t.BirthTime().UTC()
	}

	tx, err := db.db.BeginTx(db.ctx, nil)
	if err != nil {
		return &models.File{}, err
	}
	defer tx.Rollback()

	// filepath.Ext returns with dot
	var newFile models.File

	newFile.Path = book
	newFile.Name = filepath.Base(book)
	newFile.Size = stat.Size()
	newFile.Extension = strings.ToLower(filepath.Ext(book)[1:])
	newFile.Title = null.StringFrom(f.title)
	newFile.Creator = null.StringFrom(f.creator)
	newFile.Subject = null.StringFrom(f.subject)
	newFile.Publisher = null.StringFrom(f.publisher)
	newFile.DateCreated = dateCreated
	newFile.DateModified = dateModified
	newFile.DateAccessed = null.TimeFrom(t.AccessTime().UTC())
	newFile.Hash = hash

	if err != nil {
		// TODO: workaround PDF metadata issue
		newFile.Path = book
		newFile.Name = filepath.Base(book)
		newFile.Size = stat.Size()
		newFile.Extension = strings.ToLower(filepath.Ext(book)[1:])
		newFile.DateCreated = dateCreated
		newFile.DateModified = dateModified
		newFile.DateAccessed = null.TimeFrom(t.AccessTime().UTC())
		newFile.Hash = hash
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
func (db *Database) RecentFiles() ([]*models.File, error) {
	files, err := models.Files(qm.SQL("SELECT * FROM file ORDER BY date_accessed DESC LIMIT 50")).All(db.ctx, db.db)
	return files, err
}

// SearchFiles Search all files from FTS5 table,
// order by rank: name, title, creator, subject, publisher, description
// Return results as slice of struct models.File, later prepped for Alfred script filter
func (db *Database) SearchFiles(query string) ([]*models.File, error) {
	queryString := stringForSQLite(query)
	var results []*models.File

	// NOTE: AND rank MATCH 'bm25(…)' ORDER BY rank faster than ORDER BY bm25(fts, …)
	err := queries.Raw(`SELECT
			file.id, file.path, file.name, file.extension, file.title, file.subject
			FROM file
			JOIN file_search on file.id = file_search.rowid
			WHERE file_search MATCH ?
			AND rank MATCH 'bm25(10.0, 2.0, 2.0, 2.0, 2.0, 2.0)'
			ORDER BY rank LIMIT 200`,
		queryString).Bind(db.ctx, db.db, &results)

	err = db.db.Close()
	return results, err
}

// UpdateFile update file on change of path, file name, or date modified
func (db *Database) UpdateFile(file models.File) error {
	stat, err := os.Stat(file.Path)
	if err != nil {
		return err
	}
	t, err := times.Stat(file.Path)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.DateAccessed = null.TimeFrom(t.AccessTime().UTC())

	tx, err := db.db.BeginTx(db.ctx, nil)

	_, err = queries.Raw(`UPDATE file SET
			path = $1, name = $2, size = $3,
			title = $4, creator = $5, subject = $6, publisher = $7,
			date_modified = $8, date_accessed = $9, hash = $10
			WHERE id = $11`,
		file.Path, filepath.Base(file.Path), stat.Size(),
		file.Title.String, file.Creator.String, file.Subject.String, file.Publisher.String,
		file.ID).Exec(db.db)

	if err != nil {
		// TODO: PDF metadata "unrecognized token", workaround don't update metadata
		_, err = queries.Raw(`UPDATE file SET
			path = ?, name = ?, size = ?, date_modified = ?, date_accessed = ?, hash = ?
			WHERE id = ?`,
			file.Path, filepath.Base(file.Path), stat.Size(), file.DateModified, file.DateAccessed,
			file.Hash, file.ID).Exec(db.db)
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error committing:", err)
	}
	return err
}

// UpdateDateAccessed update last opened
func (db *Database) UpdateDateAccessed(file *models.File) {
	currentTime := time.Now().UTC()
	_, _ = queries.Raw(`UPDATE file SET date_accessed = ? WHERE id = ?`,
		currentTime, file.ID).Exec(db.db)
}

// UpdateMetadata check for updates to metadata
func (db *Database) UpdateMetadata(file *models.File) (bool, error) {
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
func (db *Database) UpdateSubject(file *models.File, subject string) error {
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
