package main

import "fmt"

// go routine synchronizer
func task(done chan bool) {
	defer func() {done <- true}() // use defer, as even if func has error, onces, it gets exited, it'll be executed
	fmt.Println("Processing...")
}

func main() {
	done := make(chan bool) // unbuffered channel, creates blocking while sending & receiving
	go task(done)

	<- done // block till, something is being sent over the channel : synchronizer
}

// we are doing the same thing which we did using wait groups