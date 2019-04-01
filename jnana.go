package main

import (
	"encoding/json"
	"fmt"
	"os"
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
    jnana bookmarks <file> [<query>]
    jnana epub [<query>]
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
    bookmarks	Retrieve bookmarks for file, all or filtered by query
    epub	Retrieve or filter bookmarks for EPUB in Calibre.
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
	Bookmarks bool
	Epub      bool
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

// Bookmark all or filtered for file, from database or imported, return results
func bookmarksForFile(file string, query string) {
	forFile(file, query)
	// TODO: pass to Alfred
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
			Type:  aw.IconTypeFileType,
		}
	} else {
		icon = &aw.Icon{
			Value: filePath,
			Type:  aw.IconTypeFileType,
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

func printLastQuery() {
	lastQuery := getLastQuery()
	fmt.Println(lastQuery)
}

// Query database for all bookmarks
func searchAllBookmarks(query string) {
	usr, _ := user.Current()
	dbFile := filepath.Join(usr.HomeDir, dataDir, dbFileName)
	conn := initDatabase(dbFile)

	cacheLastQuery(query)

	results, err := searchAll(conn, query)
	if err != nil {
		wf.FatalError(err)
	}

	returnBookmarks(results)
}

// Parse database search results and return items to Alfred
func returnBookmarks(bookmarks []SearchAllResult) {
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
	// fmt.Println(options)

	// normalize white space, remove dupes
	query := strings.Join(strings.Fields(strings.TrimSpace(options.Query)), " ")

	if options.All == true {
		searchAllBookmarks(query)
	}
	if options.Bookmarks == true {
		bookmarksForFile(options.File, options.Query)
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
