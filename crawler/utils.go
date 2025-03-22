package crawler

import (
	"net/url"
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
