package mathutil

import (
	"errors"
)

// Sum returns sum of nums
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Div returns a / b
func Div(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("Can't devide by zero")
	}

	return a / b, nil
}

// Fibo returns fibonacci number of a sequence
func Fibo(a int) int {
	if a <= 1 {
		return a
	}

	return Fibo(a-1) + Fibo(a-2)
}
