package main

import "fmt"

func main() {
	//	La funcion opeAritmetica nos retorna una funcion(operacion) dependiento el parametro que le pasemos
	operation := opeAritmetica("Suma")

	result := operation(2, 3)
	fmt.Println(result)
}

// Orquestador
func opeAritmetica(operation string) func(val1, val2 float64) float64 {
	switch operation {
	case "Suma":
		return opSum
	case "Resta":
		return opRes
	case "Division":
		return opDivi
	case "Multiplicacion":
		return opMulti
	}
	return nil
}

// Operaciones
func opSum(val1, val2 float64) float64 {
	return val1 + val2
}
func opRes(val1, val2 float64) float64 {
	return val1 - val2
}
func opMulti(val1, val2 float64) float64 {
	return val1 * val2
}
func opDivi(val1, val2 float64) float64 {
	if val2 == 0 {
		return 0
	}
	return val1 / val2
}
