package models

import (
	"reflect"
	"testing"
)

func TestAreInternalUrls(t *testing.T) {
	tests := []struct {
		name     string
		url1     string
		url2     string
		expected bool // true = same origin, false = different origin
	}{
		// ✅ Basic Matching
		{"Exact match", "https://example.com", "https://example.com", true},
		{"Same domain with path", "https://example.com/home", "https://example.com/about", true},

		// ✅ Different Schemes
		{"HTTP vs HTTPS (different origins)", "http://example.com", "https://example.com", false},

		// ✅ Different Domains
		{"Subdomain vs Root Domain", "https://sub.example.com", "https://example.com", false},
		{"Different domains", "https://example.com", "https://other.com", false},

		// ✅ Ports Affect Origin
		{"Same domain different port", "https://example.com:8080", "https://example.com:9090", false},
		{"Default HTTP port match", "http://example.com", "http://example.com:80", true},
		{"Default HTTPS port match", "https://example.com", "https://example.com:443", true},

		// ✅ Query Params & Fragments (Ignored)
		{"Query params ignored", "https://example.com/page?user=1", "https://example.com/page?user=2", true},
		{"Fragment ignored", "https://example.com/page#section", "https://example.com/page", true},

		// ✅ Different TLDs
		{"Same name different TLD", "https://example.com", "https://example.org", false},

		// ✅ Edge Cases
		{"Same IP address", "http://192.168.1.1", "http://192.168.1.1", true},
		{"Same IP different protocol", "http://192.168.1.1", "https://192.168.1.1", false},
		{"Localhost match", "http://localhost", "http://localhost", true},
		{"Localhost different ports", "http://localhost:3000", "http://localhost:4000", false},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			result := AreInternalHrefs(testItem.url1, testItem.url2)
			if !reflect.DeepEqual(result, testItem.expected) {
				t.Fatalf("expected %v, got %v", testItem.expected, result)
			}
		})
	}
}
