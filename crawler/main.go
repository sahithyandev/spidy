package crawler

import (
	"database/sql"
	"fmt"
	"spidy/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Crawl(db *sql.DB) {
	urlToCrawl := models.ChooseNextUrlToCrawl(db)
	if urlToCrawl == "" {
		// No URL to crawl
		return
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
		disallowedUrls = FetchAndParseRobotsTxt(hostname)
		models.RemoveDisallowList(db, hostname)
		for _, disallowedUrl := range disallowedUrls {
			models.AddDisallowListEntry(db, hostname, disallowedUrl)
		}
	} else {
		disallowedUrls = models.GetDisallowList(db, hostname)
	}

	if !IsUrlAllowed(urlToCrawl, disallowedUrls) {
		return
	}

	models.RemoveToCrawlEntry(db, urlToCrawl)
}
