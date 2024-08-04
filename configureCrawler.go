package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages   map[string]int
	baseURL *url.URL
	mu      *sync.RWMutex
	done    chan struct{}
	wg      *sync.WaitGroup
}

func (cfg *config) incrementPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func newCrawler(rawBaseURL string, maxWorkers int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:   make(map[string]int),
		baseURL: baseURL,
		mu:      &sync.RWMutex{},
		done:    make(chan struct{}, maxWorkers),
		wg:      &sync.WaitGroup{},
	}, nil
}
