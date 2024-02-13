package main

func main() {
	a := 10
	b := 20

	if a > b {
		println("A é maior que B")
	} else {
		println("B é maior que A")
	}

	switch a {
	case 10:
		println("A é 10")
	case 20:
		println("A é 20")
	default:
		println("A não é 10 nem 20")
	}

	if a >= 10 && b >= 10 {
		println("A e B são maiores que 10")
	}

	if a >= 10 || b >= 10 {
		println("A ou B são maiores que 10")
	}
}
