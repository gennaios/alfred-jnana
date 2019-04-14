package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocraft/dbr"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/docopt/docopt-go"
)

var (
	repo = "gennaios/alfred-jnana"

	usage = `jnana [command] [<query>...]

usage:
    jnana all [<query>]
    jnana allepub [<query>]
    jnana allpdf [<query>]
    jnana bm <file>
    jnana bmf <file> <query>
    jnana epub [<query>]
    jnana import <file>
    jnana getepub
    jnana openepub <query> [<file>]
    jnana pdf <file> [<query>]
    jnana test <file>
    jnana lastquery
    jnana test <file>
    jnana update [<file>]
    jnana -h

options:
    -h --help          Show this message and exit.
    --version          Show workflow version and exit.

commands:
    all		Search all bookmarks.
    allepub	Search all EPUB bookmarks.
    allpdf	Search all PDF bookmarks.
    bm		Bookmarks for file
    bmf		Bookmarks for file filtered by query
    epub	Bookmarks for EPUB in calibre
    epubf	Bookmarks for EPUB in calibre filtered by query
    getepub     Return opened EPUB
    import      Import file or files from folder	
    openepub	open calibre to bookmark
    pdf		Retrieve or filter bookmarks for opened PDF in Acrobat, Preview, or Skim.
    lastquery	Retrieve cached last query string for script filter
    test        Testing stuff
    update      Update path metadata
`

	wf *aw.Workflow

	dbFileName     = "jnana.db"
	coversCacheDir string // directory generated icons are stored in
)

var options struct {
	// commands
	All       bool
	Allepub   bool
	Allpdf    bool
	Bm        bool
	Bmf       bool
	Epub      bool
	Getepub   bool
	Import    bool
	Openepub  bool
	Pdf       bool
	Lastquery bool
	Test      bool
	Update    bool

	// parameters
	Query  string
	File   string
	Fileid string
}

func init() {
	// Create a new Workflow using default settings.
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))

	coversCacheDir = filepath.Join(wf.DataDir(), "covers")
}

// initDatabase: initialize SQLite database
func initDatabase(dbFile string) Database {
	db := Database{}
	db.Init(dbFile)
	return db
}

// initDatabase: initialize SQLite database
func initDatabaseForReading(dbFile string) Database {
	db := Database{}
	db.InitForReading(dbFile)
	return db
}

// Bookmarks all for file, from database or imported, return results
func bookmarksForFile(file string) {
	dbFile := filepath.Join(wf.DataDir(), dbFileName)
	db := initDatabase(dbFile)

	bookmarks, err := db.BookmarksForFile(file)
	if err == nil {
		returnBookmarksForFile(file, bookmarks)
	} else {
		wf.FatalError(err)
	}
}

func bookmarksForFileEpub(query string) {
	epub := calibreEpubFile()
	if query != "" {
		bookmarksForFileFiltered(epub, query)
	} else {
		bookmarksForFile(epub)
	}
}

// Bookmarks filtered for file, from database or imported, return results
func bookmarksForFileFiltered(file string, query string) {
	dbFile := filepath.Join(wf.DataDir(), dbFileName)
	db := initDatabaseForReading(dbFile)

	bookmarks, err := db.BookmarksForFileFiltered(file, query)

	if err == nil {
		returnBookmarksForFileFiltered(file, bookmarks)
	} else {
		wf.FatalError(err)
	}
}

func calibreEpubFile() string {
	usr, _ := user.Current()
	var path string
	calibreJsonFile := "~/Library/Preferences/calibre/viewer.json"
	path = filepath.Join(usr.HomeDir, calibreJsonFile[2:])

	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		wf.FatalError(err)
	}
	var jsonData map[string][]string
	// JSON unmarshal returns some BOOL error
	_ = json.Unmarshal(fileBytes, &jsonData)
	return jsonData["viewer_open_history"][0]
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

func cacheLastQuery(queryString string) {
	if err := wf.Cache.StoreJSON("LastQuery", queryString); err != nil {
		wf.FatalError(err)
	}
}

func getCurrentEpub() {
	fmt.Println(calibreEpubFile())
}

func getLastQuery() string {
	var lastQuery string
	if err := wf.Cache.LoadJSON("LastQuery", &lastQuery); err != nil {
		wf.FatalError(err)
	}
	return lastQuery
}

// ImportFiles: import file or all files in folder
func ImportFile(db Database, file string) {
	if strings.HasSuffix(file, ".epub") || strings.HasSuffix(file, ".pdf") {

		_, err := db.GetFileFromPath(file)

		if err == dbr.ErrNotFound {
			fmt.Println("trying:", file)
			fileRecord, changed, _ := db.GetFile(file, false)

			if fileRecord.ID > 1 && changed == true {
				bookmarks, _ := db.BookmarksForFile(file)
				if len(bookmarks) != 0 {
					log.Println("Imported:", fileRecord.FileName)
				}
			}
		}
	}
}

// ImportFiles: import file or all files in folder
func ImportFiles(file string) {
	dbFile := filepath.Join(wf.DataDir(), dbFileName)
	db := initDatabase(dbFile)

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
			ImportFile(db, aFile)
			return nil
		})
	case mode.IsRegular():
		aFile, _ := filepath.Abs(file)
		ImportFile(db, aFile)
	}
}

// receive bookmark title as query from script filter and open calibre
func openCalibreBookmark(query string, file string) {
	command := "/Applications/calibre.app/Contents/MacOS/ebook-viewer"
	if file == "" {
		file = calibreEpubFile()
	}
	file = "\"" + file + "\"" // for shell script
	// TODO: "--continue" needed?
	cmdArgs := []string{"--open-at=toc:\"" + query + "\"", file}
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

func printLastQuery() {
	fmt.Println(getLastQuery())
}

// Query database for all bookmarks
func searchAllBookmarks(query string) {
	dbFile := filepath.Join(wf.DataDir(), dbFileName)
	db := initDatabaseForReading(dbFile)

	cacheLastQuery(query)
	results, err := db.searchAll(query)
	if err != nil {
		wf.FatalError(err)
	}
	returnSearchAllResults(results)
}

func returnBookmarksForFile(file string, bookmarks []*Bookmark) {
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

			wf.NewItem(bookmarks[i].Title).
				Subtitle(subtitle).
				UID(strconv.FormatInt(bookmarks[i].ID, 10)).
				Valid(true).
				Icon(icon).
				Arg(destination)
		}
	} else {
		icon = &aw.Icon{Value: "org.idpf.epub-container", Type: aw.IconTypeFileType}

		for i := range bookmarks {
			wf.NewItem(bookmarks[i].Title).
				Subtitle(bookmarks[i].Section.String).
				UID(strconv.FormatInt(bookmarks[i].ID, 10)).
				Valid(true).
				Icon(icon).
				Arg(bookmarks[i].Title)
		}

	}
	wf.SendFeedback()
}

func returnBookmarksForFileFiltered(file string, bookmarks []*SearchAllResult) {
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

			wf.NewItem(bookmarks[i].Title).
				Subtitle(subtitle).
				UID(strconv.FormatInt(bookmarks[i].ID, 10)).
				Valid(true).
				Icon(icon).
				Arg(bookmarks[i].Destination)
		}
	} else {
		icon = &aw.Icon{Value: "org.idpf.epub-container", Type: aw.IconTypeFileType}

		for i := range bookmarks {
			wf.NewItem(bookmarks[i].Title).
				Subtitle(bookmarks[i].Section.String).
				UID(strconv.FormatInt(bookmarks[i].ID, 10)).
				Valid(true).
				Icon(icon).
				Arg(bookmarks[i].Title)
		}
	}
	wf.SendFeedback()
}

// Parse database search results and return items to Alfred
func returnSearchAllResults(bookmarks []*SearchAllResult) {
	var title string
	var subtitle string
	var arg string
	for i := range bookmarks {
		icon := iconForFileID(bookmarks[i].FileID, bookmarks[i].Path)

		if bookmarks[i].Section.String != "" {
			title = bookmarks[i].Title + " | " + bookmarks[i].Section.String
		} else {
			title = bookmarks[i].Title
		}

		if strings.HasSuffix(bookmarks[i].FileName, ".pdf") {
			subtitle = "Page " + bookmarks[i].Destination + ". " + bookmarks[i].FileName
			arg = bookmarks[i].Path + "/Page:" + bookmarks[i].Destination
		} else {
			subtitle = bookmarks[i].FileName
			arg = bookmarks[i].Path + "/Page:" + bookmarks[i].Title
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

func TestStuff(file string) {
	//bookmarks, _ := bookmarksForPDF(path)
	//fmt.Println("bookmarks", len(bookmarks))
	f := File{}
	if err := f.Init(file); err != nil {
		log.Println(err)
	}
	_ = f.Init(file)
	bookmarks2, _ := f.Bookmarks()
	fmt.Println("Bookmarks", len(bookmarks2))
}

// UpdateFile: check one file for metadata updates, not including bookmarks
func UpdateFile(db Database, fileRecord *DatabaseFile) {
	updated, _ := db.UpdateMetadata(fileRecord)
	if updated == true {
		fmt.Println("Updated:", fileRecord.Path)
	}
}

// UpdateFiles: check passed file or all files for metadata changes, not including bookmarks
func UpdateFiles(file string) {
	dbFile := filepath.Join(wf.DataDir(), dbFileName)
	db := initDatabase(dbFile)

	if _, err := os.Stat(file); err == nil {
		fileRecord, _, _ := db.GetFile(file, false)
		UpdateFile(db, fileRecord)
	} else {
		files, err := db.AllFiles()
		if err != nil {
			log.Fatal(err)
		}
		for _, aFile := range files {
			if fileExists(aFile.Path) {
				UpdateFile(db, aFile)
			}
		}
	}
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
	case options.Epub:
		bookmarksForFileEpub(query)
	case options.Import:
		ImportFiles(options.File)
	case options.Getepub:
		getCurrentEpub()
	case options.Openepub:
		openCalibreBookmark(query, options.File)
	case options.Lastquery:
		printLastQuery()
	case options.Test:
		TestStuff(options.File)
	case options.Update:
		UpdateFiles(options.File)
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
