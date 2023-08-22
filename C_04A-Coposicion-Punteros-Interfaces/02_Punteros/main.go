package main

import "fmt"

/*
Ejemplo de uso de estructuras, composición y métodos con receptores.
	- Vamos a poder hacer referencia a la estructura original: cualquier modificación en el objeto dentro del método, se verá reflejado en el objeto original.
*/

type Compania struct {
	Nombre string
	Puesto string
}

type Empleado struct {
	Nombre   string
	Apellido string
	Compania Compania
}

// Metodo
func (e Empleado) Informacion() {
	fmt.Printf("\nEmpleado: %s %s\nPuesto: %s\nCompañia: %s\n", e.Nombre, e.Apellido, e.Compania.Puesto, e.Compania.Nombre)
}

func (c *Compania) CambiarPuesto(nuevoPuesto string) {
	c.Puesto = nuevoPuesto
}

func main() {
	empleado := Empleado{
		Nombre:   "Juan",
		Apellido: "Lopez",
		Compania: Compania{
			Nombre: "IT Solutions",
			Puesto: "Junior Dev",
		},
	}

	empleado.Informacion()
	empleado.Compania.CambiarPuesto("Senior Dev")
	empleado.Informacion()
}
