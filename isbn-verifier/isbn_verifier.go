package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	trimmed := []rune(strings.ReplaceAll(isbn, "-", ""))

	if len(trimmed) != 10 {
		return false
	}

	sum := 0
	for i, j := 0, 10; i <= 9; i++ {
		var val int
		if unicode.IsDigit(trimmed[i]) {
			val, _ = strconv.Atoi(string(trimmed[i]))
		} else if trimmed[i] == 'X' && i == 9 {
			val = 10
		} else {
			return false
		}
		sum += val * j
		j--

	}

	return sum%11 == 0
}
