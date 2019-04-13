package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

// run the bookmarksForFile function b.N times
func BenchmarkBookmarksForFile(b *testing.B) {
	file, _ := filepath.Abs("./tests/pdf.pdf")
	if _, err := os.Stat(file); err != nil {
		log.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		bookmarksForFile(file)
	}
}

// run the searchAllBookmarks function b.N times
func BenchmarkSearchAllBookmarks100(b *testing.B) {
	query := "emblica"
	for n := 0; n < b.N; n++ {
		searchAllBookmarks(query)
	}
}
