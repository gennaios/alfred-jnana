package fulltext

import (
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"jnana/internal/database"
	"jnana/models"
)

type BookmarkSearch struct {
	ID        int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title     null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Section   null.String `boil:"section" json:"section,omitempty" toml:"section" yaml:"section,omitempty"`
	Name      string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	FileTitle string      `boil:"file_title" json:"file_title" toml:"file_title" yaml:"file_title"`
	Creator   string      `boil:"creator" json:"creator" toml:"creator" yaml:"creator"`
	Subject   string      `boil:"subject" json:"subject" toml:"subject" yaml:"subject"`
	Publisher string      `boil:"publisher" json:"publisher" toml:"publisher" yaml:"publisher"`
}

// BookmarksCreate create FTS entries
func BookmarksCreate(db *database.Database, file models.File) {
	var results []*BookmarkSearch

	_ = queries.Raw(`SELECT
		id, title, section, name, file_title, creator, subject, publisher from bookmark_view
		WHERE file_id = ?
		ORDER BY id`, file.ID).Bind(db.Ctx, db.Db, &results)

	for i := range results {
		_, _ = queries.Raw(`INSERT INTO
			bookmark_search(rowid, title, section, name, file_title, creator, subject, publisher)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?)`,
			results[i].ID,
			results[i].Title,
			results[i].Section,
			results[i].Name,
			results[i].FileTitle,
			results[i].Creator,
			results[i].Subject,
			results[i].Publisher,
		).Exec(db.Db)
	}
}

// BookmarksDelete delete FTS entries
func BookmarksDelete(db *database.Database, file models.File) {
	var results []*BookmarkSearch

	_ = queries.Raw(`SELECT
		id, title, section, name, file_title, creator, subject, publisher from bookmark_view
		WHERE file_id = ?
		ORDER BY id`, file.ID).Bind(db.Ctx, db.Db, &results)

	for i := range results {
		_, _ = queries.Raw(`INSERT INTO 
			bookmark_search(bookmark_search, rowid, title, section, name, file_title, creator, subject, publisher)
			VALUES('delete', ?, ?, ?, ?, ?, ?, ?, ?)`,
			results[i].ID,
			results[i].Title,
			results[i].Section,
			results[i].Name,
			results[i].FileTitle,
			results[i].Creator,
			results[i].Subject,
			results[i].Publisher,
		).Exec(db.Db)
	}
}

// FileCreate create FTS entry
func FileCreate(db *database.Database, file models.File) {
	var results []*models.File

	_ = queries.Raw(`SELECT
		id, name, title, series, creator, publisher, subject, isbn, description from file_view
		WHERE id = ?`, file.ID).Bind(db.Ctx, db.Db, &results)

	_, _ = queries.Raw(`INSERT INTO
			file_search(rowid, name, title, series, creator, publisher, subject, isbn, description)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		results[0].ID,
		results[0].Name,
		results[0].Title,
		results[0].Series,
		results[0].Creator,
		results[0].Publisher,
		results[0].Subject,
		results[0].Isbn,
		results[0].Description,
	).Exec(db.Db)
}

// FileDelete delete FTS entry
func FileDelete(db *database.Database, file models.File) {
	var results []*models.File

	_ = queries.Raw(`SELECT
		id, name, title, series, creator, publisher, subject, isbn, description from file_view
		WHERE id = ?`, file.ID).Bind(db.Ctx, db.Db, &results)

	_, _ = queries.Raw(`INSERT INTO 
			file_search(file_search, rowid, name, title, series, creator, publisher, subject, isbn, description)
			VALUES('delete', ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		results[0].ID,
		results[0].Name,
		results[0].Title,
		results[0].Series,
		results[0].Creator,
		results[0].Publisher,
		results[0].Subject,
		results[0].Isbn,
		results[0].Description,
	).Exec(db.Db)
}
