package utils

import (
	"C_09S-API_Persona_Producto/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// CreateJSON - Crear un archivo .JSON
func CreateJSON(products []models.Product) {
	file, err := os.Create("products.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//	Crear un codificador JSON
	encoder := json.NewEncoder(file)

	// Codificar la lista de productos y escribir en el archivo
	if err := encoder.Encode(products); err != nil {
		log.Fatal(err)
	}

	fmt.Println("File JSON 'products.json' successfully created")
}

func ReadJSON(fileName string) []models.Product {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Decodificar el contenido JSON en la variable 'productos'
	var products []models.Product
	if err := decoder.Decode(&products); err != nil {
		fmt.Println("Error al decodificar el archivo JSON:", err)
		return nil
	}

	// Ejemplo: Imprimir la lista de productos
	fmt.Println("Lista de productos:")
	fmt.Println(products)
	return products
}
