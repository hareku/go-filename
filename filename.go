// Package filename provides functions to escape reserved characters in Unix and Windows.
package filename

import "regexp"

// ReservedChars references https://en.wikipedia.org/wiki/Filename#Reserved_characters_and_words.
var ReservedChars = regexp.MustCompile(`[.\\/<>|:&?%*"]`)

// EscapeString replaces all reserved characters in Unix and Windows with string repl.
// The reserved characters matche filename.ReservedChars.
func EscapeString(s string, repl string) string {
	return ReservedChars.ReplaceAllString(s, repl)
}
