package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", sum(10, 10))

	valor, err := sum2(10, 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)
}

func sum(a, b int) int {
	return a + b
}

func sum2(a, b int) (int, error) {
	if a+b > 10 {
		return a + b, errors.New("A soma é maior que 10")
	}
	return a + b, nil
}
