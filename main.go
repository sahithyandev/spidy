package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"spidy/crawler"
	"spidy/models"
	"time"

	"github.com/PuerkitoBio/goquery"
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
			crawler.Crawl(db)
		}
	}
	chosenUrl := ""
	htmlResponse, err := crawler.Fetch(chosenUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer htmlResponse.Body.Close()

	doc, err := goquery.NewDocumentFromReader(htmlResponse.Body)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		innerText := s.Text()
		link, _ := s.Attr("href")
		fmt.Printf("%s --> %s\n", innerText, link)
	})
}
