package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"spidy/crawler"
	"spidy/models"

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

	chosenUrl := models.ChooseNextUrlToCrawl(db)
	fmt.Println("Chosen URL: " + chosenUrl)
	hostname := crawler.UrlToHostname(chosenUrl)
	fmt.Println("Hostname: " + hostname)

	robotsTxtUrl := crawler.RobotsTxtUrl(chosenUrl)
	fmt.Println("Robots.txt URL: " + robotsTxtUrl)
	response, err := crawler.Fetch(robotsTxtUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := crawler.ResponseToString(response)
	disallowedList := crawler.ParseRobotsTxt(body, "Spidy")
	fmt.Println(disallowedList)

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
