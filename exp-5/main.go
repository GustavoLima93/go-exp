package main

import "fmt"

func main() {
	var arr [5]int

	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	fmt.Printf("%T\n", arr)
	fmt.Println(arr)
	fmt.Printf("%v\n", arr[0])
	fmt.Printf("%v\n", arr[len(arr)-1])

	for i, v := range arr {
		fmt.Println(i, v)
		fmt.Printf("O valor do meu i é: %v\n e o valor é %v\n", i, v)
	}
}
