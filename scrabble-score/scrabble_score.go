package scrabble

import "strings"

func Score(word string) int {
	sum := 0
	for _, val := range strings.ToLower(word) {
		sum += getRuneValue(val)
	}

	return sum
}

func getRuneValue(letter rune) int {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	}
	return 0
}
