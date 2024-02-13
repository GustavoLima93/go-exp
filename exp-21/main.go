package main

import (
	"exp-modules/go/carro"
	"exp-modules/go/matematica"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	soma := matematica.Soma(1, 2)
	carro := carro.NewCarro("Fiat")
	fmt.Printf("Soma: %d\n", soma)
	fmt.Printf("Carro: %v\n", *carro)
	fmt.Printf("Carro: %v\n", carro.GetMarca())
	fmt.Printf("%v\n", uuid.New())
}
