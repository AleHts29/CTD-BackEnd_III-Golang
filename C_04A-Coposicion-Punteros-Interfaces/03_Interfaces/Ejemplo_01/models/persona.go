package models

import "fmt"

type Persona struct {
	ID     int64
	Nombre string
	Edad   int64
	Estado bool
}

/*
Hablar - Metodo (con receptores: por valor o por referencia(punteros))

	--> Estos metodos cumplen con la firma de la interface Humano
*/
func (p *Persona) Hablar() {
	fmt.Printf("\n%s esta hablando.. ", p.Nombre)
}

func (p *Persona) Correr() {
	fmt.Printf("\n%s esta corriendo.. ", p.Nombre)
}

func (p *Persona) Caminar() {
	fmt.Printf("\n%s esta caminando.. ", p.Nombre)
}
