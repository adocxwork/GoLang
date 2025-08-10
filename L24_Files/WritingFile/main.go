package main

import (
	"os"
)

func main() {
	f, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("Hi there...\n")
	f.WriteString("Nice Language...\n")
	bytes := []byte("Hello Golang....")
	f.Write(bytes)

}