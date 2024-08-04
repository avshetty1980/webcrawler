package client

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRetrieveHTML(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		body        string
		contentType string
		expectErr   bool
		errMsg      string
	}{
		{
			name:       "Successful HTML retrieval",
			statusCode: http.StatusOK,
			body: `
<html>
	<body>
		<a href="https://domain.example.dev">
			<span>Example.dev</span>
		</a>
	</body>
</html>
`,
			contentType: "text/html",
			expectErr:   false,
		},
		{
			name:       "Network error",
			statusCode: 0, // simulates a network error.
			expectErr:  true,
			errMsg:     "network error",
		},
		{
			name:        "Non-HTML content type",
			statusCode:  http.StatusOK,
			body:        "Not an HTML",
			contentType: "application/json",
			expectErr:   true,
			errMsg:      "returned non-HTML content-type",
		},
		{
			name:       "HTTP error status",
			statusCode: http.StatusInternalServerError,
			expectErr:  true,
			errMsg:     "status 500",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var client *httptest.Server

			if tt.statusCode == 0 {
				client = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "network error", http.StatusInternalServerError)
				}))
			} else {
				client = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", tt.contentType)
					w.WriteHeader(tt.statusCode)
					_, _ = w.Write([]byte(tt.body))
				}))
			}

			defer client.Close()

			url := client.URL
			if tt.statusCode == 0 {
				url = "http://invalid.url" // This will mock a network error
			}

			got, err := RetrieveHTML(url)
			if (err != nil) != tt.expectErr {
				t.Errorf("error = %v, wantErr %v", err, tt.expectErr)
				return
			}
			if tt.expectErr && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("error = %v, expected error message to contain %v", err, tt.errMsg)
			}
			if !tt.expectErr && got != tt.body {
				t.Errorf("got = %v, want %v", got, tt.body)
			}

		})
	}
}
