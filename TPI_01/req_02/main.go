package main

import (
	"TPI_01/internal/tickets"
	"fmt"
)

func main() {
	fileName := "tickets.csv"
	period := "manana" // Cambia el período según tus necesidades

	count, err := tickets.GetCountByPeriod(fileName, period)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Personas que viajan en %s: %d\n", period, count)
	}
}
