package main

import "fmt"

func main() {
	p := Persona{
		Nombre: "Gabriel",
		Edad:   24,
	}
	p.Descripcion()
}

// definimos la estructura
type Persona struct {
	Nombre string
	Edad   int
}

// Metodo
func (p *Persona) Descripcion() {
	fmt.Printf("%s tiene %d años de edad", p.Nombre, p.Edad)

}
