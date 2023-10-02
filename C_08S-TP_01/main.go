package main

import (
	"fmt"
	"time"

	"github.com/AleHts29/parcial_backend_III_g6/internal/utils"

	"github.com/AleHts29/parcial_backend_III_g6/internal/tickets"
)

func main() {

	//1 requerimiento
	go func(destination string) {
		totalTickets, err := tickets.GetTotalTickets(destination)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("El total de personas que viajaron a %s son: %d", destination, totalTickets)
	}("Brazil")

	//2 requerimiento
	go func(period string) {
		count, err := tickets.GetCountByPeriod(period)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("\nEl total de personas que viajaron de %s son %d", period, count)
	}("noche")

	//3 requerimiento
	go func(destination string) {
		totalTickets, err := utils.GetAllTickets()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		porcentage, err := tickets.PercentageDestination(destination, totalTickets)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("\nEl porcentaje de personas que viajaron a %s son %.1f %%", destination, porcentage)
	}("Japan")

	time.Sleep(time.Second)
}
