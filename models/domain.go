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

func CreateDomain(db *sql.DB, domain string) Domain {
	query := `INSERT INTO domain (domain) VALUES (?) RETURNING domain, added_at`
	row := db.QueryRow(query, domain)

	var d Domain
	err := row.Scan(&d.Domain, &d.AddedAt)
	if err != nil {
		panic(err)
	}
	return d
}

func GetDomain(db *sql.DB, domain string) (Domain, error) {
	query := `SELECT domain, added_at FROM domain WHERE domain = ?`
	row := db.QueryRow(query, domain)

	var d Domain
	err := row.Scan(&d.Domain, &d.AddedAt)
	if err != nil {
		return d, err
	}
	return d, nil
}
