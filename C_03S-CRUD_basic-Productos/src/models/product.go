package models

import (
	"errors"
	"fmt"
)

var Products = []Product{
	{ID: 1, Name: "Producto 1", Price: 19.99, Description: "Description 1", Category: "A"},
	{ID: 2, Name: "Producto 2", Price: 9.99, Description: "Description 2", Category: "A"},
	{ID: 3, Name: "Producto 3", Price: 10.05, Description: "Description 3", Category: "B"},
}

type Product struct {
	ID          int     `json:"ID"`
	Name        string  `json:"Name"`
	Price       float64 `json:"Price"`
	Description string  `json:"Description"`
	Category    string  `json:"Category"`
}

// Metodos
func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p *Product) GetAll() {
	for _, product := range Products {
		fmt.Println(product)
	}
}

// Función para obtener un producto por su ID
func GetById(id int) (Product, error) {
	for _, product := range Products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, errors.New("producto no encontrado")
}
