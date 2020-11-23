package filename_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hareku/filename"
)

func TestEscapeString(t *testing.T) {
	tests := []struct {
		name      string
		unescaped string
		escaped   string
	}{
		{
			name:      "alreadyEscaped",
			unescaped: "abc",
			escaped:   "abc",
		},
		{
			name:      "dots",
			unescaped: "abc...",
			escaped:   "abc---",
		},
		{
			name:      "pipes",
			unescaped: "a|b|c",
			escaped:   "a-b-c",
		},
		{
			name:      "doubleQuotes",
			unescaped: "\"abc\"",
			escaped:   "-abc-",
		},
		{
			name:      "asteriesks",
			unescaped: "a*b*c",
			escaped:   "a-b-c",
		},
		{
			name:      "persents",
			unescaped: "a%b%c",
			escaped:   "a-b-c",
		},
		{
			name:      "URL",
			unescaped: "https://example.com/path?a=b&c=d",
			escaped:   "https---example-com-path-a=b-c=d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filename.EscapeString(tt.unescaped, "-"); got != tt.escaped {
				t.Errorf("EscapeString(%v) = %v, want %v", tt.unescaped, got, tt.escaped)
			}
		})
	}
}

func ExampleEscapeString() {
	// Use URL as a file name.
	fmt.Println(filepath.Join("/tmp", filename.EscapeString("https://example.com", "-")))
	// Output: /tmp/https---example-com
}
