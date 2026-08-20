// synchronization is a technique that allows multiple goroutines to coordinate their execution
// and share data safely.
// In this example, we use a channel to synchronize the completion of a goroutine.
// The `making` function sends a signal to the `done` channel when it finishes processing,
// allowing the main function to wait for its completion before exiting.
package main

import (
	"fmt"
)

// goroutine synchronizer
func making(done chan bool) {
	defer func() {
		// defer is used to ensure that the signal is sent to the done channel even if an error occurs
		// during processing.
		done <- true
	}()
	fmt.Println("Processing.....")
}
func main() {
	// unbuffered channel is the default type of channel in Go. 
	// It has no capacity and will block the sending goroutine until another goroutine receives the 
	// value from the channel.
	done := make(chan bool)
	go making(done)
	<-done // blocking

	// dones := <- done
	// fmt.Println(dones)	// output: true
}
