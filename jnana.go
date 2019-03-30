package main

import (
	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/docopt/docopt-go"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	repo = "gennaios/alfred-jnana"

	usage = `jnana [command] [<query>...]

usage:
    jnana all [<query>]
    jnana allepub [<query>]
    jnana allpdf [<query>]
    jnana epub [<query>]
    jnana pdf <file> [<query>]
    jnana -h

options:
    -h --help          Show this message and exit.
    --version          Show workflow version and exit.

commands:
    all		Search all bookmarks.
    allepub	Search all EPUB bookmarks.
    allpdf	Search all PDF bookmarks.
    epub	Retrieve or filter bookmarks for EPUB in Calibre.
    pdf		Retrieve or filter bookmarks for opened PDF in Acrobat, Preview, or Skim.
`

	wf *aw.Workflow

	dataDir        = "Library/Application Support/Alfred 3/Workflow Data/io.github.gennaios.gnosis"
	dbFileName     = "gnosis.db"
	coversCacheDir string // directory generated icons are stored in

	query string
)

var options struct {
	// commands
	All     bool
	Allepub bool
	Allpdf  bool
	Epub    bool
	Pdf     bool

	// parameters
	Query string
	File  string
}

func init() {
	// Create a new Workflow using default settings.
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))

	usr, _ := user.Current()
	coversCacheDir = filepath.Join(usr.HomeDir, dataDir, "covers")
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

// Query database for all bookmarks
func searchAllBookmarks(query string) {
	usr, _ := user.Current()
	dbFile := filepath.Join(usr.HomeDir, dataDir, dbFileName)
	conn := initDatabase(dbFile)

	results, err := searchAll(conn, query)
	if err != nil {
		wf.FatalError(err)
	}

	returnBookmarks(results)

	//resultsJson, _ := json.Marshal(results)
	//fmt.Println("Results: ", string(resultsJson))
}

// Parse database search results and return items to Alfred
func returnBookmarks(bookmarks []SearchAllResult) {
	log.Printf("%d total bookmark(s)", len(bookmarks))

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
	if options.All == true {
		searchAllBookmarks(options.Query)
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

	// fmt.Println("global arguments:", args)
	runCommand()
}

func main() {
	// calls run()
	wf.Run(run)
}
