package crawler

import (
	"reflect"
	"testing"
)

func TestUrlToHostname(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		// ✅ Standard Domains
		{"Simple domain", "https://example.com", "example.com"},
		{"Subdomain", "https://sub.example.com", "sub.example.com"},
		{"Multiple subdomains", "https://deep.sub.example.com", "deep.sub.example.com"},

		// ✅ Public Suffix Domains
		{"Public Suffix (co.uk)", "https://shop.example.co.uk", "shop.example.co.uk"},
		{"Public Suffix (gov.uk)", "https://service.gov.uk", "service.gov.uk"},
		{"Public Suffix (appspot.com)", "https://shop.mywebsite.appspot.com", "shop.mywebsite.appspot.com"},
		{"Public Suffix (vercel.app)", "https://shop.vercel.app", "shop.vercel.app"},

		// ✅ Edge Cases
		{"IP Address", "http://192.168.1.1", "192.168.1.1"},
		{"Localhost", "http://localhost", "localhost"},
		{"URL with port", "https://example.com:8080", "example.com"},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			result := UrlToHostname(testItem.url)
			if !reflect.DeepEqual(result, testItem.expected) {
				t.Fatalf("expected %v, got %v", testItem.expected, result)
			}
		})
	}
}
