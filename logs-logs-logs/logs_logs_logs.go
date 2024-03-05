package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	applicationMap := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for _, val := range log {
		app, found := applicationMap[val]
		if found {
			return app
		}
	}

	return "default"

}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	newString := ""
	for _, curr := range log {
		if curr == oldRune {
			newString += string(newRune)
		} else {
			newString += string(curr)
		}
	}
	return newString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {

	return utf8.RuneCountInString(log) <= limit
}
