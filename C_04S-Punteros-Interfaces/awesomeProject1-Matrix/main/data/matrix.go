package data

import (
	"awesomeProject1/main/utils"
	"errors"
	"fmt"
)

type Calculator interface {
	Max(numbers []int) (int, error)
}

type Matrix struct {
	values       []int
	wide, length int
	calculator   Calculator
}

func NewMatrix(wide, length int, calculator Calculator) (Matrix, error) {

	if wide == 0 && length == 0 {
		return Matrix{}, errors.New("Error: wide == 0 && length == 0 ")
	}

	//si no se recibe una implemantacion se hace una implementacion por default para no romper el programa
	if calculator == nil {
		calculator = utils.NormalMax{}
	}

	return Matrix{
		wide:       wide,
		length:     length,
		calculator: calculator,
	}, nil
}

func (m *Matrix) Set(values []int) error {
	if len(values) != m.wide*m.length {
		errors.New("matrix len is not equal with values len")
	}
	m.values = values
	return nil
}

func (m *Matrix) IsSquare() bool {
	return m.wide == m.length
}

func (m *Matrix) Print() {
	if len(m.values) == 0 {
		fmt.Println("La matriz está vacía")
	}
	for fila := 0; fila < m.wide; fila++ {
		fmt.Println(m.values[fila*m.length : fila*m.length+m.length])
	}
}

// Max - esta es la forma en la que esta Matrix utiliza la interfaz por adentro
func (m *Matrix) Max() int {
	values, _ := m.calculator.Max(m.values)
	return values
}
