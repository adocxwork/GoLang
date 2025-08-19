package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func main() {
	// Check if the user provided a file path as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: file-reader <file-path>")
		os.Exit(1)
	}

	// Get the file path from the command line arguments
	filePath := os.Args[1]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Create a new scanner to read through the file
	scanner := bufio.NewScanner(file)

	// Loop through the lines in the file and print them to the console
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Print each line
	}

	// Check for errors while reading the file
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading the file:", err)
	}
}
