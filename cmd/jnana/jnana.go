package main

import (
	"database/sql"
	bookFile "jnana/internal/file"
	"jnana/internal/files"
	"jnana/internal/util"

	"jnana/internal/bookmarks"
	"jnana/internal/database"
	"jnana/models"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/deanishe/awgo"
	"github.com/docopt/docopt-go"
)

var (
	repo = "gennaios/alfred-jnana"

	usage = `jnana [command] [<query>...]

usage:
    jnana all [<query>]
    jnana bm <file>
    jnana bmf <file> <query>
    jnana clean
    jnana epub [<query>]
    jnana filecreators <file> [<query>]
    jnana filetitle <file> [<query>]
    jnana files [<query>]
    jnana import <file>
    jnana isbn <file> [<query>]
    jnana getepub
    jnana openepub <query> [<file>]
    jnana openfile <file>
    jnana pdf <file> [<query>]
    jnana recent
    jnana test <file>
    jnana lastquery
    jnana lastfilequery
    jnana savequery <query>
    jnana savefilequery <query>
    jnana subject <file> [<query>]
    jnana test <file>
    jnana update [<file>]
    jnana updateread <file>
    jnana -h

options:
    -h --help          Show this message and exit.
    --version          Show workflow version and exit.

commands:
    all				Search all bookmarks.
    bm				Bookmarks for file
    bmf				Bookmarks for file filtered by query
    clean			Clean database, remove bookmarks for deleted files
    epub		   	Bookmarks for EPUB in calibre
	filecreators	Get or set file creators
	filetitle		Get or set file title
    files			Search all files by name and metadata
    getepub     	Return opened EPUB
    import      	Import file or files from folder
	isbn 	     	Set ISBN for EPUB
    openepub		open calibre to bookmark
	openfile		open file
    pdf				Retrieve or filter bookmarks for opened PDF in Acrobat, Preview, or Skim
	recent			Show recently opened files
    lastquery		Retrieve cached last query string for script filter
    lastfilequery	Retrieve cached last file query string for script filter
    savefilequery	Save last file query string for script filter
    savequery		Save last query string for script filter
    test        	Testing stuff
    update      	Update path metadata
	updateread		Update date read
`

	wf *aw.Workflow

	dbFileName     = "jnana.db"
	coversCacheDir string // directory generated icons are stored in
)

var (
	options struct {
		// commands
		All           bool
		Bm            bool
		Bmf           bool
		Clean         bool
		Epub          bool
		Filecreators  bool
		Filetitle     bool
		Files         bool
		Getepub       bool
		Import        bool
		Isbn          bool
		Openepub      bool
		Openfile      bool
		Pdf           bool
		Lastquery     bool
		Lastfilequery bool
		Recent        bool
		Savefilequery bool
		Savequery     bool
		Subject       bool
		Test          bool
		Update        bool
		Updateread    bool

		// parameters
		File   string
		Fileid int64
		Query  string
	}
)

func init() {
	// Create a new Workflow using default settings.
	wf = aw.New(aw.HelpURL(repo + "/issues"))

	coversCacheDir = filepath.Join(wf.DataDir(), "covers")
}

// initDatabase: initialize SQLite database
func initDatabase() *database.Database {
	dbFile := filepath.Join(wf.DataDir(), dbFileName)
	db := &database.Database{}
	db.Init(dbFile)
	return db
}

// initDatabase: initialize SQLite database
func initDatabaseForReading() *database.Database {
	dbFile := filepath.Join(wf.DataDir(), dbFileName)
	db := &database.Database{}
	db.InitForReading(dbFile)
	return db
}

// Bookmarks all for file, from database or imported, return results
func bookmarksForFile(file string) {
	db := initDatabase()

	bookmarksRecord, err := bookmarks.ForFile(db, file, coversCacheDir)

	if err == nil {
		returnBookmarksForFile(file, bookmarksRecord)
	} else {
		wf.FatalError(err)
	}
}

func bookmarksForFileEpub(query string) {
	epub := calibreEpubFile()
	if query != "" {
		// TODO: should already know file if filtered
		bookmarksForFileFiltered(epub, query)
	} else {
		bookmarksForFile(epub)
	}
}

// Bookmarks filtered for file, from database or imported, return results
func bookmarksForFileFiltered(file string, query string) {
	db := initDatabaseForReading()

	bookmarksRecord, err := bookmarks.ForFileFiltered(db, file, query)

	if err == nil {
		returnBookmarksForFileFiltered(file, bookmarksRecord)
	} else {
		wf.FatalError(err)
	}
}

func calibreEpubFile() string {
	usr, _ := user.Current()
	var path string

	// find calibre preferences

	// pre 4.x
	//calibreJsonFile := "~/Library/Preferences/calibre/viewer.json"

	// 4.x
	calibreJsonFile := "~/Library/Preferences/calibre/viewer-webengine.json"
	path = filepath.Join(usr.HomeDir, calibreJsonFile[2:])

	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		wf.FatalError(err)
	}

	// find most recently opened file = current

	// pre 4.x
	//return jsonData["viewer_open_history"][0]

	// 4.x
	var result map[string]interface{}
	if err := json.Unmarshal(jsonFile, &result); err != nil {
		log.Fatal(err)
	}

	fileName := result["session_data"].(map[string]interface{})["standalone_recently_opened"].([]interface{})[0].(map[string]interface{})["pathtoebook"]

	return fileName.(string)
}

func cleanDatabase() {
	db := initDatabase()

	all, err := files.All(db)
	if err != nil {
		log.Println(err)
	}

	log.Println("Looking for deleted files out of %d" + strconv.Itoa(len(all)))

	var file string
	for i := range all {
		file = all[i].Path
		if _, err := os.Stat(file); err != nil {
			log.Println("File:", file)
		}
	}
}

// iconForFileID: retrieve cover image from covers folder, or return generic icon
func iconForFileID(fileId string, filePath string) *aw.Icon {
	iconFile := filepath.Join(coversCacheDir, fileId+".png")

	if _, err := os.Stat(iconFile); err == nil {
		return &aw.Icon{
			Value: iconFile,
			Type:  aw.IconTypeImage,
		}
	} else {
		return &aw.Icon{
			Value: filePath,
			Type:  aw.IconTypeFileIcon,
		}
	}
}

func cacheLastFileQuery(queryString string) {
	if err := wf.Cache.StoreJSON("LastFileQuery", queryString); err != nil {
		wf.FatalError(err)
	}
}

func cacheLastQuery(queryString string) {
	if err := wf.Cache.StoreJSON("LastQuery", queryString); err != nil {
		wf.FatalError(err)
	}
}

func fileCreators(file string, creators string) {
	db := initDatabase()
	fileRecord, _, err := files.Get(db, file, false)
	if err != nil {
		wf.FatalError(err)
	}

	// return current from File for editing
	if creators == "" {
		oldCreators := fileRecord.Creator.String
		oldCreators = files.GetCreators(oldCreators)
		fmt.Println(oldCreators)
		return
	}

	// set new creators if specified
	err = files.SetCreators(db, fileRecord, creators)
	if err != nil {
		wf.FatalError(err)
	}
	_ = db.Db.Close()
	fmt.Println(creators)
}

func fileISBN(file string, isbn string) {
	db := initDatabase()
	fileRecord, _, err := files.Get(db, file, false)
	if err != nil {
		wf.FatalError(err)
	}

	// set new ISBN if specified
	err = files.UpdateISBN(fileRecord, isbn)
	if err != nil {
		wf.FatalError(err)
	}
	_ = db.Db.Close()
	fmt.Println(isbn)
}

func fileSubject(file string, subject string) {
	db := initDatabase()

	if subject == "" {
		fileRecord, _, err := files.Get(db, file, false)
		if err != nil {
			wf.FatalError(err)
		}

		fmt.Println(fileRecord.Subject.String)
	} else {
		fileRecord, _, err := files.Get(db, file, false)

		err = files.UpdateSubject(db, fileRecord, subject)
		if err != nil {
			wf.FatalError(err)
		}
		_ = db.Db.Close()
	}
}

func fileTitle(file string, newTitle string) {
	db := initDatabase()
	fileRecord, _, err := files.Get(db, file, false)
	if err != nil {
		wf.FatalError(err)
	}

	// get title
	if newTitle == "" {
		fmt.Println(fileRecord.Title.String)
	}

	// set new title if specified
	err = files.UpdateTitle(db, fileRecord, newTitle)
	if err != nil {
		wf.FatalError(err)
	}
	_ = db.Db.Close()
	fmt.Println(newTitle)
}

func getCurrentEpub() {
	fmt.Println(calibreEpubFile())
}

func getLastFileQuery() string {
	var lastQuery string
	if err := wf.Cache.LoadJSON("LastFileQuery", &lastQuery); err != nil {
		wf.FatalError(err)
	}
	return lastQuery
}

func getLastQuery() string {
	var lastQuery string
	if err := wf.Cache.LoadJSON("LastQuery", &lastQuery); err != nil {
		wf.FatalError(err)
	}
	return lastQuery
}

// ImportFile import file or all files in folder
func ImportFile(db *database.Database, file string) error {
	var err error

	if strings.HasSuffix(file, ".epub") || strings.HasSuffix(file, ".pdf") {
		_, err := files.GetFromPath(db, file)

		if err == sql.ErrNoRows {
			fileRecord, changed, err := files.Get(db, file, false)
			if err != nil {
				return err
			}

			if changed == true {
				bookmarksRecord, err := bookmarks.ForFile(db, file, coversCacheDir)
				if err != nil {
					return err
				}
				if len(bookmarksRecord) != 0 {
					log.Println("Imported:", fileRecord.Name)
				}
			}
		}
	}
	return err
}

// ImportFiles import file or all files in folder
func ImportFiles(file string) {
	db := initDatabase()

	fi, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		_ = filepath.Walk(file, func(path string, f os.FileInfo, err error) error {
			//ImportFile(db, path)
			aFile, _ := filepath.Abs(path)
			if err := ImportFile(db, aFile); err != nil {
				log.Println("ERROR:", aFile+" / "+err.Error())
			}
			return nil
		})
	case mode.IsRegular():
		aFile, _ := filepath.Abs(file)
		if err := ImportFile(db, aFile); err != nil {
			log.Println("ERROR:", aFile+" / "+err.Error())
		}
	}
}

// receive bookmark title as query from script filter and open calibre
func openCalibreBookmark(destination string, file string) {
	command := "/Applications/calibre.app/Contents/MacOS/ebook-viewer"
	if file == "" {
		file = calibreEpubFile()
	}
	file = "\"" + file + "\"" // for shell script

	// --open-at=toc: "first Table of Contents entry that contains"
	// --open-at=toc-href: "href (internal link destination) of toc nodes"
	// --open-at=toc-href-contains: href substring
	cmdArgs := []string{"--open-at=toc-href-contains:\"" + destination + "\"", file}

	openCalibreBookmarkCommand(command, cmdArgs)
	//_ = exec.Command(command, cmdArgs...).Start()
}

// open calibre to bookmark
// workaround command exec issues by creating shell script
func openCalibreBookmarkCommand(command string, cmdArgs []string) {
	//os.RemoveAll(output_path)
	temp := "/tmp"
	file, _ := os.Create(filepath.Join(temp, "alfred-jnana.sh"))
	defer func() {
		_ = file.Close()
	}()

	_, _ = file.WriteString("#!/bin/sh\n")
	_, _ = file.WriteString(command + " " + strings.Join(cmdArgs, " "))
	_, _ = file.WriteString("\n")
	_ = os.Chdir(temp)
	_ = exec.Command("sh", "alfred-jnana.sh").Start()
}

// openFile opens EPUB or PDF from catalog
func openFile(file string) {
	if exists, _ := util.Exists(file); exists == false {
		_ = util.Notification("Does not exist: " + file)
		return
	}

	// open with reader if EPUB, else handle in Workflow per active PDF app
	if strings.HasSuffix(file, ".epub") {
		command := "open"
		cmdArgs := []string{file}
		_ = exec.Command(command, cmdArgs...).Start()
	}
}

func printLastQuery() {
	fmt.Println(getLastQuery())
}

func printLastFileQuery() {
	fmt.Println(getLastFileQuery())
}

// RecentFiles List recently opened files
func RecentFiles() {
	db := initDatabaseForReading()

	results, err := files.Recent(db)
	if err != nil {
		wf.FatalError(err)
	}
	returnSearchFilesResults(results, "")
}

// Query database for all bookmarks
func searchAllBookmarks(query string) {
	db := initDatabaseForReading()

	results, err := bookmarks.SearchAll(db, query)
	if err != nil {
		wf.FatalError(err)
	}
	returnSearchAllResults(results, query)
}

// Query database for all files
func searchAllFiles(query string) {
	db := initDatabaseForReading()

	results, err := files.Search(db, query)
	if err != nil {
		wf.FatalError(err)
	}
	returnSearchFilesResults(results, query)
}

// Bookmarks for single EPUB or PDF
// - input: bookmarks
// - output to Alfred
func returnBookmarksForFile(file string, bookmarks []*models.Bookmark) {
	var icon *aw.Icon

	if strings.HasSuffix(file, "pdf") {
		var destination string
		var subtitle string
		icon = &aw.Icon{Value: "com.adobe.pdf", Type: aw.IconTypeFileType}

		for i := range bookmarks {
			destination = bookmarks[i].Destination

			if bookmarks[i].Section.String != "" {
				subtitle = "Page " + bookmarks[i].Destination + ". " + bookmarks[i].Section.String
			} else {
				subtitle = "Page " + bookmarks[i].Destination
			}

			wf.NewItem(bookmarks[i].Title.String).
				Subtitle(subtitle).
				UID(strconv.FormatInt(bookmarks[i].ID, 10)).
				Valid(true).
				Icon(icon).
				Arg(destination)
		}
	} else {
		icon = &aw.Icon{Value: "org.idpf.epub-container", Type: aw.IconTypeFileType}

		// Title.String = TOC entry name
		// Destination = TOC entry destination
		for i := range bookmarks {
			wf.NewItem(bookmarks[i].Title.String).
				Subtitle(bookmarks[i].Section.String).
				UID(strconv.FormatInt(bookmarks[i].ID, 10)).
				Valid(true).
				Icon(icon).
				Arg(bookmarks[i].Destination)
		}

	}
	wf.SendFeedback()
}

func returnBookmarksForFileFiltered(file string, bookmarks []*bookmarks.SearchAllResult) {
	var icon *aw.Icon

	if strings.HasSuffix(file, "pdf") {
		var subtitle string
		icon = &aw.Icon{Value: "com.adobe.pdf", Type: aw.IconTypeFileType}

		for i := range bookmarks {
			if bookmarks[i].Section.String != "" {
				subtitle = "Page " + bookmarks[i].Destination + ". " + bookmarks[i].Section.String
			} else {
				subtitle = "Page " + bookmarks[i].Destination
			}

			wf.NewItem(bookmarks[i].Title.String).
				Subtitle(subtitle).
				UID(strconv.FormatInt(bookmarks[i].ID, 10)).
				Valid(true).
				Icon(icon).
				Arg(bookmarks[i].Destination)
		}
	} else {
		icon = &aw.Icon{Value: "org.idpf.epub-container", Type: aw.IconTypeFileType}

		// Title.String = TOC entry name
		// Destination = TOC entry destination
		for i := range bookmarks {
			wf.NewItem(bookmarks[i].Title.String).
				Subtitle(bookmarks[i].Section.String).
				UID(strconv.FormatInt(bookmarks[i].ID, 10)).
				Valid(true).
				Icon(icon).
				Arg(bookmarks[i].Destination)
		}
	}
	wf.SendFeedback()
}

// Parse database search results and return items to Alfred
func returnSearchAllResults(bookmarks []*bookmarks.SearchAllResult, query string) {
	var title string
	var subtitle string
	var arg string

	wf.Var("JNANA_QUERY", query)

	for i := range bookmarks {
		icon := iconForFileID(bookmarks[i].FileID, bookmarks[i].Path)

		if bookmarks[i].Section.String != "" {
			title = bookmarks[i].Title.String + " | " + bookmarks[i].Section.String
		} else {
			title = bookmarks[i].Title.String
		}

		if strings.HasSuffix(bookmarks[i].Name, ".pdf") {
			subtitle = "Page " + bookmarks[i].Destination + ". " + bookmarks[i].Name
			arg = bookmarks[i].Path + "/Page:" + bookmarks[i].Destination
		} else {
			subtitle = bookmarks[i].Name
			arg = bookmarks[i].Path + "/Page:" + bookmarks[i].Destination
		}

		wf.NewItem(title).
			Subtitle(subtitle).
			UID(strconv.FormatInt(bookmarks[i].ID, 10)).
			Valid(true).
			Icon(icon).
			Arg(arg)
	}
	wf.SendFeedback()
}

// Parse database search results and return items to Alfred
func returnSearchFilesResults(files []*models.File, query string) {
	var title string

	// return query variable for next search
	if query != "" {
		wf.Var("JNANA_FILE_QUERY", query)
	}

	for i := range files {
		icon := iconForFileID(strconv.FormatInt(files[i].ID, 10), files[i].Path)

		if files[i].Format == 2 {
			title = files[i].Name
		} else {
			if files[i].Creator.String != "" {
				title = files[i].Title.String + " - " + files[i].Creator.String
			} else {
				title = files[i].Title.String
			}
		}

		wf.NewItem(title).
			Subtitle(files[i].Subject.String).
			UID(strconv.FormatInt(files[i].ID, 10)).
			Valid(true).
			Icon(icon).
			Arg(files[i].Path)
	}
	wf.SendFeedback()
}

func TestStuff(file string) {
	//bookmarks, _ := bookmarksForPDF(path)
	//fmt.Println("bookmarks", len(bookmarks))
	f := bookFile.File{}
	if err := f.Init(file); err != nil {
		log.Println(err)
	}
	_ = f.Init(file)
	bookmarks2, _ := f.Bookmarks()
	fmt.Println("Bookmarks", len(bookmarks2))
}

// UpdateFile check one file for metadata updates, not including bookmarks
func UpdateFile(db *database.Database, fileRecord *models.File) {
	updated, err := files.UpdateMetadata(db, fileRecord)
	if err != nil {
		log.Println("Error:", fileRecord.Path)
	}
	if updated == true {
		fmt.Println("Updated:", fileRecord.Path)
	}
}

// UpdateFiles check passed file or all files for metadata changes, not including bookmarks
func UpdateFiles(file string) {
	db := initDatabase()

	if _, err := os.Stat(file); err == nil {
		fileRecord, _, _ := files.Get(db, file, false)
		UpdateFile(db, fileRecord)
	} else {
		filesRecord, err := files.All(db)
		if err != nil {
			log.Fatal(err)
		}
		for _, aFile := range filesRecord {
			if util.FileExists(aFile.Path) {
				UpdateFile(db, aFile)
			}
		}
	}
}

// UpdateRead set date read (date_accessed)
func UpdateRead(file string) {
	if exists, _ := util.Exists(file); exists == false {
		return
	}

	db := initDatabase()
	fileRecord, _, _ := files.Get(db, file, false)
	files.UpdateDateAccessed(db, fileRecord)
}

func runCommand() {
	// show options for debug
	//fmt.Println(options)

	// normalize white space, remove dupes
	query := strings.Join(strings.Fields(strings.TrimSpace(options.Query)), " ")

	switch true {
	case options.All:
		searchAllBookmarks(query)
	case options.Bm:
		bookmarksForFile(options.File)
	case options.Bmf:
		bookmarksForFileFiltered(options.File, query)
	case options.Clean:
		cleanDatabase()
	case options.Epub:
		bookmarksForFileEpub(query)
	case options.Subject:
		fileSubject(options.File, query)
	case options.Filecreators:
		fileCreators(options.File, query)
	case options.Filetitle:
		fileTitle(options.File, query)
	case options.Files:
		searchAllFiles(query)
	case options.Import:
		ImportFiles(options.File)
	case options.Isbn:
		fileISBN(options.File, query)
	case options.Getepub:
		getCurrentEpub()
	case options.Openepub:
		openCalibreBookmark(query, options.File)
	case options.Openfile:
		openFile(options.File)
	case options.Lastquery:
		printLastQuery()
	case options.Savequery:
		cacheLastQuery(query)
	case options.Lastfilequery:
		printLastFileQuery()
	case options.Savefilequery:
		cacheLastFileQuery(query)
	case options.Recent:
		RecentFiles()
	case options.Test:
		TestStuff(options.File)
	case options.Update:
		UpdateFiles(options.File)
	case options.Updateread:
		UpdateRead(options.File)
	}
}

// workflow start
func run() {
	parser := &docopt.Parser{
		HelpHandler:  docopt.PrintHelpOnly,
		OptionsFirst: true,
	}
	args, _ := parser.ParseArgs(usage, nil, wf.Version())
	err := args.Bind(&options)
	if err != nil {
		wf.FatalError(err)
	}

	runCommand()
}

func main() {
	// calls run()
	wf.Run(run)
}
