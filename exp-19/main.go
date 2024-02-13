package main

func main() {
	var y interface{} = "Hello world"
	println(y.(string))
	res, ok := y.(int)
	println(res, ok)
	res2 := y.(int)
	println(res2)
}
