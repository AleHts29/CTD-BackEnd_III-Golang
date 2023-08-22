package main

import (
	"fmt"
	models2 "github.com/AleHts29/CTD-BackEnd_III-Golang/C_04A-Coposicion-Punteros-Interfaces/03_Interfaces/Ejemplo_01/models"
)

func main() {
	p := models2.Persona{
		ID:     1,
		Nombre: "Ale",
		Edad:   37,
		Estado: true,
	}

	a := models2.Animal{
		ID:     1,
		Nombre: "Homero",
		Edad:   3,
		Estado: true,
	}

	Acciones(&p)
	Acciones(&a)

	ValidarPersona(p) // cumple el criterio
	ValidarPersona(a) // No cumple

}

// ValidarPersona - se recibe una interface y se realiza un casteo para validar si cumple con la firma
func ValidarPersona(dato interface{}) {
	result, ok := dato.(models2.Persona)
	if ok {
		fmt.Println("Si es una persona", result)
	} else {
		fmt.Println("No es una persona")
	}
}

// Acciones - declaramos una funcion que recibe una interface.
func Acciones(humano models2.Humano) {
	humano.Hablar()
	humano.Correr()
	humano.Caminar()
}
