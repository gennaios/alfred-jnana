package files

import (
	"database/sql"
	"jnana/internal/database"
	bookFile "jnana/internal/file"
	"jnana/internal/fulltext"
	"jnana/internal/util"
	"jnana/models"
	"os/exec"
	"regexp"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"fmt"
	"github.com/campoy/unique"
	"github.com/djherbis/times"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// All get all files, used for update
func All(db *database.Database) ([]*models.File, error) {
	var files []*models.File
	var err error

	files, err = models.Files(qm.SQL("SELECT * FROM file")).All(db.Ctx, db.Db)
	return files, err
}

// CoverForFile creates thumbnails using ImageMagick
// 160x160 - 80x80 2x
// 7.x: magick convert -resize 160x160 -background transparent -colorspace srgb -depth 8 -gravity center -extent 160x160 -strip  …
func CoverForFile(fileRecord *models.File, coversCacheDir string) bool {
	var err error
	coverPath := filepath.Join(coversCacheDir, strconv.FormatInt(fileRecord.ID, 10)+".png")

	_, err = os.Stat(coverPath)
	if !os.IsNotExist(err) {
		return true
	} else {
		_ = util.Notification("File ID: " + strconv.FormatInt(fileRecord.ID, 10) + ".png")
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

func Get(db *database.Database, book string, check bool) (*models.File, bool, error) {
	var file *models.File
	var err error
	var hash string
	changed := false // return value, to later recheck bookmarks

	// first lookup by file path
	file, err = GetFromPath(db, book)

	// not found, possible file moved, look up by file hash
	if err == sql.ErrNoRows {
		hash, err = util.FileHash(book)
		file, err = GetFromHash(db, hash)
	}

	// not found by path or hash, create new
	if err == sql.ErrNoRows {
		file, err = New(db, book)
		if err != nil {
			return file, true, err
		}
	} else if file != nil {
		if file.ID != 0 && book != file.Path {
			// found by hash, verify not dupe
			if _, err = os.Stat(file.Path); err != nil {
				if os.IsNotExist(err) {
					// old path doesn't exist, moved, same hash
					_ = util.Notification("File moved: " + file.Path)
					file.Path = book
					err = Update(db, *file)
					// hash match, no changes needed
					return file, false, err
				}
			} else {
				// check file exists, or notification will be triggered if not
				if _, err = os.Stat(book); err != nil {
					if !os.IsNotExist(err) {
						_ = util.Notification("Dupe of: " + file.Path)
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
			file.Hash, _ = util.FileHash(book)
			file.DateModified = modDate
			err = Update(db, *file)
		}
	}
	return file, changed, err
}

// GetFromPath Look for existing record by file path
// return columns needed by Get, all in case of update
func GetFromPath(db *database.Database, book string) (*models.File, error) {
	file, err := models.Files(qm.Where("path = ?", book)).One(db.Ctx, db.Db)
	return file, err
}

// GetFromHash look for existing by file hash (sha256)
// return columns needed by Get, all in case of update
func GetFromHash(db *database.Database, hash string) (*models.File, error) {
	file, err := models.Files(qm.Where("hash = ?", hash)).One(db.Ctx, db.Db)
	return file, err
}

// New create new file entry.
// models.File struct comes in with only path.
// Required fields: path, name, extension, created, modified, hash
func New(db *database.Database, book string) (*models.File, error) {
	stat, err := os.Stat(book)
	if err != nil {
		return &models.File{}, err
	}
	// format string for insert, strange set then get by format doesn't work
	dateModified := stat.ModTime().UTC().Truncate(time.Millisecond)
	hash, _ := util.FileHash(book)

	f := bookFile.File{}
	if err = f.Init(book); err != nil {
		return &models.File{}, err
	}

	t, err := times.Stat(book)
	if err != nil {
		log.Fatal(err.Error())
	}
	var dateCreated time.Time
	if t.HasBirthTime() {
		dateCreated = t.BirthTime().UTC().Truncate(time.Millisecond)
	}

	tx, err := db.Db.BeginTx(db.Ctx, nil)
	if err != nil {
		return &models.File{}, err
	}
	defer tx.Rollback()

	_, _ = queries.Raw("INSERT INTO file (path, name, extension, size, title, publisher, creator, subject, date_created, date_modified, date_accessed, hash) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		book,
		filepath.Base(book),
		strings.ToLower(filepath.Ext(book)[1:]),
		stat.Size(),
		null.StringFrom(f.Title),
		null.StringFrom(f.Publisher),
		null.StringFrom(f.Creator),
		null.StringFrom(f.Subject),
		dateCreated,
		dateModified,
		null.TimeFrom(t.AccessTime().UTC().Truncate(time.Millisecond)),
		hash).Exec(db.Db)

	err = tx.Commit()

	//err = f.file.Close() // TODO: invalid memory address or nil pointer, maybe closed in File?
	if strings.HasSuffix(book, ".epub") {
		f.Epub.Close()
	}

	file, err := GetFromPath(db, book)
	return file, err
}

// Recent list of recently opened files
func Recent(db *database.Database) ([]*models.File, error) {
	files, err := models.Files(qm.SQL("SELECT * FROM file ORDER BY date_accessed DESC LIMIT 50")).All(db.Ctx, db.Db)
	return files, err
}

// Search all files from FTS5 table,
// order by rank: name, title, creator, subject, publisher, description
// Return results as slice of struct models.File, later prepped for Alfred script filter
func Search(db *database.Database, query string) ([]*models.File, error) {
	queryString := util.StringForSQLite(query)
	var results []*models.File

	// NOTE: AND rank MATCH 'bm25(…)' ORDER BY rank faster than ORDER BY bm25(fts, …)
	err := queries.Raw(`SELECT
			file.id, file.path, file.name, file.extension, file.title, file.subject
			FROM file
			JOIN file_search on file.id = file_search.rowid
			WHERE file_search MATCH ?
			AND rank MATCH 'bm25(10.0, 2.0, 2.0, 2.0, 2.0, 2.0)'
			ORDER BY rank LIMIT 200`,
		queryString).Bind(db.Ctx, db.Db, &results)

	err = db.Db.Close()
	return results, err
}

// Update on change of path, file name, or date modified
func Update(db *database.Database, file models.File) error {
	stat, err := os.Stat(file.Path)
	if err != nil {
		return err
	}
	t, err := times.Stat(file.Path)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.DateAccessed = null.TimeFrom(t.AccessTime().UTC())

	tx, err := db.Db.BeginTx(db.Ctx, nil)

	// delete old FTS
	oldFile, _ := models.Files(qm.Where("id = ?", file.ID)).One(db.Ctx, db.Db)
	if compare(*oldFile, file) != true {
		fulltext.FileDelete(db, file)
		fulltext.BookmarksDelete(db, file)
	}

	// update record
	_, err = queries.Raw(`UPDATE file SET
			path = $1, name = $2, size = $3,
			title = $4, creator = $5, subject = $6, publisher = $7,
			date_modified = $8, date_accessed = $9, hash = $10
			WHERE id = $11`,
		file.Path, filepath.Base(file.Path), stat.Size(),
		file.Title.String, file.Creator.String, file.Subject.String, file.Publisher.String,
		file.DateModified, file.DateAccessed, file.Hash,
		file.ID).Exec(db.Db)

	// create new FTS
	if compare(*oldFile, file) != true {
		fulltext.FileCreate(db, file)
		fulltext.BookmarksCreate(db, file)
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error committing:", err)
	}
	return err
}

// UpdateDateAccessed update last opened
func UpdateDateAccessed(db *database.Database, file *models.File) {
	currentTime := time.Now().UTC()
	_, _ = queries.Raw(`UPDATE file SET date_accessed = ? WHERE id = ?`,
		currentTime, file.ID).Exec(db.Db)
}

// UpdateMetadata check for updates to metadata
func UpdateMetadata(db *database.Database, file *models.File) (bool, error) {
	var err error
	if _, err = os.Stat(file.Path); err != nil {
		return false, err
	}
	update := false

	f := bookFile.File{}
	if err = f.Init(file.Path); err != nil {
		return false, err
	}

	if file.Title.String != f.Title && f.Title != "" {
		file.Title.String = f.Title
		update = true
	}
	if file.Creator.String != f.Creator && f.Creator != "" {
		file.Creator.String = f.Creator
		update = true
	}
	if strings.HasSuffix(file.Path, "epub") {
		if file.Publisher.String != f.Publisher && f.Publisher != "" {
			file.Publisher.String = f.Publisher
			update = true
		}
	}

	if update == true {
		err = Update(db, *file)
		if err == nil {
			return true, err
		}
	}

	return false, err
}

func ParseCreators(newCreators string) string {
	// fix creators
	creatorsArray := util.StringSplitAny(newCreators, ";&")
	for i := range creatorsArray {
		creatorsArray[i] = strings.Replace(creatorsArray[i], ", Jr", " Jr", -1)
		creatorsArray[i] = strings.Replace(creatorsArray[i], " (EDT)", "", -1)

		if strings.Contains(creatorsArray[i], ",") {
			newCreator := strings.Split(creatorsArray[i], ",")
			newCreator[0], newCreator[1] = newCreator[1], newCreator[0]
			creatorsArray[i] = strings.Join(newCreator[:], " ")
		}

		// remove double spaces
		creatorsArray[i] = strings.TrimSpace(creatorsArray[i])
		space := regexp.MustCompile(`\s+`)
		creatorsArray[i] = space.ReplaceAllString(creatorsArray[i], " ")
	}
	return strings.Join(creatorsArray[:], ", ")
}

// UpdateCreators set creators for file
func UpdateCreators(db *database.Database, file *models.File, newCreators string) error {
	var err error

	if newCreators == file.Creator.String || newCreators == "" {
		return err
	}

	// set creators
	calibreMeta := "/Applications/calibre.app/Contents/MacOS/ebook-meta"
	if file.Extension == "epub" && util.FileExists(calibreMeta) {
		if err = exec.Command(calibreMeta, file.Path, "-a", strings.Replace(newCreators, ", ", " & ", -1)).Run(); err == nil {
			file.Creator = null.StringFrom(newCreators)
			err = Update(db, *file)
			if err == nil {
				util.Notification("Creators updated: " + file.Creator.String)
			}
		}
	}

	return err
}

// UpdateSubject set subject/keywords for file
func UpdateSubject(db *database.Database, file *models.File, subject string) error {
	var err error

	if subject == "" {
		return err
	}

	terms := strings.Split(strings.ToLower(subject), ",")
	s := bookFile.TrimMetadata(terms)

	less := util.LessString(&s)
	unique.Slice(&s, less)

	newSubject := strings.Join(s, ", ")

	if newSubject != file.Subject.String {
		file.Subject = null.StringFrom(newSubject)
		err = Update(db, *file)
	}

	return err
}

// UpdateTitle set title for file
func UpdateTitle(db *database.Database, file *models.File, newTitle string) error {
	var err error

	if newTitle == file.Title.String || newTitle == "" {
		return err
	}

	// set title
	calibreMeta := "/Applications/calibre.app/Contents/MacOS/ebook-meta"
	if file.Extension == "epub" && util.FileExists(calibreMeta) {
		// get title: " | head -n 1 | cut -d ':' -f2- | awk '{$1=$1};1'"
		if err = exec.Command(calibreMeta, file.Path, "-t", strings.TrimSpace(newTitle)).Run(); err == nil {
			file.Title = null.StringFrom(newTitle)
			err = Update(db, *file)
			if err == nil {
				util.Notification("Title updated: " + file.Title.String)
			}
		}
	}

	return err
}

func compare(a, b models.File) bool {
	if a == b {
		return true
	}

	return cmp.Equal(a, b, cmpopts.IgnoreFields(models.File{}, "Path", "Extension", "Size", "Language", "DateCreated",
		"DateModified", "DateAccessed", "Hash"))
}
