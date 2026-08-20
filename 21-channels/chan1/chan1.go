package main

import (
	"fmt"
	"time"
)

// Channels are used to communicate between goroutines.
// Channels are a way to send and receive values between goroutines.
// Channels are a way to communicate between goroutines.
// channels is blocked until some other goroutine sends a value into it or receives a value from it.

func Process(numChan chan int) {
	fmt.Println("Process: ", <-numChan)
}

func main() {
	fmt.Println("Channels in Go!")

	numChan := make(chan int)   // Create a channel of type int
	go Process(numChan)         // Start a goroutine to process the channel
	numChan <- 42               // Send a value into the channel
	time.Sleep(time.Second * 2) // Sleep for a second to allow the goroutine to finish

}
