package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Domain struct {
	Domain  string    `json:"domain"`
	AddedAt time.Time `json:"added_at"`
}

func SeedDomain(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS domain (
		domain TEXT PRIMARY KEY,
		added_at TIMESTAMP NOT NULL DEFAULT current_timestamp)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
