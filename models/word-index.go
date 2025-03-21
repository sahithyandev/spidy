package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type WordIndex struct {
	WordId        string  `json:"id"`
	Word          string  `json:"word"`
	TermFrequency float64 `json:"term_frequency"`
	PageId        string  `json:"from_page_id"`
}

func SeedWordIndex(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS word_index (
		word_id TEXT PRIMARY KEY,
		word TEXT,
		term_frequency FLOAT,
		page_id TEXT)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
