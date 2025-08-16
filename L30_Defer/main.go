package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Start")
	data := add(5, 11)

	defer fmt.Println("Data is :", data)
	defer fmt.Println("Middle")
	// Multiple defer statements -> executed in LIFO order

	fmt.Println("End")
}