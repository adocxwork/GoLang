package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You are adult..")
	} else {
		fmt.Println("You are just a kid..")
	}

	// we can declare variable inside if construct
	if age2 := 12; age2 >= 18 {
		fmt.Println("Adult")
	} else if age2 >= 10 {
		fmt.Println("Teen")
	}

	// golang does not have ternary operator
}
