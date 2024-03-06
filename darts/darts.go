package darts

import "math"

func distance(x, y float64) float64 {
	return math.Sqrt(math.Abs(x*x) + math.Abs(y*y))
}

func Score(x, y float64) int {

	d := distance(x, y)
	switch {
	case d <= 1:
		return 10
	case d <= 5:
		return 5
	case d <= 10:
		return 1
	default:
		return 0
	}
}
