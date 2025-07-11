package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock() // unlocks the resource
		wg.Done()
	}()

	p.mu.Lock() // locks the resource, no other process can modify the resource till it completes
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)
}

// Locking through mutex helps to maintain consistency
// Disadvantage is, resource gets locked for other goroutines