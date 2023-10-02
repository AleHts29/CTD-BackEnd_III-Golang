package main

import (
	"C_10A_03/db"
	"C_10A_03/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//	Cargar el slice de empleados al iniciar la API
	db.CargarEmpleados()

	//	Definir rutas de la API
	server.GET("/", routes.MainPage)
	gopher := server.Group("/employees")
	{
		gopher.GET("/", routes.Employees)
		gopher.GET("/:id", routes.EmployeesById)
		gopher.GET("/employeesparams", routes.Employeesparams)
		gopher.GET("/employeesactive", routes.Employeesactive)
	}

	server.Run(":9090")
}
