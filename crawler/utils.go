package crawler

import (
	"io"
	"net/http"
	"net/url"
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

func Fetch(targetUrl string) (string, error) {
	client := &http.Client{
		Timeout: 9 * time.Second,
	}

	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Spidy")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
