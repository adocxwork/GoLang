package main

import (
	// "crypto/rand"
	"math/rand"
	"fmt"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing Number :", num)
		time.Sleep(time.Second)
	}
	// fmt.Println("Processing Number :", <- numChan)
}

func main() {
	// --- Deadlock ---
	// messageChan := make(chan string)
	// messageChan <- "ping" // blocking (channels are blocking)
	// msg := <- messageChan
	// fmt.Println(msg)

	numChan := make(chan int)
	go processNum(numChan)

	// numChan <- 5
	// time.Sleep(time.Second * 2)

	for {
		numChan <- rand.Intn(100) // send a random number to the channel
	}


}