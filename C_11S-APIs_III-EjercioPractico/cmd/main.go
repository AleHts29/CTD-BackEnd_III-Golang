package main

import (
	"C_11S-practica_arquitectura/cmd/server/external/memory"
	"C_11S-practica_arquitectura/cmd/server/handler"
	"C_11S-practica_arquitectura/internal/products"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//Ruta para telemetria de la app
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})

	repository, error := memory.NewService("products.json")
	if error != nil {
		panic(error)
	}
	productsService := products.NewService(repository)

	productsHandler := handler.NewProductsHandler(productsService, productsService)

	productsGroup := router.Group("/products")
	productsGroup.GET("/:id", productsHandler.GetProductByID)
	productsGroup.PUT("/:id", productsHandler.PutProduct)

	router.Run(":8080")
}
