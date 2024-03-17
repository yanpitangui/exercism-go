package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid n")
	}

	primes := []int{2, 3, 5, 7, 11, 13}

	c := 6
	for next := primes[c-1] + 2; c < n; next += 2 {
		if isPrime(next, primes) {
			primes = append(primes, next)
			c++
		}
	}

	return primes[n-1], nil

}

func isPrime(n int, primes []int) bool {
	m := int(math.Ceil(math.Sqrt(float64(n))))

	for _, val := range primes {
		if val > m {
			break
		}
		if n%val == 0 {
			return false
		}
	}
	return true
}
