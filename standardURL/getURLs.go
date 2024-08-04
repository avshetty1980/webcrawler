package standardURL

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetURLsFromPage(htmlBody string, baseURL *url.URL) ([]string, error) {

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("could not parse HTML from %v, %w", baseURL, err)
	}

	var urls []string
	var traverseNodes func(*html.Node)
	traverseNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, anchor := range node.Attr {
				if anchor.Key != "href" {
					//not a link
					continue
				}
				href, err := url.Parse(anchor.Val)
				if err != nil {
					fmt.Printf("couldn't parse href '%v': %v\n", anchor.Val, err)
					continue
				}

				resolvedURL := baseURL.ResolveReference(href)
				urls = append(urls, resolvedURL.String())

			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child)
		}
	}
	traverseNodes(doc)

	return urls, nil
}