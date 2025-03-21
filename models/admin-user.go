package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type AdminUser struct {
	Username      string `json:"username"`
	Password_hash string `json:"password_hash"`
}

func SeedAdminUser(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS admin_users (
		username TEXT PRIMARY KEY,
		password_hash TEXT)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
