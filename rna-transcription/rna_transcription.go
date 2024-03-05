package strand

import "strings"

var proteinMap = map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

func ToRNA(dna string) string {
	builder := &strings.Builder{}

	for _, val := range dna {
		builder.WriteRune(proteinMap[val])
	}

	return builder.String()
}
