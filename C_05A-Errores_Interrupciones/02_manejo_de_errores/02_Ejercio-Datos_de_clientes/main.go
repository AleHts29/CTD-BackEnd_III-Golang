package main

import (
	"fmt"
	"os"
)

func main() {
	//it will be executed at the end of then main function
	defer fmt.Println("Ejecucion finalizada")

	//We tried to read the file "customers.txt"
	data, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o esta dañado")
	}

	//if we succeed in reading the file, we convert the data to a string and print it to the console
	fmt.Println(string(data))
}
