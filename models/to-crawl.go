package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ToCrawl struct {
	Url        string    `json:"url"`
	Priority   int       `json:"priority"`
	CrawlAfter time.Time `json:"crawl_after"`
	AddedOn    time.Time `json:"added_on"`
}

func SeedToCrawl(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS to_crawl (
		url TEXT PRIMARY KEY,
		priority INTEGER NOT NULL DEFAULT 1,
		crawl_after TIMESTAMP NOT NULL DEFAULT current_timestamp,
		added_on TIMESTAMP NOT NULL DEFAULT current_timestamp)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	// Seed the table with some data
	query = `INSERT INTO to_crawl (url) VALUES ('https://sahithyan.dev')`
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func ChooseNextUrlToCrawl(db *sql.DB) string {
	query := `SELECT url FROM to_crawl WHERE crawl_after <= CURRENT_TIMESTAMP ORDER BY priority DESC, crawl_after ASC LIMIT 1`
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	url := ""
	if rows.Next() {
		rows.Scan(&url)
	}
	return url
}
