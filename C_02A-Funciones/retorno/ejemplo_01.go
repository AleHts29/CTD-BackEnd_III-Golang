package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := division(2, 0)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}

// Manejo de error
func division(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("el divisor no puede ser cero")
	}
	return dividendo / divisor, nil
}
