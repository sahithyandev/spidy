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
