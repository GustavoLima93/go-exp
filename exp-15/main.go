package main

func main() {
	// memory -> address -> value

	a := 42

	//pointer to a
	var pointer *int = &a

	//dereferencing
	*pointer = 10

	// b is a pointer to a
	b := &a
	// dereferencing
	*b = 20

	// print the value of a and b
	println(a, *b)

}
