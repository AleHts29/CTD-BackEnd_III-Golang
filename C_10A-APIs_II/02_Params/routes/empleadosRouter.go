package routes

import (
	"C_10A/db"
	"github.com/gin-gonic/gin"
)

func PaginaPrincipal(ctxt *gin.Context) {
	ctxt.String(200, "Bienvenido a CoffeeCoder")
}

func BuscarEmpleado(ctxt *gin.Context) {
	empleado, ok := db.Empleados[ctxt.Param("id")]

	if ok {
		ctxt.String(200, "Informacion del empleado %s, nombre: %s", ctxt.Param("id"), empleado)
	} else {
		ctxt.String(400, "Informacion del empleado no existe!!")
	}
}
