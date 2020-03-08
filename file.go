package main

import (
	"fmt"
	"github.com/gen2brain/go-fitz"
	"github.com/meskio/epubgo"
	"io/ioutil"
	"strings"
)

type File struct {
	path string
	file *fitz.Document

	epub *epubgo.Epub
	nav  *epubgo.NavigationIterator

	title     string
	creator   string
	subject   string
	publisher string
}

type FileBookmark struct {
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
		f.epub, err = epubgo.Open(file)
		f.nav, err = f.epub.Navigation()
	}
	f.Metadata()
	return err

}

// Bookmarks: for EPUB and PDF path, using go-fitz
func (f *File) Bookmarks() ([]*FileBookmark, error) {
	var outlines []fitz.Outline
	var err error

	outlines, err = f.file.ToC()
	if err != nil {
		fmt.Println("error:", err)
	}
	return f.parseBookmarks(f.path, outlines), err
}

func (f *File) CoverForEPUB() ([]byte, error) {
	metaCoverId, err := f.epub.Metadata("cover")
	if err != nil {
		cover, _ := f.epub.OpenFileId(metaCoverId[0])
		defer cover.Close()

		buff1, err := ioutil.ReadAll(cover)
		if err != nil {
			_ = notification("Error EPUB cover:" + err.Error())
			return nil, err
		} else {
			return buff1, err
		}
	}
	return nil, nil
}

func (f *File) Metadata() {
	f.title = ""
	f.creator = ""
	f.subject = ""
	f.publisher = ""

	if strings.HasSuffix(f.path, ".pdf") {
		f.MetadataForPDF()
	} else if strings.HasSuffix(f.path, ".epub") {
		f.MetadataForEPUB()
	}
}

// MetadataForEPUB: for PDF path, using go-fitz
func (f *File) MetadataForEPUB() {
	var title []string
	var creator []string
	var subject []string
	var publisher []string

	title, _ = f.epub.Metadata("title")
	title = trimMetadata(title)
	f.title = strings.TrimSpace(strings.Join(title[:], "; "))

	creator, _ = f.epub.Metadata("creator")
	creator = trimMetadata(creator)
	f.creator = strings.TrimSpace(strings.Join(creator[:], "; "))

	subject, _ = f.epub.Metadata("subject")
	subject = trimMetadata(subject)
	f.subject = strings.ToLower(strings.TrimSpace(strings.Join(subject[:], ", ")))

	publisher, _ = f.epub.Metadata("publisher")
	publisher = trimMetadata(publisher)
	f.publisher = strings.TrimSpace(strings.Join(publisher[:], "; "))
}

// MetadataForPDF: for PDF path, using go-fitz
func (f *File) MetadataForPDF() {
	fileMetadata := f.file.Metadata()

	f.title = strings.Trim(fileMetadata["title"], `'"; `)
	f.creator = strings.Trim(fileMetadata["author"], `'"; `)
}

// Parse bookmarks from go-fitz
func (f *File) parseBookmarks(file string, outline []fitz.Outline) []*FileBookmark {
	sections := []string{"", "", "", "", "", "", "", "", "", "", "", "", ""}
	var parsedBookmarks []*FileBookmark
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

			newBookmark := FileBookmark{
				Title:       title,
				Section:     section,
				Destination: destination,
			}
			parsedBookmarks = append(parsedBookmarks, &newBookmark)
		}
	}

	return parsedBookmarks
}

func trimMetadata(slc []string) []string {
	for i := range slc {
		slc[i] = strings.Trim(slc[i], `'"; `)
	}
	return slc
}
