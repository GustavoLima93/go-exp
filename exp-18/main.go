package main

func main() {
	var x interface{} = 10
	var y interface{} = "Hello world"
	var t interface{}

	z(x)
	z(y)
	z(t)

}

func z(t interface{}) {
	switch t.(type) {
	case int:
		println("int")
	case string:
		println("string")
	default:
		println("unknown")
	}
}
