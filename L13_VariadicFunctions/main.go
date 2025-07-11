package main

import "fmt"

// func sum(nums ... interface{}) int { // for any data type...
func sum(nums ... int) int {
	s := 0
	for _, i := range nums { // nums -> slice
		s += i
	}
	return s
}

func main() {
	result := sum(1, 3, 4) // any no of parameters
	fmt.Println(result)

	result2 := []int {4,3,5}
	fmt.Println(sum(result2...)) // ... se slice ko khol diya
}