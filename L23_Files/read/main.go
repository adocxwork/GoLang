package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 10)

	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {
		fmt.Println("Data :", d, string(buf[i]))
	}
}
