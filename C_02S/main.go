package main

import "fmt"

/*
Ejercicio 1 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
Categoría C: su salario es de $1.000 por hora.
Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

func main() {
	employeeName := "Camilo"
	totalsalary := calSalary(120, "A")
	fmt.Printf("nombre: %s\nTotal a pagar: $%.2f, mes actual", employeeName, totalsalary)

}

func calSalary(min float64, category string) float64 {
	oneHour := 60.0
	switch category {
	case "A":
		salaryByHour := 3000.0
		additional := 1.50
		cantHours := min / oneHour
		totalSalary := (salaryByHour * cantHours) * additional
		return totalSalary
	case "B":
		salaryByHour := 1500.0
		additional := 1.20
		cantHours := min / oneHour
		totalSalary := (salaryByHour * cantHours) * additional
		return totalSalary
	case "C":
		salaryByHour := 1000.0
		cantHours := min / oneHour
		return salaryByHour * cantHours
	}
	return 0
}
