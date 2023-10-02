package main

import (
	"C_09S-API_Persona_Producto/models"
	"C_09S-API_Persona_Producto/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	personaUno := models.Persona{
		Nombre:    "Ale",
		Apellido:  "Huertas",
		Edad:      30,
		Direccion: "CABA",
		Telefono:  "12123231",
		Activo:    false,
	}

	//create product list file
	products := []models.Product{
		{
			Id:         1,
			Name:       "Producto 1",
			Price:      10.99,
			Stock:      100,
			Code:       "ABC123",
			Publish:    true,
			CreateDate: time.Now(),
		},
		{
			Id:         2,
			Name:       "Producto 2",
			Price:      15.49,
			Stock:      50,
			Code:       "DEF456",
			Publish:    true,
			CreateDate: time.Now(),
		},
		// Agrega más productos aquí
	}
	utils.CreateJSON(products)

	//json.Marshal --> nos retorna una cadena de bytes
	personaByte, err := json.Marshal(personaUno)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(personaByte)
	fmt.Println(string(personaByte)) //convirtiendo a string

	//Routers
	contextGetFunc := func(context *gin.Context) {
		context.JSON(http.StatusOK, personaUno)
	}
	router.GET("/persona", contextGetFunc)

	router.GET("productos", func(context *gin.Context) {
		//Read product File
		productList := utils.ReadJSON("products.json")
		context.JSON(http.StatusOK, productList)
	})

	router.Run()

}
