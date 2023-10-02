package routes

import (
	"C_10A_03/db"
	"C_10A_03/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MainPage(ctxt *gin.Context) {
	ctxt.String(200, "Bienvenido a CoffeeCoder")
}

func Employees(ctxt *gin.Context) {
	employees := db.EmpleadosDB
	ctxt.JSON(200, employees)
}

func EmployeesById(ctxt *gin.Context) {
	idStr := ctxt.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctxt.JSON(http.StatusBadRequest, gin.H{"error": "ID no valido"})
		return
	}
	//Buscar el empleado por ID
	var employee models.Empleado
	encontrado := false
	for _, e := range db.EmpleadosDB {
		if e.Id == id {
			employee = e
			encontrado = true
			break
		}
	}
	if encontrado {
		ctxt.JSON(http.StatusOK, employee)
	} else {
		ctxt.JSON(http.StatusNotFound, gin.H{"error": "Empleado no encontrado"})
	}
}

func Employeesparams(ctxt *gin.Context) {
	// Obtener los parámetros de la consulta
	idStr := ctxt.Query("id")
	nombre := ctxt.Query("nombre")
	activoStr := ctxt.Query("activo")

	// Convertir los parámetros en sus tipos correspondientes
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctxt.JSON(http.StatusBadRequest, gin.H{"error": "ID no valido"})
		return
	}

	activo, err := strconv.ParseBool(activoStr)
	if err != nil {
		ctxt.JSON(http.StatusBadRequest, gin.H{"error": "Valor 'activo' no válido"})
		return

	}

	//Crear un nuevo empleado
	nuevoEmpleado := models.Empleado{
		Id:     id,
		Nombre: nombre,
		Activo: activo,
	}

	//Agregar el nuevo empleado a la lista
	db.EmpleadosDB = append(db.EmpleadosDB, nuevoEmpleado)

	// Devolver el nuevo empleado en formato JSON
	ctxt.JSON(http.StatusCreated, nuevoEmpleado)
}

func Employeesactive(ctxt *gin.Context) {
	// Obtener el parámetro de consulta "activo"
	activoStr := ctxt.DefaultQuery("activo", "true")

	// Convertir el parámetro en un valor booleano
	activo, err := strconv.ParseBool(activoStr)
	if err != nil {
		ctxt.JSON(http.StatusBadRequest, gin.H{"error": "Valor 'activo' no válido"})
		return
	}

	// Filtrar empleados por estado activo o inactivo
	empleadosFiltrados := []models.Empleado{}
	for _, e := range db.EmpleadosDB {
		if e.Activo == activo {
			empleadosFiltrados = append(empleadosFiltrados, e)
		}
	}
	ctxt.JSON(http.StatusOK, empleadosFiltrados)
}
