package main

// quando queremos variaveis mutaveis utilziar ponteiros, caso contrario nao é necessario

func sum(a, b *int) int {
	*a = 50
	return *a + *b

}

func main() {

	myVar := 10
	myVar2 := 20

	sum(&myVar, &myVar2)

	println(myVar)
}
