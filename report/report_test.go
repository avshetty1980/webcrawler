package report

import (
	"bytes"
	"testing"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name     string
		pages    map[string]int
		baseURL  string
		expected string
	}{
		{
			name:    "Valid URL",
			pages:   map[string]int{"https://example.com": 5},
			baseURL: "https://example.com",
			expected: `===== REPORT for https://example.com======
Found 5 internal links to https://example.com
`,
		},
		{
			name:    "Empty Page",
			pages:   map[string]int{},
			baseURL: "https://example.com",
			expected: `===== REPORT for https://example.com======
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var buf bytes.Buffer

			Print(tt.pages, tt.baseURL, &buf)

			got := buf.String()
			if got != tt.expected {
				t.Errorf("Print() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
