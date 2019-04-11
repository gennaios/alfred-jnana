package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/gen2brain/go-fitz"
)

type FileBookmark struct {
	Title       string `json:"title"`
	Section     string `json:"section"`
	Destination string `json:"destination"`
	Uri         string
}

// FileBookmarks: for EPUB and PDF file, using go-fitz
func FileBookmarks(file string) ([]*FileBookmark, error) {
	doc, err := fitz.New(file)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer doc.Close()

	outlines, err := doc.ToC()
	if err != nil {
		fmt.Println("error:", err)
	}
	return parseBookmarks(outlines), err
}

// FileBookmarks: for EPUB and PDF file, using go-fitz
func FileMetadata(file string) map[string]string {
	doc, err := fitz.New(file)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer doc.Close()
	return doc.Metadata()
}

// FBookmarks for EPUB and PDF file, using Python script ./pdf.py
func bookmarksForPDF(file string) ([]*FileBookmark, error) {
	cmdArgs := []string{"FileBookmarks", file}

	output, err := exec.Command("./pdf.py", cmdArgs...).Output()
	if err != nil {
		log.Fatal(err)
	}

	// JSON stdout as []bytes, convert before return
	var bookmarks []*FileBookmark
	bookmarks, err = bookmarksFromJson(output)

	return bookmarks, err
}

// Take JSON []bytes and return as slice of FileBookmark structs
func bookmarksFromJson(jsonBytes []byte) ([]*FileBookmark, error) {
	var bookmarks []*FileBookmark

	err := json.Unmarshal(jsonBytes, bookmarks)
	if err != nil {
		fmt.Println(err)
	}
	return bookmarks, err
}

// Parse bookmarks from go-fitz
func parseBookmarks(outline []fitz.Outline) []*FileBookmark {
	sections := []string{"", "", "", "", "", "", "", "", "", "", "", "", ""}
	var parsedBookmarks []*FileBookmark
	var page int
	var currentLevel int
	var section string
	var title string

	for _, bookmark := range outline {
		if bookmark.Page != -1 {
			page = bookmark.Page + 1
		} else {
			page = -1
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
			Destination: fmt.Sprintf("%d", page),
			Uri:         strings.TrimSpace(bookmark.URI),
		}
		parsedBookmarks = append(parsedBookmarks, &newBookmark)
	}
	return parsedBookmarks
}
