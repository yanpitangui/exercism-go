package sieve

import "math"

func Sieve(limit int) []int {

	lst := []bool{false, false}

	for i := 2; i <= limit; i++ {
		lst = append(lst, true)
	}

	primes := []int{}
	for i := 2; i <= int(math.Sqrt(float64(limit))); i++ {
		if lst[i] {
			for j := i * i; j <= limit; j += i {
				lst[j] = false
			}
		}
	}

	for i := range lst {
		if lst[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
