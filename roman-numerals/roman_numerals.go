package romannumerals

import (
	"errors"
	"strings"
)

var numbers = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var numberToRoman = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func ToRomanNumeral(input int) (string, error) {

	builder := strings.Builder{}
	if input <= 0 || input > 3999 {
		return builder.String(), errors.New("invalid input")
	}

	for _, val := range numbers {
		for input >= val && input >= 0 {
			input -= val
			builder.WriteString(numberToRoman[val])
		}
	}

	return builder.String(), nil
}
