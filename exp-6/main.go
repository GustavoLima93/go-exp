package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	println(s)
	fmt.Printf("len: %v\ncap: %v\nvalor: %v\n", len(s), cap(s), s)
	fmt.Printf("len: %v\ncap: %v\nvalor: %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len: %v\ncap: %v\nvalor: %v\n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len: %v\ncap: %v\nvalor: %v\n", len(s[4:]), cap(s[4:]), s[4:])

	s = append(s, 10)
	fmt.Printf("len: %v\ncap: %v\nvalor: %v\n", len(s), cap(s), s)

}
