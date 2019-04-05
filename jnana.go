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
    jnana epubf <query>
    jnana openepub [<file>] <query>
    jnana pdf <file> [<query>]
    jnana lastquery
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
    openepub	open calibre to bookmark
    pdf		Retrieve or filter bookmarks for opened PDF in Acrobat, Preview, or Skim.
    lastquery	Retrieve cached last query string for script filter
`

	wf *aw.Workflow

	dataDir        = "Library/Application Support/Alfred 3/Workflow Data/io.github.gennaios.gnosis"
	dbFileName     = "gnosis.db"
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
	Epubf     bool
	Openepub  bool
	Pdf       bool
	Lastquery bool

	// parameters
	Query  string
	File   string
	Fileid string
}

func init() {
	// Create a new Workflow using default settings.
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))

	usr, _ := user.Current()
	coversCacheDir = filepath.Join(usr.HomeDir, dataDir, "covers")
}

// Bookmarks all for file, from database or imported, return results
func bookmarksForFile(file string) {
	if _, err := os.Stat(file); err != nil {
		wf.FatalError(err)
	}

	usr, _ := user.Current()
	dbFile := filepath.Join(usr.HomeDir, dataDir, dbFileName)
	db := Database{}
	db.Init(dbFile)

	bookmarks, err := db.BookmarksForFile(file)
	if err == nil {
		if strings.HasSuffix(file, "pdf") {
			returnBookmarksForPdf(file, bookmarks)
		} else {
			returnBookmarksForEpub(bookmarks)
		}
	} else {
		wf.FatalError(err)
	}
}

func bookmarksForFileEpub() {
	epub := calibreEpubFile()
	bookmarksForFile(epub)
}

// Bookmarks filtered for file, from database or imported, return results
func bookmarksForFileFiltered(file string, query string) {
	usr, _ := user.Current()
	dbFile := filepath.Join(usr.HomeDir, dataDir, dbFileName)

	db := Database{}
	db.Init(dbFile)
	bookmarks, err := db.BookmarksForFileFiltered(file, query)

	if err == nil {
		returnBookmarksForPdfFiltered(file, bookmarks)
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

func iconForFileID(fileId string, filePath string) *aw.Icon {
	iconFile := filepath.Join(coversCacheDir, fileId+".jp2")

	var icon *aw.Icon

	if _, err := os.Stat(iconFile); err == nil {
		icon = &aw.Icon{
			Value: iconFile,
			Type:  aw.IconTypeImage,
		}
	} else if os.IsNotExist(err) {
		icon = &aw.Icon{
			Value: filePath,
			Type:  aw.IconTypeFileIcon,
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
		// Cache "key" (filename) and the value to store
		name  = "LastQuery"
		value = queryString
	)

	cache := aw.NewCache(wf.CacheDir())

	// The API uses bytes
	data, _ := json.Marshal(value)

	if err := cache.Store(name, data); err != nil {
		panic(err)
	}
}

func getLastQuery() string {
	var (
		// Cache "key" (filename) and the value to store
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
func openCalibreBookmark(file string, query string) {
	command := "/Applications/calibre.app/Contents/MacOS/ebook-viewer"
	if file == "false" {
		file = calibreEpubFile()
	}
	// TODO: "--continue" needed?
	cmdArgs := []string{"--open-at=toc:\"" + query + "\"", file}

	_, err := exec.Command(command, cmdArgs...).Output()
	if err != nil {
		log.Fatal(err)
	}
}

func printLastQuery() {
	lastQuery := getLastQuery()
	fmt.Println(lastQuery)
}

// Query database for all bookmarks
func searchAllBookmarks(query string) {
	usr, _ := user.Current()
	dbFile := filepath.Join(usr.HomeDir, dataDir, dbFileName)
	db := Database{}
	db.Init(dbFile)

	cacheLastQuery(query)

	results, err := db.searchAll(query)
	if err != nil {
		wf.FatalError(err)
	}
	returnSearchAllResults(results)
}

func returnBookmarksForEpub(bookmarks []BookmarkRecord) {
	var icon *aw.Icon
	icon = &aw.Icon{Value: "org.idpf.epub-container", Type: aw.IconTypeFileType}

	for _, bookmark := range bookmarks {
		section := ""
		if bookmark.Section.String != "" {
			section = bookmark.Section.String
		}

		wf.NewItem(bookmark.Title).
			Subtitle(section).
			UID(strconv.FormatInt(bookmark.ID, 10)).
			Valid(true).
			Icon(icon).
			Arg(bookmark.Destination)
	}
	wf.SendFeedback()
}

func returnBookmarksForPdf(file string, bookmarks []BookmarkRecord) {
	var icon *aw.Icon
	if strings.HasSuffix(file, "pdf") {
		icon = &aw.Icon{Value: "com.adobe.pdf", Type: aw.IconTypeFileType}
	} else {
		icon = &aw.Icon{Value: "org.idpf.epub-container", Type: aw.IconTypeFileType}
	}

	for _, bookmark := range bookmarks {
		subtitleSuffix := ""
		if bookmark.Section.String != "" {
			subtitleSuffix = ". " + bookmark.Section.String
		}

		wf.NewItem(bookmark.Title).
			Subtitle("Page " + bookmark.Destination + subtitleSuffix).
			UID(strconv.FormatInt(bookmark.ID, 10)).
			Valid(true).
			Icon(icon).
			Arg(bookmark.Destination)
	}
	wf.SendFeedback()
}

func returnBookmarksForPdfFiltered(file string, bookmarks []SearchAllResult) {
	var icon *aw.Icon
	if strings.HasSuffix(file, "pdf") {
		icon = &aw.Icon{Value: "com.adobe.pdf", Type: aw.IconTypeFileType}
	} else {
		icon = &aw.Icon{Value: "org.idpf.epub-container", Type: aw.IconTypeFileType}
	}

	for _, bookmark := range bookmarks {
		subtitleSuffix := ""
		if bookmark.Section.String != "" {
			subtitleSuffix = ". " + bookmark.Section.String
		}

		wf.NewItem(bookmark.Title).
			Subtitle("Page " + bookmark.Destination + subtitleSuffix).
			UID(strconv.FormatInt(bookmark.ID, 10)).
			Valid(true).
			Icon(icon).
			Arg(bookmark.Destination)
	}

	wf.SendFeedback()
}

// Parse database search results and return items to Alfred
func returnSearchAllResults(bookmarks []SearchAllResult) {
	for _, bookmark := range bookmarks {
		uid := strconv.FormatInt(bookmark.ID, 10)
		icon := iconForFileID(bookmark.FileID, bookmark.Path)

		var title string
		if bookmark.Section.String != "" {
			title = bookmark.Title + " | " + bookmark.Section.String
		} else {
			title = bookmark.Title
		}

		var arg string
		var subtitle string
		if strings.HasSuffix(bookmark.FileName, ".epub") {
			subtitle = bookmark.FileName
			arg = bookmark.Path + "/Page:\"" + bookmark.Destination + "\""
		} else {
			subtitle = "Page " + bookmark.Destination + ". " + bookmark.FileName
			arg = bookmark.Path + "/Page:" + bookmark.Destination
		}

		wf.NewItem(title).
			Subtitle(subtitle).
			UID(uid).
			Valid(true).
			Icon(icon).
			Arg(arg)
	}

	wf.SendFeedback()
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
		bookmarksForFileFiltered(options.File, options.Query)
	}
	if options.Epub == true {
		bookmarksForFileEpub()
	}
	if options.Openepub == true {
		openCalibreBookmark(options.File, options.Query)
	}
	if options.Lastquery == true {
		printLastQuery()
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
