package main

import "fmt"

// type stack struct {
// 	elements []int
// }
type stack[T any] struct {
	elements []T
}

func main() {
	myStack := stack[string] {
		elements: []string {"golang", "typescript"},
	}
	fmt.Println(myStack)
}