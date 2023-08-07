package main

import "fmt"

/*
Ejercicio 1 - Listado de nombres

 1. Una profesora de la universidad quiere tener un listado con todos sus estudiantes. Es necesario crear una aplicación que contenga dicha lista. Sus estudiantes son: Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.

 2. Luego de dos clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado, sin modificar el código que escribiste inicialmente. El nombre de la nueva estudiante es Gabriela.
*/
func listadoDeNombres(arr []string) {
	// 1 - listado
	for _, name := range arr {
		fmt.Println(name)
	}

	//	2 - add to array
	//var c []string
	//c = arr
	arr = append(arr, "Gabriela")
	for _, name := range arr {
		fmt.Println(name)
	}
}

/*
Ejercicio 2 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente map, debemos imprimir la edad de Benjamin.

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado, también es necesario:
  - Saber cuántos de sus empleados son mayores de 21 años.
  - Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
  - Eliminar a Pedro del map.
*/
func search(m map[string]int) {
	count := 0

	for key, value := range m {
		if key == "Benjamin" {
			fmt.Println(key, "tiene", value, "años")
		}

		if value > 20 {
			count++
		}
	}
	fmt.Println("empleados mayores a 21 años: ", count)

	m["Federico"] = 25
	delete(m, "Pedro")
	fmt.Println("lista: ", m)

}

func main() {
	//ListNames := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	//listadoDeNombres(ListNames)

	employees := map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30,
	}
	search(employees)

}
