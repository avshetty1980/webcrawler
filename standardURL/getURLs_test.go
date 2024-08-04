package standardURL

import (
	"net/url"
	"testing"
)

func TestGetURLsFromPage(t *testing.T) {
	tests := []struct {
		name      string
		htmlBody  string
		baseURL   string
		wantURLs  []string
		expectErr bool
	}{
		{
			name:      "Simple HTML with links",
			htmlBody:  `<html><body><a href="https://example.com">Example</a><a href="/relative">Relative</a></body></html>`,
			baseURL:   "https://base.com",
			wantURLs:  []string{"https://example.com", "https://base.com/relative"},
			expectErr: false,
		},
		{
			name:      "Invalid HTML",
			htmlBody:  `<html><body><a href="https://example.com">Example<a href="/relative">Relative</body></html>`,
			baseURL:   "https://base.com",
			wantURLs:  []string{"https://example.com", "https://base.com/relative"},
			expectErr: false,
		},
		{
			name:      "HTML with no links",
			htmlBody:  `<html><body><p>No links here!</p></body></html>`,
			baseURL:   "https://base.com",
			wantURLs:  []string{},
			expectErr: false,
		},
		{
			name:      "Malformed URL",
			htmlBody:  `<html><body><a href="::invalid-url::">Invalid URL</a></body></html>`,
			baseURL:   "https://base.com",
			wantURLs:  []string{},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base, err := url.Parse(tt.baseURL)
			if err != nil {
				t.Fatalf("Failed to parse base URL: %v", err)
			}

			gotURLs, err := GetURLsFromPage(tt.htmlBody, base)
			if (err != nil) != tt.expectErr {
				t.Errorf("GetURLsFromPage() error = %v, expectErr %v", err, tt.expectErr)
				return
			}

			if len(gotURLs) != len(tt.wantURLs) {
				t.Errorf("GetURLsFromPage() got %v URLs, want %v URLs", len(gotURLs), len(tt.wantURLs))
			}

			for i, gotURL := range gotURLs {
				if gotURL != tt.wantURLs[i] {
					t.Errorf("GetURLsFromPage() got URL %v, want URL %v", gotURL, tt.wantURLs[i])
				}
			}
		})
	}
}
