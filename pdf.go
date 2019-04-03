package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type Bookmark struct {
	Title       string `json:"title"`
	Section     string `json:"section"`
	Destination string `json:"destination"`
}

// Bookmark for PDF file, from Python script ./pdf.py
func bookmarksForPDF(file string) ([]Bookmark, error) {
	cmdArgs := []string{"bookmarks", file}

	output, err := exec.Command("./pdf.py", cmdArgs...).Output()
	if err != nil {
		log.Fatal(err)
	}

	// JSON stdout as []bytes, convert before return
	var bookmarks []Bookmark
	bookmarks, err = bookmarksFromJson(output)

	return bookmarks, err
}

// Take JSON []bytes and return as slice of Bookmark structs
func bookmarksFromJson(jsonBytes []byte) ([]Bookmark, error) {
	var bookmarks []Bookmark

	err := json.Unmarshal(jsonBytes, &bookmarks)
	if err != nil {
		fmt.Println(err)
	}

	return bookmarks, err
}
