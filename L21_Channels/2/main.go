package main

import "fmt"

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func main() {
	result := make(chan int)
	go sum(result, 4, 3)
	res := <-result // blocking (result aayega, tabhi aage badega na...)
	fmt.Println(res)
}
