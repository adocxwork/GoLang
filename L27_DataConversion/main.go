package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ----------------------------
	// Integer to Float Conversion
	// ----------------------------
	var intVal int = 42
	fmt.Println("Original Integer Value:", intVal)
	fmt.Printf("Type: %T\n\n", intVal)

	// Convert int to float64
	var floatVal float64 = float64(intVal)
	fmt.Println("Converted to Float64:", floatVal)
	fmt.Printf("Type: %T\n\n", floatVal)

	// ----------------------------
	// Integer to String Conversion
	// ----------------------------
	// Using strconv.Itoa (Integer to ASCII)
	var strFromInt string = strconv.Itoa(intVal)
	fmt.Println("Converted to String:", strFromInt)
	fmt.Printf("Type: %T\n\n", strFromInt)

	// ----------------------------
	// String to Integer Conversion
	// ----------------------------
	// Using strconv.Atoi (ASCII to Integer)
	intFromStr, err := strconv.Atoi(strFromInt)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("Converted back to Integer:", intFromStr)
		fmt.Printf("Type: %T\n\n", intFromStr)
	}

	// ----------------------------
	// String to Float Conversion
	// ----------------------------
	floatStr := "5.4"
	floatFromStr, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Println("Converted String to Float64:", floatFromStr)
		fmt.Printf("Type: %T\n", floatFromStr)
	}
}
