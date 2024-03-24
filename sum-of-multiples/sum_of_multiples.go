package summultiples

import "slices"

func SumMultiples(limit int, divisors ...int) int {
	multiples := []int{}

	for _, val := range divisors {
		for i := val; i < limit && i != 0; i += val {
			if slices.Index(multiples, i) < 0 {
				multiples = append(multiples, i)
			}
		}
	}

	sum := 0
	for _, val := range multiples {
		sum += val
	}

	return sum
}
