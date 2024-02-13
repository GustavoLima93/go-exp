package main

import "fmt"

func main() {
	total := func() int {
		return sum(10, 10, 10, 10, 10, 10, 10, 10, 10, 10)
	}()

	fmt.Printf("%v", total)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
