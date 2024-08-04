package report

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

func Print(pages map[string]int, baseURL string, w io.Writer) {

	writer := csv.NewWriter(w)
	defer writer.Flush()

	if err := writer.Write([]string{
		"===== REPORT for " + baseURL + "======"}); err != nil {
		log.Printf("could not write header: %v", err)
	}

	for URL, count := range pages {
		record := []string{"Found " + strconv.Itoa(count) + " internal links to " + URL}
		if err := writer.Write(record); err != nil {
			log.Printf("could not write record: %v", err)
		}
	}
}
