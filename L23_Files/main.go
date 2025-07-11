package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		// Log the error
		panic(err)
	}
	
	fileInfo, err := f.Stat()
	if err != nil {
		// Log the error
		panic(err)
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
}