package main

import (
	"fmt"
	"os"
	// "io"
)

func main() {
	// Creating, Writing
	/*
	// Creating file
	file, err1 := os.Create("f1.txt")
	if err1 != nil {
		fmt.Println("Error while creating file :", err1)
		return
	}
	defer file.Close()
	
	// Writing
	content := "Hello World..."
	byte, err2 := io.WriteString(file, content)
	if err2 != nil {
		fmt.Println("Error while writing into the file :", err2)
		return
	}

	fmt.Println("File created with", byte, "bytes.")
	*/






	// Opening file, Reading it using buffer
	/*
	file, err1 := os.Open("f1.txt")
	if err1 != nil {
		fmt.Println("Error while opening file :", err1)
		return
	}
	defer file.Close()

	// Creating a buffer to read file : buffer is a temporary storage
	buffer := make([]byte, 1024)
	// Read the file content into buffer
	for {
		n, err2 := file.Read(buffer)
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			fmt.Println("Error while reading file :", err2)
			return
		}

		// Process the read content
		fmt.Println(string(buffer[:n]))
	}
	*/







	// Reading file on a single go (not recommended for large files)
	content, err1 := os.ReadFile("f1.txt")
	if err1 != nil {
		fmt.Println("Error while reading file :", err1)
		return
	}
	fmt.Println(string(content))
}