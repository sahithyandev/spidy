package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DisallowList struct {
	Domain        string    `json:"domain"`
	DisallowedUrl string    `json:"disallowed_url"`
	AddedAt       time.Time `json:"added_at"`
}

func SeedDisallowList(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS disallow_list (
		domain TEXT PRIMARY KEY,
		disallowed_url TEXT,
		added_at TIMESTAMP)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
