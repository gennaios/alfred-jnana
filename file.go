package main

import (
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

type File struct {
	ID            int64
	Path          string         `db:"path"`
	FileName      string         `db:"file_name"`
	FileExtension string         `db:"file_extension"`
	Title         dbr.NullString `db:"title"`
	Authors       dbr.NullString `db:"authors"`
	Subjects      dbr.NullString `db:"subjects"`
	DateCreated   string         `db:"date_created"`
	DateModified  dbr.NullString `db:"file_modified_date"` // TODO: rename -> date_modified
	FileHash      string         `db:"hash"`
}
