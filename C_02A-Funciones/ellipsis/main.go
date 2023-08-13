package main

import "fmt"

// Calculadora simple
const (
	Suma           = "+"
	Resta          = "-"
	Multiplicaicon = "*"
	Division       = "/"
)

func main() {
	fmt.Println(operacionAritmetica(Suma, 2, 3, 4, 5, 6))
}

// func calculator
func operacionAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case Suma:
		return orquestadorOperaciones(valores, opSuma)
	case Resta:
		return orquestadorOperaciones(valores, opResta)
	case Multiplicaicon:
		return orquestadorOperaciones(valores, opMultip)
	case Division:
		return orquestadorOperaciones(valores, opDivi)
	}
	return 0
}

// func orquestador
func orquestadorOperaciones(valores []float64, operacion func(value1, value2 float64) float64) float64 {
	var result float64
	for i, valor := range valores {
		if i == 0 {
			result = valor
		} else {
			result = operacion(result, valor)
			//fmt.Println(result)
		}
	}
	return result
}

// Operadores
func opSuma(num1, num2 float64) float64 {
	return num1 + num2
}

func opResta(num1, num2 float64) float64 {
	return num1 - num2
}

func opMultip(num1, num2 float64) float64 {
	return num1 * num2
}

func opDivi(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}
