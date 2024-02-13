package carro

type Carro struct {
	marca string
}

func NewCarro(marca string) *Carro {
	return &Carro{marca: marca}
}

func (c *Carro) GetMarca() string {
	return c.marca
}
