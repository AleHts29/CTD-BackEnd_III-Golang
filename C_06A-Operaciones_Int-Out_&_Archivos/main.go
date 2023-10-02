package main

import (
	models2 "C_06-Files/models"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	products := []models2.Product{
		{1, "Detergente", 3.99},
		{2, "Limpiavidrios", 2.49},
		{3, "Jabón en barra", 1.29},
		{4, "Esponja", 0.99},
		{5, "Desinfectante", 4.49},
	}

	//Create .csv
	file, err := os.Create("ourchasedProducts.csv")
	if err != nil {
		fmt.Println("Error to create file:", err)
		return
	}
	defer file.Close()

	//Se crea un escritor CSV
	write := csv.NewWriter(file)
	write.Comma = ';'

	/*
		Cuando se escribe en un archivo CSV en Go utilizando el escritor CSV, los datos no se escriben directamente en el archivo en cada llamada a writer.Write(). En su lugar, se almacenan en búfer en memoria para mejorar la eficiencia de escritura. writer.Flush() se utiliza para vaciar ese búfer y escribir los datos en el archivo. Sin esta llamada, los datos podrían quedar en el búfer y no se escribirían en el archivo hasta que se llame a Flush() o hasta que el búfer se llene por sí solo.
	*/
	defer write.Flush()

	// addProducts to CSV file
	for _, product := range products {
		record := []string{
			fmt.Sprintf("%d", product.ID),
			product.Name,
			fmt.Sprintf("%.2f", product.Price),
		}
		err := write.Write(record)
		if err != nil {
			fmt.Println("Error writing product data: ", err)
			return
		}
	}
}
