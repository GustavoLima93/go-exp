package main

import "fmt"

func main() {
	salario := map[string]float64{"Jose": 1000.0, "Maria": 2000.0}

	fmt.Printf("%T\n", salario)
	fmt.Println("Salario do Jose:", salario["Jose"])
	fmt.Println("Salario do Maria:", salario["Maria"])

	sal := make(map[string]float64)

	sal["Jose"] = 1000.0
	sal["Maria"] = 2000.0

	fmt.Println("Salario do Jose:", sal["Jose"])

	delete(salario, "Jose")

	fmt.Println("Salario do Jose:", salario["Jose"])

	salario["Gustavo"] = 5000.0
	fmt.Println("Salario do Gustavo:", salario["Gustavo"])

	for k, v := range salario {
		fmt.Println(k, v)
	}
}
