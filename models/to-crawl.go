package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ToCrawl struct {
	Url      string    `json:"url"`
	Priority int       `json:"priority"`
	AddedOn  time.Time `json:"added_on"`
}

func SeedToCrawl(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS to_crawl (
		url TEXT PRIMARY KEY,
		priority INTEGER NOT NULL DEFAULT 1,
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

func AddToCrawlEntry(db *sql.DB, crawlUrl string) {
	query := `INSERT INTO to_crawl (url) VALUES (?)`
	_, err := db.Exec(query, crawlUrl)
	if err != nil {
		panic(err)
	}
}

func RemoveToCrawlEntry(db *sql.DB, crawlUrl string) {
	query := `DELETE FROM to_crawl WHERE url = ?`
	_, err := db.Exec(query, crawlUrl)
	if err != nil {
		panic(err)
	}
}

func ChooseNextUrlToCrawl(db *sql.DB) string {
	query := `SELECT COALESCE(
		(SELECT url FROM to_crawl ORDER BY priority DESC LIMIT 1),
		(SELECT url FROM pages ORDER BY crawled_at ASC LIMIT 1)
	)`
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
func IsUrlInToCrawl(db *sql.DB, url string) bool {
	query := `SELECT EXISTS(SELECT 1 FROM to_crawl WHERE url = ?)`
	row := db.QueryRow(query, url)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}

func IncreasePriority(db *sql.DB, url string) {
	query := `UPDATE to_crawl SET priority = priority + 1 WHERE url = ?`
	_, err := db.Query(query, url)
	if err != nil {
		panic(err)
	}
}
