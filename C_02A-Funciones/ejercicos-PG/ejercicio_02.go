package main

import "fmt"

//Calcular promedio
/*
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.
*/

func main() {
	studentCalifications("Marcos", 5, 6.3, 7.5)
}

func studentCalifications(name string, califications ...float64) {
	averageCalifications := average(califications...)
	fmt.Printf("Nombre: %s\nPromedio de calificaciones: %.2f\n", name, averageCalifications)
}

func average(values ...float64) float64 {
	var result float64
	for _, value := range values {
		result = result + value
	}
	return result / float64(len(values))
}
