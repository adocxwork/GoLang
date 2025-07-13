package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		// Log the error
		panic(err) // abhi kuch nhi ho sakta hai : panic mode
	}
	
	fileInfo, err := f.Stat()
	if err != nil {
		// Log the error
		panic(err)
	}

	fmt.Println("File Name :", fileInfo.Name())
	fmt.Println("IsFolder :", fileInfo.IsDir())
	fmt.Println("Size :", fileInfo.Size())
	fmt.Println("Mode :", fileInfo.Mode())
	fmt.Println("Modification time :", fileInfo.ModTime())
}