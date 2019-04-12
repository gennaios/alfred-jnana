package main

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
)

// 	"github.com/stretchr/testify/require"

func assertPragma(t *testing.T, row *sql.Row, expected string, pragma string) {
	var result string
	if err := row.Scan(&result); err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, result, expected, pragma)
}

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
