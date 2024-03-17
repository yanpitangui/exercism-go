package lsproduct

import (
	"errors"
	"strconv"
)

func getProduct(series string) (int64, error) {

	product := int64(1)

	for _, val := range series {
		n, err := strconv.Atoi(string(val))
		if err != nil {
			return 0, errors.New("invalid series")
		}
		product = product * int64(n)
	}

	return product, nil
}

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span <= 0 {
		return 0, errors.New("invalid span")
	}

	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}

	largest := int64(0)

	for i := 0; i <= len(digits)-span; i++ {
		slc := digits[i : i+span]
		product, err := getProduct(slc)
		if err != nil {
			return 0, err
		}
		if product > largest {
			largest = product
		}
	}

	return largest, nil
}
