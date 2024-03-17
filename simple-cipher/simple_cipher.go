package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

type shift struct {
	distance int
}

type vigenere struct {
	key []rune
}

func NewCaesar() Cipher {

	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance >= 26 || distance == 0 || distance <= -26 {
		return nil
	}
	return shift{distance: distance}
}

func (c shift) Encode(input string) string {

	builder := strings.Builder{}

	for _, val := range []rune(input) {
		if unicode.IsLetter(val) {
			builder.WriteRune(applyDistance(val, c.distance))
		}
	}
	return builder.String()
}

func (c shift) Decode(input string) string {
	builder := strings.Builder{}

	for _, val := range []rune(input) {
		if unicode.IsLetter(val) {
			builder.WriteRune(applyDistance(val, -c.distance))
		}
	}
	return builder.String()
}

func applyDistance(letter rune, distance int) rune {
	if l := unicode.ToLower(letter) + rune(distance); l > 'z' {
		return 'a' + l%'z' - 1
	} else if l < 'a' {
		return 'z' - ('a' - l - 1)
	} else {
		return l
	}
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}

	re := regexp.MustCompile(`^(a)+$`)
	if re.MatchString(key) {
		return nil
	}
	for _, val := range []rune(key) {
		if !unicode.IsLetter(val) || !unicode.IsLower(val) {
			return nil
		}
	}
	return vigenere{key: []rune(key)}
}

func (v vigenere) Encode(input string) string {
	builder := strings.Builder{}

	i := 0
	for _, val := range []rune(input) {
		if unicode.IsLetter(val) {
			builder.WriteRune(applyDistance(val, int(v.key[i%len(v.key)]-'a')))
			i++
		}
	}
	return builder.String()
}

func (v vigenere) Decode(input string) string {
	builder := strings.Builder{}

	i := 0
	for _, val := range []rune(input) {
		if unicode.IsLetter(val) {
			builder.WriteRune(applyDistance(val, -int(v.key[i%len(v.key)]-'a')))
			i++
		}
	}
	return builder.String()
}
