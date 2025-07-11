package main

import (
	"fmt"
	"slices"
)

// slices - arraylist of java
// most used construct
func main() {
	// uninitialized slice is nil
	var num1[] int
	fmt.Println(num1)
	fmt.Println(len(num1))
	fmt.Println(num1 == nil)
	fmt.Println()
	
	var num2 = make([]int, 2, 5)
	// capacity -> maximum numbers of elements can fit
	fmt.Println("Length", len(num2))
	fmt.Println("Capacity", cap(num2))
	fmt.Println(num2)
	num2 = append(num2, 1)
	num2[1] = 7
	fmt.Println(num2)
	fmt.Println()

	// declare & initialize
	num3 := []int {1,2,3,4,5}
	fmt.Println(num3)
	// slice operator
	fmt.Println(num3[:2])
	fmt.Println(num3[2:])
	fmt.Println(num3[2:4])
	fmt.Println()

	var num4 = make([]int, 3)
	copy(num4, num2)
	fmt.Println("num4 & num2 :-")
	fmt.Println(num4, num2)
	fmt.Println()

	// 2D slices
	// var arr = [][]int {{1,2,3}, {4,5,6}}
	arr := [][]int {{1,2,3}, {4,5,6}}
	fmt.Println(arr)
	fmt.Println()

	ar := []int {1,2,3}
	ar2 := []int {1,2,3}
	fmt.Println(slices.Equal(ar, ar2))
}