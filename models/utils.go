package models

import (
	"crypto/sha1"
	"encoding/base64"
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
