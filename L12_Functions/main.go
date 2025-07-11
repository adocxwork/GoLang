package main

import "fmt"

// func add(a, b int) int {
// func add(a int, b int) int {
// 	return a + b
// }

// fucntion can return multiple values
func getLanguages() (string, string, bool) {
	return "golang", "js", true
}

// function can take another func as parameter
func processIt(fn func(a int) int) {
	fn(1)
}

// func can return another func
func processIt2() func(a int) int {
	return func(a int) int {
		return 3
	}
}

func main() {
	l1, l2, _ := getLanguages() // use _ if u don't want to use a value
	fmt.Println(l1, l2)

	fn := func(a int) int {
		return 2
	}
	processIt(fn)
}
