// buffered channel is a type of channel in Go that has a capacity to hold a certain number of values.
package main

import (
	"fmt"
	"time"
)

func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true
	}()

	// func emailSender(emailChan chan <- string, done <- chan bool)
	// this means that the emailSender function can only send values to the emailChan channel
	// and doesn't receive values from it.
	// <-done  so this is not running because the done channel is not being used to signal when
	// the emailSender function is done sending emails.

	for email := range emailChan {
		fmt.Println("Email sending to ", email)
		time.Sleep(time.Second) // Sleep for a second to simulate email sending time
	}
}
func main() {
	emailChan := make(chan string, 10) // buffered channel with a capacity of 10
	done := make(chan bool)

	go emailSender(emailChan, done)

	// Send emails to the buffered channel
	for i := 0; i < 10; i++ {
		emailChan <- fmt.Sprintf("%d@example.com", i)
	}
	fmt.Println("Done Sending")
	close(emailChan) // Close the channel to signal that no more emails will be sent
	// (no errors will be sent to the channel after it is closed)
	<-done // Wait for the emailSender goroutine to finish
}

// // Create a buffered channel with a capacity of 2
// 	ch := make(chan int, 2)

// 	// Send values to the buffered channel
// 	ch <- 1
// 	ch <- 2

// 	// Receive values from the buffered channel
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
