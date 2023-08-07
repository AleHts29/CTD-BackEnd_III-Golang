package main

import "fmt"

/*
Ejercicio 1 - descuento
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos. Para ello necesitan una aplicación que les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos.
*/

func Descuento(price, discount float64) float64 {
	porcentaje := discount / 100
	return price - (price * porcentaje)
}

/*
Ejercicio 2 - Préstamo
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar: solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga, no les cobrará interés a los que su sueldo sea mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes.
*/

func Prestamo(age int, employed bool, seniority int, salary float64) string {
	if age > 22 && employed == true && seniority > 1 {
		if salary > 100000 {
			return "Felicitaciones, el banco le otorga un prestamo a tasa 0."
		}
		return "Felicitaciones, el banco le otorga un prestamo pero se te cobraran intereses."
	}
	return "Lo siento, no cumple con los requisitos minimos establecidos por el banco."
}

func main() {
	//Ejercicio 1 - descuento
	price := 2300.00
	discount := 30.0
	result := Descuento(price, discount)
	fmt.Printf("Valor del producto: %.2f\nDescuento del %.1f %% \nTotal a pagar: %.2f \n\n", price, discount, result)

	//Ejercicio 2 - Préstamo
	age := 25
	employed := true
	seniority := 4
	salary := 50000.00
	result2 := Prestamo(age, employed, seniority, salary)
	fmt.Println(result2)
}
