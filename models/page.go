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
		url TEXT UNIQUE,
		title TEXT,
		description TEXT ,
		page_rank FLOAT NOT NULL DEFAULT 0.0,
		crawled_at TIMESTAMP NOT NULL DEFAULT current_timestamp)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func AddPage(db *sql.DB, url string, title string, description string) {
	query := `INSERT INTO pages (id, url, title, description) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, HashUrl(url), url, title, description)
	if err != nil {
		panic(err)
	}
}
