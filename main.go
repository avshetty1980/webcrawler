package main

import (
	"fmt"
	"log"
	"os"

	"github.com/avshetty1980/webcrawler/report"
)

func main() {

	rawRootURL := "https://swapi.dev/"

	const maxWorkers = 20 //set maximum number of goroutines

	cfg, err := newCrawler(rawRootURL, maxWorkers)
	if err != nil {
		fmt.Printf("error configuring crawler %v", err)
		return
	}

	fmt.Printf("starting crawling of %v", rawRootURL)

	cfg.wg.Add(1)
	go cfg.crawl(rawRootURL)
	cfg.wg.Wait()

	file, err := os.Create("crawl-report.csv")
	if err != nil {
		log.Printf("error creating csv file %v", err)
	}
	defer file.Close()

	report.Print(cfg.pages, rawRootURL, file)
}
