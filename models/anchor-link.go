package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type AnchorLink struct {
	Id         string    `json:"id"`
	Text       string    `json:"text"`
	FromPageId string    `json:"from_page_id"`
	ToPageId   string    `json:"to_page_id"`
	IsInternal bool      `json:"is_internal"`
	CrawledAt  time.Time `json:"crawled_at"`
}

func SeedAnchorLink(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS anchor_links (
		id TEXT PRIMARY KEY,
		text TEXT,
		from_page_id TEXT,
		to_page_id TEXT,
		is_internal BOOLEAN,
		crawled_at TIMESTAMP)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
