package main

import (
	"fmt"
	"os"
)

func main() {
	// Reading file at ones
	// data, err := os.ReadFile("example.txt") // not good for large files, loads everything at ones
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))


	// Reading Directory
	// dir, err := os.Open(".")
	dir, err := os.Open("../../")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1) // returns all for -ve no

	for _, fi := range fileInfo {
		fmt.Println(fi.Name(), fi.IsDir())
	}
}