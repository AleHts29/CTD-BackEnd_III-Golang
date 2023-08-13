package main

import "fmt"

//Impuestos de salario
/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo.
Para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario, teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará, además, un 10 % (27 % en total).
*/

func main() {
	depositarSueldo("Juan", 40000)
}

func depositarSueldo(name string, salary float64) {
	if salary < 50000 {
		fmt.Printf("Nombre: %s \nDescuentos: No aplica\nTotal: %.2f", name, salary)
	} else {
		grossSalary := calImpuesto(salary)
		fmt.Printf("Nombre: %s\nNeto: %.2f\nImpuestos: -%.2f\nTotal: %.2f", name, salary, grossSalary, salary-grossSalary)
	}
}

// funcion que devuelva el impuesto de un salario
func calImpuesto(salary float64) float64 {
	if (salary > 50000) && (salary < 150000) {
		return salary * 0.17
	} else {
		return salary * 0.27
	}
}
