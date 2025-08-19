package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parseInputs parses two string values into integers, returning an error if parsing fails.
func parseInputs(v1, v2 string) (int, int, error) {
	val1, err := strconv.Atoi(v1)
	if err != nil {
		return 0, 0, errors.New("invalid input: value 1 must be a number")
	}
	val2, err := strconv.Atoi(v2)
	if err != nil {
		return 0, 0, errors.New("invalid input: value 2 must be a number")
	}
	return val1, val2, nil
}

// add returns the sum of two integers.
func add(v1, v2 int) int {
	return v1 + v2
}

// subtract returns the difference of two integers.
func subtract(v1, v2 int) int {
	return v1 - v2
}

// multiply returns the product of two integers.
func multiply(v1, v2 int) int {
	return v1 * v2
}

// divide returns the quotient of two float64 numbers or an error.
func divide(v1, v2 string) (float64, error) {
	val1, err := strconv.ParseFloat(v1, 64)
	if err != nil {
		return 0, errors.New("invalid input: value 1 must be a number")
	}
	val2, err := strconv.ParseFloat(v2, 64)
	if err != nil {
		return 0, errors.New("invalid input: value 2 must be a number")
	}

	if val2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return val1 / val2, nil
}

func main() {
	fmt.Println("--- Welcome to Calculator App ---")
	fmt.Println("Help -> help | Exit -> exit")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		parts := strings.SplitN(input, " ", 3)
		command := parts[0]

		switch command {
		case "add", "subtract", "multiply":
			if len(parts) != 3 {
				fmt.Println("Invalid command. Usage:", command, "<val1> <val2>")
				continue
			}
			val1, val2, err := parseInputs(parts[1], parts[2])
			if err != nil {
				fmt.Println(err)
				continue
			}

			var result int
			switch command {
			case "add":
				result = add(val1, val2)
			case "subtract":
				result = subtract(val1, val2)
			case "multiply":
				result = multiply(val1, val2)
			}
			fmt.Printf("Output: %d\n", result)

		case "divide":
			if len(parts) != 3 {
				fmt.Println("Invalid command. Usage: divide <val1> <val2>")
				continue
			}
			result, err := divide(parts[1], parts[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Output: %f\n", result)

		case "help":
			fmt.Println("--- Commands ---")
			fmt.Println("add <val1> <val2>")
			fmt.Println("subtract <val1> <val2>")
			fmt.Println("multiply <val1> <val2>")
			fmt.Println("divide <val1> <val2>")
			fmt.Println("help")
			fmt.Println("exit")

		case "exit":
			fmt.Println("Exiting...")
			return

		case "":
			continue

		default:
			fmt.Println("Invalid command. Type 'help' for a list of commands.")
		}
	}
}