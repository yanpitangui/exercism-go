package pangram

import "unicode"

func IsPangram(input string) bool {

	letterMap := map[rune]bool{}

	for i := 97; i <= 122; i++ {
		letterMap[rune(i)] = false
	}

	for _, run := range input {
		if cased := unicode.ToLower(run); cased <= 122 && cased >= 97 {
			letterMap[cased] = true
		}
	}

	for _, i := range letterMap {
		if !i {
			return false
		}
	}
	return true

}
