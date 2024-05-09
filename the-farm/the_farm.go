package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function

func DivideFood(calculator FodderCalculator, cows int) (float64, error) {
	fodder, err := calculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fat, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fodder * fat / float64(cows), nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(calculator, cows)
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	} else {
		return nil
	}
}
