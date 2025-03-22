package crawler

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func UrlToHostname(targetUrl string) string {
	u, err := url.Parse(targetUrl)
	if err != nil {
		panic(err)
	}
	return u.Hostname()
}

func RobotsTxtUrl(targetUrl string) string {
	u, err := url.Parse(targetUrl)
	if err != nil {
		panic(err)
	}
	return u.Scheme + "://" + u.Hostname() + "/robots.txt"
}

func Fetch(targetUrl string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 9 * time.Second,
	}

	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Spidy")

	return client.Do(req)
}

func ResponseToString(resp *http.Response) (string, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func FetchAndParseRobotsTxt(hostname string) []string {
	robotsTxtUrl := RobotsTxtUrl(hostname)
	response, err := Fetch(robotsTxtUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ResponseToString(response)
	disallowedList := ParseRobotsTxt(body, "Spidy")
	return disallowedList
}

func IsUrlAllowed(targetUrl string, disallowedUrlPatterns []string) bool {
	if len(disallowedUrlPatterns) == 0 {
		return true
	}
	u, err := url.Parse(targetUrl)
	if err != nil {
		panic(err)
	}
	for _, pattern := range disallowedUrlPatterns {
		if strings.HasPrefix(u.Path, pattern) {
			return false
		}
		if strings.HasSuffix(pattern, "$") {
			frontPartMatches := strings.HasPrefix(u.Path, pattern[:len(pattern)-1])
			if frontPartMatches && len(u.Path) == len(pattern)-1 {
				return false
			}
		}
		if strings.Contains(pattern, "*") {
			patternParts := strings.Split(pattern, "*")
			if strings.HasPrefix(u.Path, patternParts[0]) && strings.HasSuffix(u.Path, patternParts[1]) {
				return false
			}
		}
	}
	return true
}
