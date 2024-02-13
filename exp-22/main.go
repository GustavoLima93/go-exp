package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []int{1, 2, 3, 4, 5}
	for i, v := range numeros {
		println(i, v)
	}

	pessoas := map[string]string{"nome": "João", "sobrenome": "Silva"}

	for _, v := range pessoas {
		println(v)
	}

	for {
		println("Loop infinito")
		break
	}

}
