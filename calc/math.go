package calc

import "errors"

// This function returns the sum of numbers along with an error
func Add(numbers ...int) (int, error) {
	if len(numbers) < 2 {
		return 0, errors.New("at least 2 arguments required")
	}

	sum := 0

	for _, num := range numbers {
		sum += num
	}
	return sum, nil
}
