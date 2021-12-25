package file

import (
	"fmt"
	"github.com/gen2brain/go-fitz"
	"github.com/meskio/epubgo"
	"io/ioutil"
	"jnana/internal/util"
	"strings"
)

type File struct {
	path string
	file *fitz.Document

	Epub *epubgo.Epub
	nav  *epubgo.NavigationIterator

	Title     string
	Creator   string
	Subject   string
	Publisher string
}

type Bookmark struct {
	Title       string `json:"title"`
	Section     string `json:"section"`
	Destination string `json:"destination"`
}

// Init EPUB or PDF path and fill struct fields with metadata
func (f *File) Init(file string) error {
	var err error
	f.path = file

	f.file, err = fitz.New(f.path)
	if err != nil {
		return err
	}

	if strings.HasSuffix(f.path, ".epub") {
		f.Epub, err = epubgo.Open(file)
		f.nav, err = f.Epub.Navigation()
	}
	f.Metadata()
	return err

}

// Bookmarks for EPUB and PDF path, using go-fitz
func (f *File) Bookmarks() ([]*Bookmark, error) {
	var outlines []fitz.Outline
	var err error

	outlines, err = f.file.ToC()
	if err != nil {
		fmt.Println("error:", err)
	}
	return f.parseBookmarks(f.path, outlines), err
}

func (f *File) CoverForEPUB() ([]byte, error) {
	metaCoverId, err := f.Epub.Metadata("cover")
	if err != nil {
		cover, _ := f.Epub.OpenFileId(metaCoverId[0])
		defer cover.Close()

		buff1, err := ioutil.ReadAll(cover)
		if err != nil {
			_ = util.Notification("Error EPUB cover:" + err.Error())
			return nil, err
		} else {
			return buff1, err
		}
	}
	return nil, nil
}

func (f *File) Metadata() {
	f.Title = ""
	f.Creator = ""
	f.Subject = ""
	f.Publisher = ""

	if strings.HasSuffix(f.path, ".pdf") {
		f.MetadataForPDF()
	} else if strings.HasSuffix(f.path, ".epub") {
		f.MetadataForEPUB()
	}
}

// MetadataForEPUB for PDF path, using go-fitz
func (f *File) MetadataForEPUB() {
	var title []string
	var creator []string
	var subject []string
	var publisher []string

	title, _ = f.Epub.Metadata("title")
	title = TrimMetadata(title)
	f.Title = strings.TrimSpace(strings.Join(title[:], "; "))

	creator, _ = f.Epub.Metadata("creator")
	creator = TrimMetadata(creator)
	f.Creator = strings.TrimSpace(strings.Join(creator[:], "; "))

	subject, _ = f.Epub.Metadata("subject")
	subject = TrimMetadata(subject)
	f.Subject = strings.ToLower(strings.TrimSpace(strings.Join(subject[:], ", ")))

	publisher, _ = f.Epub.Metadata("publisher")
	publisher = TrimMetadata(publisher)
	f.Publisher = strings.TrimSpace(strings.Join(publisher[:], "; "))
}

// MetadataForPDF for PDF path, using go-fitz
func (f *File) MetadataForPDF() {
	fileMetadata := f.file.Metadata()

	f.Title = strings.Trim(fileMetadata["title"], `'"; `)
	f.Creator = strings.Trim(fileMetadata["author"], `'"; `)
}

// Parse bookmarks from go-fitz
func (f *File) parseBookmarks(file string, outline []fitz.Outline) []*Bookmark {
	sections := []string{"", "", "", "", "", "", "", "", "", "", "", "", ""}
	var parsedBookmarks []*Bookmark
	var page int
	var currentLevel int
	var section string
	var title string
	var destination string

	pdf := true

	if strings.HasSuffix(file, ".epub") {
		pdf = false
	}

	for i, bookmark := range outline {
		// skip cover
		if i != 0 && (!strings.Contains(bookmark.Title, "Cover") || !strings.Contains(bookmark.Title, "cover") || !strings.Contains(bookmark.Title, "COVER") || !strings.Contains(bookmark.Title, "Couverture")) {
			if pdf == true {
				if bookmark.Page != -1 {
					page = bookmark.Page + 1
				} else {
					page = -1
				}
				destination = fmt.Sprintf("%d", page)
			} else {
				destination = strings.TrimSpace(bookmark.URI)
				// MuPDF: workaround for destination being full path instead of HREF
				destination = strings.Replace(destination, "OEBPS/", "", 1)
				destination = strings.Replace(destination, "Oebps/", "", 1)
				destination = strings.Replace(destination, "OPS/", "", 1)
				destination = strings.Replace(destination, "Ops/", "", 1)
			}

			title = strings.TrimSpace(bookmark.Title)

			section = ""
			sections[bookmark.Level] = title

			if bookmark.Level == 1 {
				section = ""
			} else {
				currentLevel = bookmark.Level
				for currentLevel > 1 {
					if section == "" {
						section = sections[currentLevel-1]
					} else {
						section = fmt.Sprintf("%s > %s", sections[currentLevel-1], section)
					}
					currentLevel -= 1
				}
			}

			newBookmark := Bookmark{
				Title:       title,
				Section:     section,
				Destination: destination,
			}
			parsedBookmarks = append(parsedBookmarks, &newBookmark)
		}
	}

	return parsedBookmarks
}

func TrimMetadata(slc []string) []string {
	for i := range slc {
		slc[i] = strings.Trim(slc[i], `'"; `)
	}
	return slc
}
