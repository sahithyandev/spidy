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

func TestIsUrlAllowed(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		disallowedUrls []string
		expected       bool
	}{
		{"Exact match disallowed", "https://example.com/admin", []string{"/admin"}, false},
		{"No match, so allowed", "https://example.com/home", []string{"/admin"}, true},

		// ✅ Subpath Blocking
		{"Blocked subpath", "https://example.com/admin/settings", []string{"/admin"}, false},
		{"Allowed unrelated path", "https://example.com/user", []string{"/admin"}, true},

		// ✅ Wildcards (`*`)
		{"Wildcard match (should be disallowed)", "https://example.com/private/data", []string{"/private/*"}, false},
		{"Wildcard non-match (should be allowed)", "https://example.com/public/data", []string{"/private/*"}, true},

		// ✅ `$` End Matching
		{"Exact end match disallowed", "https://example.com/secret", []string{"/secret$"}, false},
		{"Not matching beyond end (allowed)", "https://example.com/secret/info", []string{"/secret$"}, true},

		// ✅ Multiple Disallowed Paths
		{"Multiple disallowed - match first", "https://example.com/admin", []string{"/admin", "/private"}, false},
		{"Multiple disallowed - match second", "https://example.com/private", []string{"/admin", "/private"}, false},
		{"Multiple disallowed - no match (allowed)", "https://example.com/home", []string{"/admin", "/private"}, true},

		// ✅ Root Path Blocking
		{"Root path blocked", "https://example.com/", []string{"/"}, false},
		{"Root path with subpath blocked", "https://example.com/home", []string{"/"}, false},

		// ✅ Complex Patterns
		{"Complex pattern match (disallowed)", "https://example.com/api/v1/user", []string{"/api/v*/user"}, false},
		{"Complex pattern non-match (allowed)", "https://example.com/api/v2/user", []string{"/api/v1/user"}, true},

		// ✅ Query Params & Fragments
		{"Ignore query params (still disallowed)", "https://example.com/admin?user=1", []string{"/admin"}, false},
		{"Ignore fragments (still disallowed)", "https://example.com/admin#top", []string{"/admin"}, false},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			result := IsUrlAllowed(testItem.url, testItem.disallowedUrls)
			if !reflect.DeepEqual(result, testItem.expected) {
				t.Fatalf("expected %v, got %v", testItem.expected, result)
			}
		})
	}

}
