package isogram

import "strings"

func IsIsogram(word string) bool {
	letters := map[rune]bool{}
	for _, val := range strings.ToUpper(word) {
		if _, found := letters[val]; found {
			return false
		}
		if val != '-' && val != ' ' {
			letters[val] = true
		}
	}
	return true

}
