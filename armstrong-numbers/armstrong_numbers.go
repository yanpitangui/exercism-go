package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	digits := len(strconv.Itoa(n))
	sum := 0
	n1 := n
	for n1 != 0 {
		sum += int(math.Pow(float64(n1%10), float64(digits)))
		n1 = n1 / 10
	}
	return n == sum

}
