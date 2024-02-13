package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	gustavo := Client{
		Name:   "Gustavo",
		Age:    25,
		Active: true,
	}

	fmt.Printf("%v", gustavo.Name)
}
