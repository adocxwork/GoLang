package main

import "fmt"

// for --> only construct for looping
func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	// for {
	// 	fmt.Println("hello")
	// }

	// classic for loop
	for k := 0; k < 2; k++ {
		// break, continue
		fmt.Println(k)
	}

	for k := range 3 {
		fmt.Println(k)
	}
}
