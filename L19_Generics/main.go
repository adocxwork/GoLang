package main

import "fmt"

// Generic function

// func printSlice[T any](items []T) {
// func printSlice[T interface{}](items []T) { // this is also valid
// func printSlice[T int | string](items []T) { //
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// comparable ->  it must support the == and != operators.

func printSlice[T comparable, V string](items []T, name V) { //
	for _, item := range items {
		fmt.Println(item, name)
	}
}

func main() {
	nums := []int{1, 2, 3}
	name := []string{"golang", "typescript"}
	// printSlice(nums)
	// printSlice(name)
	// printSlice([]bool {true, false}) // error, bool is not allowed (except using comparable)

	printSlice(nums, "Adi")
	printSlice(name, "Adi")
	printSlice([]bool {true, false}, "Adi") // error, bool is not allowed (except using comparable)
}
