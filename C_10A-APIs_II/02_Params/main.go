package main

import (
	"C_10A/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", routes.PaginaPrincipal)
	server.GET("/empleados/:id", routes.BuscarEmpleado)

	server.Run(":9090")
}
