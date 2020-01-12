package slowmath

import (
	"fmt"
)

// Abs return the absolute value of val
func Abs(val float64) float64 {
	if val < 0 {
		return -val
	}
	return val
}

// Sqrt calculate the square root of val using Newton's method
func Sqrt(val float64) (float64, error) {
	if val < 0.0 {
		return 0.0, fmt.Errorf("sqrt of negative number")
	}

	if val == 0.0 {
		return 0.0, nil // shortcut
	}

	guess, epsilon := 1.0, 0.00001
	for i := 0; i < 10000; i++ {
		if Abs(guess*guess-val) <= epsilon {
			return guess, nil
		}
		guess = (val/guess + guess) / 2.0
	}

	return 0.0, fmt.Errorf("can't find sqrt of %f", val)
}
