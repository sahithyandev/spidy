package crawler

import (
	"reflect"
	"testing"
)

func TestParseRobotsTxt(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		expected []string
	}{
		{"Empty", "", []string{}},
		{"No matching entries", "User-agent: Unknown\nDisallow: *", []string{}},
		{"Exact match for User-Agent", "User-agent: Spidy\nDisallow: *", []string{"*"}},
		{"Case insensitive User-agent", "User-agent: spidy\nDisallow: /private", []string{"/private"}},
		{"Multiple Disallow rules", "User-agent: Spidy\nDisallow: /private\nDisallow: /secret", []string{"/private", "/secret"}},
		{"Wildcard Disallow", "User-agent: Spidy\nDisallow: /images/*", []string{"/images/*"}},
		{"Global rule", "User-agent: *\nDisallow: /global/", []string{"/global/"}},
		{"Specific vs Global", "User-agent: Spidy\nDisallow: /spidy/\nUser-agent: *\nDisallow: /all/", []string{"/spidy/"}},
		{"No Disallow (Allow All)", "User-agent: Spidy", []string{}},
		{"Ignore comments", "User-agent: Spidy\n# This is a comment\nDisallow: /hidden/", []string{"/hidden/"}},
		{"Sitemap directive", "User-agent: Spidy\nDisallow: /sensitive/\nSitemap: https://example.com/sitemap.xml", []string{"/sensitive/"}},
	}
	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			result := ParseRobotsTxt(testItem.body, "Spidy")
			if !reflect.DeepEqual(result, testItem.expected) {
				t.Fatalf("expected %v, got %v", testItem.expected, result)
			}
		})
	}
}
