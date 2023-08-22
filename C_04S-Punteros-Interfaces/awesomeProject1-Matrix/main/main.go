package main

import (
	"awesomeProject1/main/data"
	"fmt"
)

func main() {

	//implMax := utils.NormalMax{}
	//matrix, err := data.NewMatrix(3, 3, implMax)

	matrix, err := data.NewMatrix(3, 3, nil)
	if err != nil {
		print(err)
		return
	}

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	err = matrix.Set(data)
	if err != nil {
		print(err)
		return
	}

	matrix.Print()

	isSquare := matrix.IsSquare()
	fmt.Printf("the flag of isSquare is %v\n", isSquare)

	maxMatrix := matrix.Max()
	fmt.Printf("the max of matrix is %v\n", maxMatrix)

}
