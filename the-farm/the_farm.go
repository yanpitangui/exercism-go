package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calculator FodderCalculator, numberOfCows int) (float64, error) {

	fodder, err := calculator.FodderAmount(numberOfCows)
	if err != nil {
		return 0, err
	}

	factor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodder * factor / float64(numberOfCows), nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, numberOfCows int) (float64, error) {

	if numberOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(calculator, numberOfCows)

}

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}

func ValidateNumberOfCows(numberOfCows int) error {

	if numberOfCows < 0 {
		return &InvalidCowsError{
			numberOfCows: numberOfCows,
			message:      "there are no negative cows",
		}
	}
	if numberOfCows == 0 {
		return &InvalidCowsError{
			numberOfCows: numberOfCows,
			message:      "no cows don't need food",
		}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
