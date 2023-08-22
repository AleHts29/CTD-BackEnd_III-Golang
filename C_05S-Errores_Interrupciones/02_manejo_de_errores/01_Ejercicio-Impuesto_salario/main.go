package main

import (
	"fmt"
)

// Definimos una estructura para nuestro error personalizado
type InsufficientSalaryError struct {
}

// Implementamos el método Error() para que nuestra estructura cumpla con la interfaz error
func (e *InsufficientSalaryError) Error() string {
	return "Error: The salary entered does not reach the taxeble minimum.."
}

func main() {
	salary := 120000

	//	comprobamos si el salario es menos a 150000
	if salary < 150000 {
		//Lanzamos el error personalizado
		err := InsufficientSalaryError{}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}
