package main

import (
	"database/sql"
	"log"
	"os"
	"spidy/crawler"
	"spidy/models"
	"spidy/utils"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	utils.SetupLogger()

	dbFile := "./spidy.db"
	_, err := os.Stat(dbFile)
	isNewDb := os.IsNotExist(err)

	utils.Logger.Infof("Opening database from %s", dbFile)
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if isNewDb {
		utils.Logger.Info("Seeding database")
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
		utils.Logger.Info("Seeding done")
	}

	crawlTicker := time.NewTicker(1 * time.Second)
	defer crawlTicker.Stop()

	for {
		select {
		case <-crawlTicker.C:
			result := crawler.Crawl(db)
			if result {
				utils.Logger.Info("No URLs to crawl. Exiting.")
				os.Exit(0)
			}
		}
	}
}
