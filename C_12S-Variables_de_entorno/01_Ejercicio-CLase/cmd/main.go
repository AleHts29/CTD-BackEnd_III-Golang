package main

import (
	"C_12S-practica_arquitectura/cmd/server/config"
	"C_12S-practica_arquitectura/cmd/server/external/memory"
	"C_12S-practica_arquitectura/cmd/server/handler"
	"C_12S-practica_arquitectura/cmd/server/middlewares"
	"C_12S-practica_arquitectura/internal/products"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	env := os.Getenv("ENV")
	fmt.Println("env load ----->" + env)
	if env == "" {
		env = "local"
	}
	cfg, err := config.NewConfig(env)
	if err != nil {
		panic(err)
	}

	authMidd := middlewares.NewAuth(cfg.PublicConfig.PublicKey, cfg.PrivateConfig.SecretKey)
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
	productsGroup.PUT("/:id", authMidd.AuthHeader, productsHandler.PutProduct)

	router.Run(":8080")
}
