package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
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
    openepub	open calibre to bookmark
    pdf		Retrieve or filter bookmarks for opened PDF in Acrobat, Preview, or Skim.
    lastquery	Retrieve cached last query string for script filter
    test        Testing stuff
    update      Update file metadata
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
	wf = aw.New(update.GitHub(repo), aw.HelpURL(fmt.Sprintf("%s%s", repo, "/issues")))

	coversCacheDir = filepath.Join(wf.DataDir(), "covers")
}

// initDatabase: initialize SQLite database
func initDatabase() Database {
	dbFile := filepath.Join(wf.DataDir(), dbFileName)
	db := Database{}
	db.Init(dbFile)
	return db
}

// Bookmarks all for file, from database or imported, return results
func bookmarksForFile(file string) {
	// ensure EPUB or PDF file exists
	if _, err := os.Stat(file); err != nil {
		wf.FatalError(err)
	}
	db := initDatabase()

	bookmarks, err := db.BookmarksForFile(file)
	if err == nil {
		returnBookmarksForFile(file, bookmarks)
	} else {
		wf.FatalError(err)
	}
}

func bookmarksForFileEpub(query string) {
	epub := calibreEpubFile()
	if query == "" {
		bookmarksForFile(epub)
	} else {
		bookmarksForFileFiltered(epub, query)
	}
}

// Bookmarks filtered for file, from database or imported, return results
func bookmarksForFileFiltered(file string, query string) {
	db := initDatabase()
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
	if strings.HasPrefix(calibreJsonFile, "~/") {
		path = filepath.Join(usr.HomeDir, calibreJsonFile[2:])
	}

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

	var icon *aw.Icon
	if _, err := os.Stat(iconFile); err == nil {
		icon = &aw.Icon{
			Value: iconFile,
			Type:  aw.IconTypeImage,
		}
	} else {
		icon = &aw.Icon{
			Value: filePath,
			Type:  aw.IconTypeFileIcon,
		}
	}
	return icon
}

func cacheLastQuery(queryString string) {
	var (
		name  = "LastQuery"
		value = queryString
	)
	cache := aw.NewCache(wf.CacheDir())
	data, _ := json.Marshal(value)
	if err := cache.Store(name, data); err != nil {
		panic(err)
	}
}

func getCurrentEpub() {
	file := calibreEpubFile()
	fmt.Println(file)
}

func getLastQuery() string {
	var (
		name = "LastQuery"
	)
	cache := aw.NewCache(wf.CacheDir())
	data, err := cache.Load(name)
	if err != nil {
		panic(err)
	}

	var lastQuery string
	err = json.Unmarshal(data, &lastQuery)
	return lastQuery
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
	defer file.Close()
	_, _ = file.WriteString("#!/bin/sh\n")
	_, _ = file.WriteString(command + " " + strings.Join(cmdArgs, " "))
	_, _ = file.WriteString("\n")
	_ = os.Chdir(temp)
	_ = exec.Command("sh", "alfred-jnana.sh").Start()
}

func printLastQuery() {
	lastQuery := getLastQuery()
	fmt.Println(lastQuery)
}

// Query database for all bookmarks
func searchAllBookmarks(query string) {
	db := initDatabase()
	cacheLastQuery(query)
	results, err := db.searchAll(query)
	if err != nil {
		wf.FatalError(err)
	}
	returnSearchAllResults(results)
}

func returnBookmarksForFile(file string, bookmarks []*Bookmark) {
	var icon *aw.Icon
	var destination string
	var subtitle string
	pdf := false

	if strings.HasSuffix(file, "pdf") {
		icon = &aw.Icon{Value: "com.adobe.pdf", Type: aw.IconTypeFileType}
		pdf = true
	} else {
		icon = &aw.Icon{Value: "org.idpf.epub-container", Type: aw.IconTypeFileType}
	}

	for _, bookmark := range bookmarks {
		if pdf == true {
			destination = bookmark.Destination
			if bookmark.Section.String != "" {
				subtitle = fmt.Sprintf("Page %s. %s", bookmark.Destination, bookmark.Section.String)
			} else {
				subtitle = fmt.Sprintf("Page %s", bookmark.Destination)
			}
		} else {
			destination = bookmark.Title
			subtitle = bookmark.Section.String
		}
		wf.NewItem(bookmark.Title).
			Subtitle(subtitle).
			UID(fmt.Sprintf("%d", bookmark.ID)).
			Valid(true).
			Icon(icon).
			Arg(destination)
	}
	wf.SendFeedback()
}

func returnBookmarksForFileFiltered(file string, bookmarks []*SearchAllResult) {
	var icon *aw.Icon
	var destination string
	var subtitle string
	pdf := false

	if strings.HasSuffix(file, "pdf") {
		icon = &aw.Icon{Value: "com.adobe.pdf", Type: aw.IconTypeFileType}
		pdf = true
	} else {
		icon = &aw.Icon{Value: "org.idpf.epub-container", Type: aw.IconTypeFileType}
	}

	for _, bookmark := range bookmarks {
		if pdf == true {
			destination = bookmark.Destination
			if bookmark.Section.String != "" {
				subtitle = fmt.Sprintf("Page %s. %s", bookmark.Destination, bookmark.Section.String)
			} else {
				subtitle = fmt.Sprintf("Page %s", bookmark.Destination)
			}
		} else {
			destination = bookmark.Title
			subtitle = bookmark.Section.String
		}
		wf.NewItem(bookmark.Title).
			Subtitle(subtitle).
			UID(fmt.Sprintf("%d", bookmark.ID)).
			Valid(true).
			Icon(icon).
			Arg(destination)
	}
	wf.SendFeedback()
}

// Parse database search results and return items to Alfred
func returnSearchAllResults(bookmarks []*SearchAllResult) {
	var title string
	var subtitle string
	var arg string
	for _, bookmark := range bookmarks {
		icon := iconForFileID(bookmark.FileID, bookmark.Path)

		if bookmark.Section.String != "" {
			title = fmt.Sprintf("%s | %s", bookmark.Title, bookmark.Section.String)
		} else {
			title = bookmark.Title
		}
		if strings.HasSuffix(bookmark.FileName, ".pdf") {
			subtitle = fmt.Sprintf("Page %s. %s", bookmark.Destination, bookmark.FileName)
			arg = fmt.Sprintf("%s/Page:%s", bookmark.Path, bookmark.Destination)
		} else {
			subtitle = bookmark.FileName
			arg = fmt.Sprintf("%s/Page:%s", bookmark.Path, bookmark.Title)
		}
		wf.NewItem(title).
			Subtitle(subtitle).
			UID(fmt.Sprintf("%d", bookmark.ID)).
			Valid(true).
			Icon(icon).
			Arg(arg)
	}
	wf.SendFeedback()
}

func TestStuff(file string) {
	//bookmarks, _ := bookmarksForPDF(file)
	//fmt.Println("bookmarks", len(bookmarks))
	bookmarks2, _ := FileBookmarks(file)
	fmt.Println("Bookmarks", len(bookmarks2))

	//if cmp.Equal(bookmarks, bookmarks2) {
	//	fmt.Println("bookmarks:", len(bookmarks), len(bookmarks2), file, " equal")
	//} else {
	//	fmt.Println("bookmarks:", len(bookmarks), len(bookmarks2), file, " NOT EQUAL")
	//}
	//for i := range bookmarks {
	//	if !cmp.Equal(bookmarks[i], bookmarks2[i]) {
	//		fmt.Println("Title", bookmarks[i].Title, "/", bookmarks2[i].Title)
	//		fmt.Println("Section", bookmarks[i].Section, "/", bookmarks2[i].Section)
	//		fmt.Println("Dest", bookmarks[i].Destination, "/", bookmarks2[i].Destination)
	//	}
	//}
}

// UpdateFile: check one file for metadata updates, not including bookmarks
func UpdateFile(db Database, fileRecord File) {
	updated, err := db.UpdateFileCheck(fileRecord)
	if err != nil {
		fmt.Println(fmt.Sprintf("Update error: %s, %s", fileRecord.Path, err))
	}
	if updated == true {
		fmt.Println("Updated:", fileRecord.Path)
	}
}

// UpdateFiles: check passed file or all files for metadata changes, not including bookmarks
func UpdateFiles(file string) {
	db := initDatabase()

	if _, err := os.Stat(file); err == nil {
		fileRecord, _, _ := db.GetFile(file)
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

	if options.All == true {
		searchAllBookmarks(query)
	}
	if options.Bm == true {
		bookmarksForFile(options.File)
	}
	if options.Bmf == true {
		bookmarksForFileFiltered(options.File, query)
	}
	if options.Epub == true {
		bookmarksForFileEpub(query)
	}
	if options.Getepub == true {
		getCurrentEpub()
	}
	if options.Openepub == true {
		openCalibreBookmark(query, options.File)
	}
	if options.Lastquery == true {
		printLastQuery()
	}
	if options.Test == true {
		TestStuff(options.File)
	}
	if options.Update == true {
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
