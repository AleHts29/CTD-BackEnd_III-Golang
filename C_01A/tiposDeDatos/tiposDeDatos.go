package main

import "fmt"

func main() {

	//	==== Arrays ====  //
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println("Array: ", a[0])
	fmt.Println("Array: ", a)

	//	==== Slice ====	//
	// 	Slice --> similar al array, solo que no se define el numero de elementos
	var s = []bool{true, false}
	fmt.Println("Slice: ", s[0])

	// 	crear un slice con make()
	b := make([]int, 5) // len (a) = 5
	fmt.Println(b[1])

	//	Obtener rango
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4]) // si no definimos un valor despues del ":" toma hasta el fin de elementos y viceversa

	//	Agregar a un slice data
	c := []int{32, 1}
	c = append(c, 2, 3, 4)
	fmt.Println(c)

	//	==== Mapas ====	//
	//	declaraciones de un map
	myMap1 := map[string]int{"Benjamin": 20, "Nahuel": 27}
	fmt.Println(len(myMap1)) // con len podemos determinar la longitud del map

	myMap2 := make(map[string]string) // make no me permite cargar datos en la misma sentencia
	//Agregar elementos
	myMap2["Brenda"] = "22"
	myMap2["Maria"] = "33"
	fmt.Println(myMap2)

	//	Actualizar valores
	myMap1["Benjamin"] = 33

	//Eliminar valores
	delete(myMap1, "Nahuel")

	fmt.Println(myMap1)

}
