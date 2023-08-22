package main

import "fmt"

func main() {
	fmt.Println("LA FUNCION ESTA INICIANDO")
	//ejemplo01()
	//ejemplo02()
	ejemplo03()
	fmt.Println("LA FUNCION ESTA FINALIZANDO")
}

func ejemplo01() {
	animals := []string{
		"dog",
		"cat",
		"hourse",
		"bird",
	}
	fmt.Println("the animal that fly is: ", animals[5])
}

func ejemplo02() {
	//a esta funcion se le pueden pasar diferente tipos de datos como argumento
	panic("This is a msg with panic-func")
}

func ejemplo03() {
	s := &Cat{"Misho"}
	s = nil
	s.Speak()
}

type Cat struct {
	name string
}

func (c *Cat) Speak() {
	fmt.Print(c.name, "The cat make miauu..")
}
