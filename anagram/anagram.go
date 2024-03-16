package anagram

import (
	"slices"
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subjectUpper := strings.ToUpper(subject)
	subjectRuneList := []rune(subjectUpper)
	sort.Slice(subjectRuneList, func(i, j int) bool {
		return subjectRuneList[i] < subjectRuneList[j]
	})

	lst := []string{}
	for _, candidate := range candidates {
		candidateUpper := strings.ToUpper(candidate)
		if len(candidateUpper) != len(subjectUpper) {
			continue
		}
		if candidateUpper == subjectUpper {
			continue
		}

		candidateLst := []rune(candidateUpper)
		sort.Slice(candidateLst, func(i, j int) bool {
			return candidateLst[i] < candidateLst[j]
		})
		if slices.Equal(subjectRuneList, candidateLst) {
			lst = append(lst, candidate)
		}

	}
	return lst
}
