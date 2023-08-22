package utils

import "errors"

type NormalMax struct{}

func (NormalMax) Max(numbers []int) (int, error) {
	if len(numbers) < 1 {
		return 0, errors.New("matrix is empty")
	}

	max := numbers[0]

	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	return max, nil
}
