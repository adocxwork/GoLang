package main

import (
	"fmt"
	"time"
)

// emailChan is receive only channel here, done is send only
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to :", email)
		time.Sleep(time.Second)
	}
}

func main() {
	emailChan := make(chan string, 100) // buffered channel
	done := make(chan bool)             // synchronizer

	go emailSender(emailChan, done)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("Done sending emails..")

	// this is important
	close(emailChan) // closing the channel
	<-done // blocking : synchronizer

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"
	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
}
