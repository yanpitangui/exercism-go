package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	builder := strings.Builder{}
	runeString := []rune(plain)

	for _, val := range runeString {
		if unicode.IsLetter(val) {
			shifted := rune(int(val) + shiftKey%26)
			if val >= 'a' && val <= 'z' {
				if shifted > 'z' {
					builder.WriteRune(shifted - 26)
				} else if shifted < 'a' {
					builder.WriteRune(shifted + 26)
				} else {
					builder.WriteRune(shifted)
				}
			} else {
				if shifted > 'Z' {
					builder.WriteRune(shifted - 26)
				} else if shifted < 'A' {
					builder.WriteRune(shifted + 26)
				} else {
					builder.WriteRune(shifted)
				}
			}
		} else {
			builder.WriteRune(val)
		}
	}

	return builder.String()
}
