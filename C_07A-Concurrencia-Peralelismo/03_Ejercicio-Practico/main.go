package main

import (
	"C_07A/models"
	"fmt"
	"time"
)

func main() {
	// Ejemplo de uso
	total := 0.0

	go func() {
		products := []models.Product{
			{"Producto 1", 10.0, 2},
			{"Producto 2", 5.0, 3},
			{"Producto 3", 8.0, 1},
		}
		priceTotal := models.SumProducts(products)
		total += priceTotal
	}()

	go func() {
		services := []models.Service{
			{"Service 1", 10.0, 1800},
			{"Service 2", 5.0, 300},
			{"Service 3", 8.0, 520},
		}
		priceTotal := models.SumServices(services)
		total += priceTotal
	}()

	go func() {
		maintenances := []models.Maintenance{
			{"Maintenance 1", 10.0},
			{"Maintenance 2", 15.0},
			{"Maintenance 3", 20.0},
		}
		priceTotal := models.SumMaintenance(maintenances)
		total += priceTotal
	}()
	time.Sleep(time.Second)
	fmt.Printf("Precio Total a pagar: %.2f\n", total)

}
