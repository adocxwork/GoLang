package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted...")
}