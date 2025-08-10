package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be zero.")
	}
	return a/b, nil
}

func main() {
	// ans, _ := divide(10, 0)
	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Division of two no :", ans)
}