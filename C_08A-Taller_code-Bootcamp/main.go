package main

import "fmt"

type Bootcamp struct {
	Organizacion string
	Lenguaje     string
	Alumnos      []Alumno
}

type Alumno struct {
	Nombre       string
	FechaIngreso string
}

func crearBootcamp(org, lng string) Bootcamp {
	bootcamp := Bootcamp{
		Organizacion: org,
		Lenguaje:     lng,
		Alumnos:      nil,
	}
	return bootcamp
}

func (b *Bootcamp) agregarAlumno(alumno Alumno) {
	b.Alumnos = append(b.Alumnos, alumno)
}
func main() {
	bootcamp_Go := crearBootcamp("CoffeeCoder", "Go")

	alum1 := Alumno{
		Nombre:       "Juan",
		FechaIngreso: "22/06/2023",
	}
	alum2 := Alumno{
		Nombre:       "Alejandro",
		FechaIngreso: "22/01/2023",
	}
	alum3 := Alumno{
		Nombre:       "Matias",
		FechaIngreso: "22/03/2023",
	}

	bootcamp_Go.agregarAlumno(alum1)
	bootcamp_Go.agregarAlumno(alum2)
	bootcamp_Go.agregarAlumno(alum3)

	fmt.Println(bootcamp_Go)
}
