package filename

import "regexp"

// reservedChars references https://en.wikipedia.org/wiki/Filename#Reserved_characters_and_words.
var reservedChars = regexp.MustCompile(`[./<>|:&?%*"]`)

// EscapeString escapes reserved characters in Unix and Windows.
func EscapeString(s string) string {
	return reservedChars.ReplaceAllString(s, "-")
}
