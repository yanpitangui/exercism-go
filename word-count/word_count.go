package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	pRunes := []rune(strings.ToLower(phrase))
	freq := Frequency{}
	currWord := []rune{}
	for _, val := range pRunes {
		if unicode.IsSymbol(val) || unicode.IsSpace(val) || unicode.IsPunct(val) {
			if len(currWord) > 0 {
				if val == '\'' && unicode.IsLetter(currWord[len(currWord)-1]) {
					currWord = append(currWord, val)
				} else {
					if currWord[len(currWord)-1] == '\'' {
						currWord = currWord[:len(currWord)-1]
					}
					freq[string(currWord)] += 1
					currWord = currWord[:0]
				}

			}
		} else {
			currWord = append(currWord, val)
		}

	}
	if len(currWord) > 0 {
		if currWord[len(currWord)-1] == '\'' {
			currWord = currWord[:len(currWord)-1]
		}
		freq[string(currWord)] += 1
	}

	return freq

}
