package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"spidy/crawler"
	"spidy/models"

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

	chosenUrl := models.ChooseNextUrlToCrawl(db)
	fmt.Println("Chosen URL: " + chosenUrl)
	hostname := crawler.UrlToHostname(chosenUrl)
	fmt.Println("Hostname: " + hostname)

	robotsTxtUrl := crawler.RobotsTxtUrl(chosenUrl)
	fmt.Println("Robots.txt URL: " + robotsTxtUrl)
	body, err := crawler.Fetch(robotsTxtUrl)
	if err != nil {
		log.Fatal(err)
	}

	disallowedList := crawler.ParseRobotsTxt(body, "Spidy")
	fmt.Println(disallowedList)
}
