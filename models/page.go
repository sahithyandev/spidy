package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Id          string    `json:"id"`
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	PageRank    float64   `json:"page_rank"`
	CrawledAt   time.Time `json:"crawled_at"`
}

func SeedPage(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS pages (
		id TEXT PRIMARY KEY,
		url TEXT,
		title TEXT,
		description TEXT,
		content TEXT,
		page_rank FLOAT,
		crawled_at TIMESTAMP)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
