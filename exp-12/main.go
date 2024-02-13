package main

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

	println(gustavo.City)
}
