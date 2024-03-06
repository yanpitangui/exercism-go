package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	output := map[string]int{}
	for i, val := range in {
		for j := range val {
			output[strings.ToLower(val[j])] = i
		}
	}

	return output
}
