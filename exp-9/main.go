package main

import "fmt"

func main() {
	fmt.Printf("%v", sum(10, 10, 10, 10, 10, 10, 10, 10, 10, 10))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
