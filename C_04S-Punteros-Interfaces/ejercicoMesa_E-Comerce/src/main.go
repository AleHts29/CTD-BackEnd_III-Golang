package main

import (
	"c_04s_e-comerce/src/utils"
	"fmt"
)

const (
	Large  = "large"
	Medium = "medium"
	Small  = "small"
)

func main() {

	product1 := utils.CrearProducto(Small, 89.56)
	fmt.Printf("Nuevo producto:  %v", product1)

}
