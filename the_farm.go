package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

var nonScaleError = errors.New("non-scale error")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	//panic("Please implement DivideFood")
	if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}

	fa, err := weightFodder.FodderAmount()

	if err != nil {
		if err == ErrScaleMalfunction {
			if fa > 0 {
				return (fa * 2) / float64(cows), nil
			}
			return 0, errors.New("negative fodder")
		} else if err == nonScaleError && fa < 0 {
			return 0, errors.New("negative fodder")
		} else {
			return 0, err
		}
	}
	if fa < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	return fa / float64(cows), err
}
