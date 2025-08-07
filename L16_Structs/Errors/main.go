package main

import (
	"errors"
	"fmt"
)

// Returning errors...
func doSomething() error {
	return errors.New("something went wrong")
}

func doSomething2() error {
	return fmt.Errorf("failed at step %d: %v", 2, "bad input")
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
