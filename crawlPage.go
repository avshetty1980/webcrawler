package main

import (
	"log"
	"net/url"

	"github.com/avshetty1980/webcrawler/client"
	"github.com/avshetty1980/webcrawler/standardURL"
)

func (cfg *config) crawl(rawCurrentURL string) {

	cfg.done <- struct{}{}
	defer func() {
		<-cfg.done
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("Error - crawlPage: could not parse current URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	visitedOnce := cfg.incrementPageVisit(rawCurrentURL)
	if !visitedOnce {
		return
	}

	crawledHtmlBody, err := client.RetrieveHTML(rawCurrentURL)
	if err != nil {
		log.Printf("Error - retrieveHTML: %v", err)
		return
	}

	urlList, err := standardURL.GetURLsFromPage(crawledHtmlBody, cfg.baseURL)
	if err != nil {
		log.Printf("Error - GetURLsFromPage: %v", err)
		return
	}

	for _, crawledRawURL := range urlList {
		cfg.wg.Add(1)
		go cfg.crawl(crawledRawURL)
	}
}
