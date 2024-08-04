package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
	},
}

func RetrieveHTML(URL string) (string, error) {

	resp, err := httpClient.Get(URL)
	if err != nil {
		return "", fmt.Errorf("network error: %v - %w", URL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("status %v returned from URL: %v", resp.StatusCode, URL)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("returned non-HTML content-type from URL %v : %v", URL, contentType)
	}

	bodyInBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("could not read response body from URL %v : %v", URL, err)
	}
	return string(bodyInBytes), nil
}
