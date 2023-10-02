package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	parChan := make(chan int)
	impChan := make(chan int)

	//	La espera de grupo (sync.WaitGroup) se utiliza para asegurarse de que ambas goroutines hayan terminado antes de que el programa finalice.
	var wg sync.WaitGroup

	wg.Add(1)
	go par(parChan, &wg)

	wg.Add(1)
	go impar(impChan, &wg)

	for _, num := range numbers {
		if num%2 == 0 {
			parChan <- num
		} else {
			impChan <- num
		}
	}

	close(parChan)
	close(impChan)

	wg.Wait()
}

func par(parChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Números pares:")
	for num := range parChan {
		fmt.Println(num)
	}
}

func impar(impChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Números impares:")
	for num := range impChan {
		fmt.Println(num)
	}
}
