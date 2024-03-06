// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func allCapital(letters []rune) bool {
	hasLetters := false
	for _, val := range letters {
		if unicode.IsLetter(val) {
			hasLetters = true
			if !unicode.IsUpper(val) {
				return false
			}
		}
	}
	return hasLetters

}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	trimmed := []rune(strings.TrimSpace(remark))

	isCapital := allCapital(trimmed)

	if len(trimmed) == 0 {
		return "Fine. Be that way!"
	}
	if last := trimmed[len(trimmed)-1]; last == '?' {
		if isCapital {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if isCapital {
		return "Whoa, chill out!"
	}

	return "Whatever."

}
