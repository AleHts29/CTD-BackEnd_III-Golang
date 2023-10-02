package main

import "fmt"

func checkSalary(salary int) error {
	minimumIncome := 150000
	if salary < minimumIncome {
		return fmt.Errorf("Error:	el mínimo imponible es de %d y el salario ingresado es de: %d", minimumIncome, salary)
	}
	return nil
}

func main() {
	salary := 120000

	//	comprobamos si el salario es menos a 150000
	err := checkSalary(salary)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Todo ok!!")
	}
}
