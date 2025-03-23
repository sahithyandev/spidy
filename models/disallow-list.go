package models

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"

	_ "github.com/mattn/go-sqlite3"
)

type DisallowList struct {
	Id            string `json:"id"`
	Domain        string `json:"domain"`
	DisallowedUrl string `json:"disallowed_url"`
}

func SeedDisallowList(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS disallow_list (
		id TEXT PRIMARY KEY,
		domain TEXT,
		disallowed_url TEXT
	)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func HashDisallowList(domain string, disallowedUrl string) string {
	hasher := sha1.New()
	hasher.Write([]byte(domain))
	hasher.Write([]byte(disallowedUrl))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func AddDisallowListEntry(db *sql.DB, domain string, disallowedUrl string) {
	query := `INSERT INTO disallow_list (id, domain, disallowed_url) VALUES (?,?,?)`
	_, err := db.Exec(query, HashDisallowList(domain, disallowedUrl), domain, disallowedUrl)
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

func GetDisallowList(db *sql.DB, domain string) []string {
	query := `SELECT disallowed_url FROM disallow_list WHERE domain = ?`
	rows, err := db.Query(query, domain)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	disallowedUrls := []string{}
	for rows.Next() {
		var disallowedUrl string
		rows.Scan(&disallowedUrl)
		disallowedUrls = append(disallowedUrls, disallowedUrl)
	}
	return disallowedUrls
}
