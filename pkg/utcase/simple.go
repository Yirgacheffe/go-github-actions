package utcase

import (
	"errors"
)

var ErrNeedParams = errors.New("please provide more than 2 numbers")

// Add return sum of two numbers
func Add(a, b int) int {
	return a + b
}

func Sum(numbers ...int) (int, error) {
	sum := 0
	if len(numbers) < 2 {
		return sum, ErrNeedParams
	}

	for _, n := range numbers {
		sum += n
	}
	return sum, nil // no error happened, sslslsllslsls
}
