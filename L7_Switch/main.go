package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple switch
	i := 7
	switch i {
	case 1:
		fmt.Println("One") // no break keyword is needed
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}


	// Multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Working day")
	}

	// Type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other")
		}
	}
	whoAmI(7)
}