package main

import (
	"fmt"
	"github.com/AleHts29/CTD-BackEnd_III-Golang/C_03S-CRUD_basic-Productos/src/models"
)

func main() {
	p := models.Product{
		ID:          4,
		Name:        "New Product",
		Price:       12.35,
		Description: "New description",
		Category:    "C",
	}
	//Dar de alta un producto
	p.Save()

	//Obtener todos los productos
	p.GetAll()

	//Obtener producto por ID
	findbyId(2)

}

func findbyId(id int) {
	searchProduct, err := models.GetById(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(searchProduct)
	}
}
