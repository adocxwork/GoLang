package main
import "fmt"

func main() {
	var v1 int = 2
	var v2 float32 = 2.9
	var v3 string = "Hey"
	var v4 bool = false
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println()

	var a1 float32 = 6.9
	var a2 = 6.9 // datatype automatically assign ho jata hai
	a3 := 6.9 // shorthand syntax : when u need to create & initialize variable at the same time
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}
