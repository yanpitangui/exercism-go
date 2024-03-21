// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	builder := strings.Builder{}
	re := regexp.MustCompile(`\s+`)
	segments := re.Split(s, -1)

	for i := range segments {
		for _, word := range strings.Split(segments[i], "-") {
			for _, val := range word {
				if unicode.IsLetter(val) {
					builder.WriteRune(unicode.ToUpper(val))
					break
				}
			}

		}
	}

	return builder.String()
}
