package main

import (
	"sync"
	"testing"
)

func TestIncrementPageVisit(t *testing.T) {
	tests := []struct {
		name           string
		initialPages   map[string]int
		url            string
		expectedFirst  bool
		expectedCounts map[string]int
	}{
		{
			name:           "First visit to URL",
			initialPages:   map[string]int{},
			url:            "https://example.com",
			expectedFirst:  true,
			expectedCounts: map[string]int{"https://example.com": 1},
		},
		{
			name:           "Second visit to URL",
			initialPages:   map[string]int{"https://example.com": 1},
			url:            "https://example.com",
			expectedFirst:  false,
			expectedCounts: map[string]int{"https://example.com": 2},
		},
		{
			name:           "First visit to a different URL",
			initialPages:   map[string]int{"https://example.com": 1},
			url:            "https://golang.org",
			expectedFirst:  true,
			expectedCounts: map[string]int{"https://example.com": 1, "https://golang.org": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config{
				pages: tt.initialPages,
				mu:    &sync.RWMutex{},
				done:  make(chan struct{}),
				wg:    &sync.WaitGroup{},
			}

			gotFirst := cfg.incrementPageVisit(tt.url)

			// Check if the first visit status is as expected
			if gotFirst != tt.expectedFirst {
				t.Errorf("gotFirst = %v, expectedFirst %v", gotFirst, tt.expectedFirst)
			}

			// Check if the page counts are as expected
			for url, expectedCount := range tt.expectedCounts {
				if count, exists := cfg.pages[url]; !exists || count != expectedCount {
					t.Errorf("For URL %s, got count = %d, expected count = %d", url, count, expectedCount)
				}
			}
		})
	}
}
