package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {

	runes := []rune(strings.ToLower(s))
	builder := strings.Builder{}

	alphabet := []rune("zyxwvutsrqponmlkjihgfedcba")

	c := 0
	for _, val := range runes {
		if isDigit, isLetter := unicode.IsDigit(val), unicode.IsLetter(val); isDigit || isLetter {
			if c%5 == 0 && c != 0 {
				builder.WriteRune(' ')
				c = 0
			}
			if isDigit {
				builder.WriteRune(val)
			} else {
				builder.WriteRune(alphabet[val-'a'])
			}
			c++
		}
	}

	return builder.String()

}
