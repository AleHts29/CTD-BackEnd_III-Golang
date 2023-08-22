package models

import "fmt"

type Animal struct {
	ID     int64
	Nombre string
	Edad   int64
	Estado bool
}

/*
Hablar - Metodo (con receptores: por valor o por referencia(punteros))

	--> Estos metodos cumplen con la firma de la interface Humano
*/
func (p *Animal) Hablar() {
	fmt.Printf("\n%s esta hablando.. ", p.Nombre)
}

func (p *Animal) Correr() {
	fmt.Printf("\n%s esta corriendo.. ", p.Nombre)
}

func (p *Animal) Caminar() {
	fmt.Printf("\n%s esta caminando.. ", p.Nombre)
}
