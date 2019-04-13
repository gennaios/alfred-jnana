package main

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestDatabase_GetFile(t *testing.T) {
	file, _ := filepath.Abs("./tests/pdf.pdf")
	db := initDatabase("memory")

	fileRecord, changed, _ := db.GetFile(file, false)

	assert.Equal(t, fileRecord.Path, file, "File not found by path")
	assert.Equal(t, changed, false, "File path changed failed")
}

func TestDatabase_GetFileFromHash(t *testing.T) {
	file, _ := filepath.Abs("./tests/pdf.pdf")
	db := initDatabase("memory")

	fileRecord, _ := db.GetFileFromHash("ebb031c3945e884e695dbc63c52a5efcd075375046c49729980073585ee13c52")

	assert.Equal(t, fileRecord.Path, file, "File not found by hash")
}
