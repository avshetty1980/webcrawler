package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock functions
type MockClient struct{}

func (m MockClient) RetrieveHTML(rawURL string) (string, error) {
	// Mocked HTML
	return "<html></html>", nil
}

type MockStandardURL struct{}

func (m MockStandardURL) GetURLsFromPage(body string, baseURL *url.URL) ([]string, error) {
	// Mocked URL extraction
	return []string{"http://example.com/page1", "http://example.com/page2"}, nil
}

func TestCrawl(t *testing.T) {
	tests := []struct {
		name          string
		rawCurrentURL string
		setupConfig   func() *config
		expectedPages map[string]int
	}{
		{
			name:          "Invalid URL",
			rawCurrentURL: "invalid-url",
			setupConfig: func() *config {
				cfg, _ := newCrawler("http://example.com", 1)
				return cfg
			},
			expectedPages: map[string]int{},
		},
		{
			name:          "Different Hostname",
			rawCurrentURL: "http://other.com",
			setupConfig: func() *config {
				cfg, _ := newCrawler("http://example.com", 1)
				return cfg
			},
			expectedPages: map[string]int{},
		},
		{
			name:          "Already Visited URL",
			rawCurrentURL: "http://example.com",
			setupConfig: func() *config {
				cfg, _ := newCrawler("http://example.com", 1)
				cfg.pages["http://example.com"] = 1
				return cfg
			},
			expectedPages: map[string]int{"http://example.com": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setupConfig()
			cfg.wg.Add(1)
			go cfg.crawl(tt.rawCurrentURL)
			cfg.wg.Wait()

			assert.Equal(t, tt.expectedPages, cfg.pages)
		})
	}
}

func BenchmarkCrawl(b *testing.B) {
	rawBaseURL := "http://example.com"
	maxWorkers := 10

	cfg, _ := newCrawler(rawBaseURL, maxWorkers)

	for i := 0; i < b.N; i++ {
		cfg.pages = make(map[string]int) // Reset pages for each iteration
		cfg.wg.Add(1)
		go cfg.crawl(rawBaseURL)
		cfg.wg.Wait()
	}
}
