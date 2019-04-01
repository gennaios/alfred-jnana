package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type BookmarkJson struct {
	Title       string `json:"title"`
	Section     string `json:"section"`
	Destination string `json:"destination"`
}

// BookmarkJson for PDF file, from Python script ./pdf.py
func bookmarksForPDF(file string) ([]BookmarkJson, error) {
	cmdArgs := []string{"bookmarks", file}

	output, err := exec.Command("./pdf.py", cmdArgs...).Output()
	if err != nil {
		log.Fatal(err)
	}

	// JSON stdout as []bytes, convert before return
	var bookmarks []BookmarkJson
	bookmarks, err = bookmarksFromJson(output)

	return bookmarks, err
}

// Take JSON []bytes and return as slice of BookmarkJson structs
func bookmarksFromJson(jsonBytes []byte) ([]BookmarkJson, error) {
	var bookmarks []BookmarkJson

	err := json.Unmarshal(jsonBytes, &bookmarks)
	if err != nil {
		fmt.Println(err)
	}

	return bookmarks, err
}
