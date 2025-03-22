package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"spidy/crawler"
	"spidy/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbFile := "./spidy.db"
	_, err := os.Stat(dbFile)
	isNewDb := os.IsNotExist(err)

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if isNewDb {
		fmt.Println("Seeding newly created database at " + dbFile)
		models.SeedAdminUser(db)
		models.SeedAnchorLink(db)
		models.SeedWordIndex(db)
		models.SeedToCrawl(db)
		models.SeedPage(db)
		models.SeedDomain(db)
		models.SeedDisallowList(db)
		err := db.Ping()
		if err != nil {
			log.Fatal(err)
		}
	}

	crawlTicker := time.NewTicker(1 * time.Second)
	defer crawlTicker.Stop()

	for {
		select {
		case <-crawlTicker.C:
			result := crawler.Crawl(db)
			if result {
				os.Exit(0)
			}
		}
	}
}
