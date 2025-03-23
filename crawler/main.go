package crawler

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"spidy/models"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
)

func Crawl(db *sql.DB) bool {
	urlToCrawl := models.ChooseNextUrlToCrawl(db)
	if urlToCrawl == "" {
		fmt.Println("No URLs to crawl")
		return true
	}
	fmt.Printf("Crawling: %s\n", urlToCrawl)

	hostname := UrlToHostname(urlToCrawl)

	domain, err := models.GetDomain(db, hostname)
	if err != nil {
		if err == sql.ErrNoRows {
			domain = models.CreateDomain(db, hostname)
		} else {
			panic(err)
		}
	}

	var disallowedUrls []string
	if err != nil || time.Since(domain.AddedAt) > 7*24*time.Hour {
		fmt.Printf("Updating robots.txt for %s\n", hostname)
		disallowedUrls = FetchAndParseRobotsTxt(fmt.Sprintf("https://%s", hostname))
		models.RemoveDisallowList(db, hostname)
		for _, disallowedUrl := range disallowedUrls {
			models.AddDisallowListEntry(db, hostname, disallowedUrl)
		}
	} else {
		disallowedUrls = models.GetDisallowList(db, hostname)
	}

	if !IsUrlAllowed(urlToCrawl, disallowedUrls) {
		return false
	}

	htmlResponse, err := Fetch(urlToCrawl)
	if err != nil {
		panic(err)
	}
	defer htmlResponse.Body.Close()

	doc, err := goquery.NewDocumentFromReader(htmlResponse.Body)
	title := doc.Find("head>title").First().Text()
	description := doc.Find("head>meta[name=description]").First().AttrOr("content", "")

	if models.IfPageExists(db, urlToCrawl) {
		models.UpdatePage(db, urlToCrawl, title, description)
	} else {
		models.AddPage(db, urlToCrawl, title, description)
	}

	// remove all existing anchors
	models.RemoveAllAnchorLinksFromPage(db, urlToCrawl)
	db.Ping()

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		innerText := strings.Trim(strings.ReplaceAll(s.Text(), "\n", ""), " ")
		if innerText == "" {
			innerText = "_"
		}
		link := s.AttrOr("href", "")
		if link == "" {
			return
		}
		if strings.HasPrefix(link, "/") {
			link = fmt.Sprintf("https://%s%s", hostname, link)
		} else if strings.HasPrefix(link, "#") {
			link = urlToCrawl
		} else if strings.HasPrefix(link, ".") {
			parsedCurrentUrl, err := url.Parse(urlToCrawl)
			if err != nil {
				return
			}
			link = parsedCurrentUrl.JoinPath(link).String()
		}
		models.AddAnchorLink(db, innerText, urlToCrawl, link)
		if models.IsUrlInToCrawl(db, link) {
			models.IncreasePriority(db, link)
		} else {
			models.AddToCrawlEntry(db, link, time.Now())
		}
	})

	models.RemoveToCrawlEntry(db, urlToCrawl)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	models.AddToCrawlEntry(db, urlToCrawl, time.Now().Add(7*24*time.Hour))
	fmt.Printf("Crawled: %s\n", urlToCrawl)
	return false
}
