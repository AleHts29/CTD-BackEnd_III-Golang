package main

import (
	"fmt"
)

func main() {
	n := 3

	////time.Sleep --> no es lo mas optimo
	//go multiplicarPorDos(n)
	//time.Sleep(time.Second)

	//Usando un chanel
	ch := make(chan int)
	go multiplicarPorDos(n, ch)
	fmt.Println(<-ch)
}

func multiplicarPorDos(num int, ch chan int) {
	res := num * 2
	ch <- res
}
