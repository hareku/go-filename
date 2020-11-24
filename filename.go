// Package filename provides functions to escape reserved characters in Unix and Windows.
package filename

import "regexp"

// reservedChars references https://en.wikipedia.org/wiki/Filename#Reserved_characters_and_words.
var reservedChars = regexp.MustCompile(`[.\\/<>|:&?%*"]`)

// EscapeString replaces all reserved characters in Unix and Windows with string repl.
// Tge reserved characters matches [.\/<>|:&?%*"].
func EscapeString(s string, repl string) string {
	return reservedChars.ReplaceAllString(s, repl)
}
