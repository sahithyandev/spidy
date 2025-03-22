package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DisallowList struct {
	Domain        string `json:"domain"`
	DisallowedUrl string `json:"disallowed_url"`
}

func SeedDisallowList(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS disallow_list (
		domain TEXT PRIMARY KEY,
		disallowed_url TEXT
	)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func AddDisallowListEntry(db *sql.DB, domain string, disallowedUrl string) {
	query := `INSERT INTO disallow_list (domain, disallowed_url) VALUES (?, ?)`
	_, err := db.Exec(query, domain, disallowedUrl)
	if err != nil {
		panic(err)
	}
}

func RemoveDisallowList(db *sql.DB, domain string) {
	query := `DELETE FROM disallow_list WHERE domain = ?`
	_, err := db.Exec(query, domain)
	if err != nil {
		panic(err)
	}
}
