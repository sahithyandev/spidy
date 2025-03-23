package models

import (
	"crypto/sha1"
	"encoding/base64"
	"net/url"
)

// HashUrl generates a SHA-1 hash from the url
func HashUrl(targetUrl string) string {
	hasher := sha1.New()
	hasher.Write([]byte(targetUrl))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func HashAnchorLink(text string, fromPageId string, toPageId string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	hasher.Write([]byte(fromPageId))
	hasher.Write([]byte(toPageId))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func ResolveDefaultPortForScheme(scheme string) string {
	if scheme == "http" {
		return "80"
	} else if scheme == "https" {
		return "443"
	}
	return ""
}

// AreInternalHrefs returns if the to link is internal to the origin of the first
func AreInternalHrefs(from string, to string) bool {
	fromUrl, err := url.Parse(from)
	if err != nil {
		panic(err)
	}
	toUrl, err := url.Parse(to)
	if err != nil {
		panic(err)
	}
	if fromUrl.Scheme != toUrl.Scheme || fromUrl.Hostname() != toUrl.Hostname() {
		return false
	}

	fromPort := fromUrl.Port()
	toPort := toUrl.Port()

	if fromPort == "" {
		fromPort = ResolveDefaultPortForScheme(fromUrl.Scheme)
	}
	if toPort == "" {
		toPort = ResolveDefaultPortForScheme(toUrl.Scheme)
	}
	if fromPort != toPort {
		return false
	}

	return true
}
