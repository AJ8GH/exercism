package thefarm

import (
	"errors"
	"fmt"
)

var errCows = errors.New("invalid number of cows")

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	ff, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	a, err := fc.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}

	return (a * ff) / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0.0, errCows
	}
	return DivideFood(fc, cows)
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %v", err.cows, err.message)
}

func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{cows, "there are no negative cows"}
	case cows == 0:
		return &InvalidCowsError{cows, "no cows don't need food"}
	default:
		return nil
	}
}
