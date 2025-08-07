package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Name : ")
	// var name string
	// fmt.Scan(&name) // single word input
	// fmt.Println("Hello Mr.", name)

	// Reading Sentences
	reader := bufio.NewReader(os.Stdin)
	fullName, _ := reader.ReadString('\n')
	fmt.Println("Full Name :", fullName)
}