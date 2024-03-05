package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	stripped := []rune(strings.ReplaceAll(id, " ", ""))
	if len(stripped) <= 1 {
		return false
	}

	sum := 0

	parity := len(stripped) % 2
	for i := len(stripped) - 1; i >= 0; i-- {
		if !unicode.IsDigit(stripped[i]) {
			return false
		}

		digit, _ := strconv.Atoi(string(stripped[i]))

		if i%2 == parity {
			if doubled := digit * 2; doubled > 9 {
				sum += doubled - 9
			} else {
				sum += doubled
			}
		} else {
			sum += digit
		}

	}

	return sum%10 == 0
}
