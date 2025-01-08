package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	if inputBase == outputBase {
		return inputDigits, nil
	}

	res := []int{}
	base10, err := toBase10(inputBase, inputDigits)
	if err != nil {
		return nil, err
	}
	if base10 == 0 {
		return []int{0}, nil
	}

	for base10 > 0 {
		divisor := base10 / outputBase
		remainder := base10 % outputBase
		res = append([]int{remainder}, res...)
		base10 = divisor
	}

	return res, nil
}

func toBase10(inputBase int, inputDigits []int) (int, error) {
	base10 := 0
	for i, digit := range inputDigits {
		if digit < 0 {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}
		if digit >= inputBase {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}

		base10 = base10 + (digit * int(math.Pow(float64(inputBase), float64(len(inputDigits)-i-1))))

	}

	return base10, nil
}
