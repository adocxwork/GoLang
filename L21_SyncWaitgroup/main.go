package main

import (
	"fmt"
	"sync"
)

// worker simulates a task and signals completion using wg.Done()
func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Marks this goroutine as done when function exits
	fmt.Println("Work started :", i)
	fmt.Println("Work ended :", i)
}

func main() {
	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish

	for i := 1; i <= 3; i++ {
		wg.Add(1)           // Increment counter before starting a goroutine
		go worker(i, &wg)   // Launch worker as a goroutine
	}

	wg.Wait() // Wait until all goroutines have called Done()
	fmt.Println("Worker's task completed.")
}
