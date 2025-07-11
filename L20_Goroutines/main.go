package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task :", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		// task(i)
		// go task(i) // sb ek saath run hoga...!

		// closure with goroutine : anonymous func
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	// Adding delay so that, we can see the work being done (before the main func gets exhausted)
	time.Sleep(time.Second * 2)
}
