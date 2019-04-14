package main

import (
	"fmt"
	"github.com/gen2brain/go-fitz"
	"github.com/meskio/epubgo"
	"strings"
)

type File struct {
	path string
	file *fitz.Document

	epub *epubgo.Epub
	nav  *epubgo.NavigationIterator

	title     string
	authors   string
	subjects  string
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

func (f *File) Metadata() {
	f.title = ""
	f.authors = ""
	f.subjects = ""
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
	var authors []string
	var subjects []string
	var publisher []string

	title, _ = f.epub.Metadata("title")
	title = trimMetadata(title)
	f.title = strings.TrimSpace(strings.Join(title[:], "; "))

	authors, _ = f.epub.Metadata("creator")
	authors = trimMetadata(authors)
	f.authors = strings.TrimSpace(strings.Join(authors[:], "; "))

	subjects, _ = f.epub.Metadata("subject")
	subjects = trimMetadata(subjects)
	f.subjects = strings.TrimSpace(strings.Join(subjects[:], "; "))

	publisher, _ = f.epub.Metadata("publisher")
	publisher = trimMetadata(publisher)
	f.publisher = strings.TrimSpace(strings.Join(publisher[:], "; "))
}

// MetadataForPDF: for PDF path, using go-fitz
func (f *File) MetadataForPDF() {
	fileMetadata := f.file.Metadata()

	f.title = strings.Trim(fileMetadata["title"], `'"; `)
	f.authors = strings.Trim(fileMetadata["author"], `'"; `)
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

	for _, bookmark := range outline {
		if pdf == true {
			if bookmark.Page != -1 {
				page = bookmark.Page + 1
			} else {
				page = -1
			}
			destination = fmt.Sprintf("%d", page)
		} else {
			destination = strings.TrimSpace(bookmark.URI)
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
	return parsedBookmarks
}

func trimMetadata(slc []string) []string {
	for i := range slc {
		slc[i] = strings.Trim(slc[i], `'"; `)
	}
	return slc
}
