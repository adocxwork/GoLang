package main

import "fmt"

func add(a int, b int) int { return a + b }

func main() {
	fmt.Println(add(3, 4))
	// fmt.Printf("%d", "hello") // go vet <file name> : catches this error
}

// go fmt <file name>
// formats the code..

// go vet <file name>
// A static analysis tool that looks for bugs.
// It checks for things go fmt won't catch (e.g., unreachable code, misused format strings).