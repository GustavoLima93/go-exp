package main

import "fmt"

type Address struct {
	City    string
	Country string
	Street  string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c Client) GetAddress() string {
	return c.City + ", " + c.Country
}

func main() {
	gustavo := Client{
		Name:   "Gustavo",
		Age:    25,
		Active: true,
		Address: Address{
			City:    "São Paulo",
			Country: "Brazil",
			Street:  "Rua das Flores",
		},
	}

	fmt.Printf(gustavo.GetAddress())
}
