package db

import "C_10A_03/models"

var EmpleadosDB []models.Empleado

func CargarEmpleados() {
	//	Simulamos la carga de empleados desde una fuente de datos
	EmpleadosDB = []models.Empleado{
		{1, "Juan", true},
		{2, "Ana", true},
		{3, "Pedro", false},
	}
}
