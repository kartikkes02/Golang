package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this is for the sending.
func ProcessNum1(numChan chan int) {
	for num := range numChan {
		fmt.Println("ProcessNum: ", num)
		time.Sleep(time.Second * 1) // Sleep for a second to simulate processing time
		// (it's not necessary, but it makes the output more readable)
	}
}

func main() {
	numChan := make(chan int) // Create a channel of type int
	go ProcessNum1(numChan)   // Start a goroutine to process the channel

	for {
		numChan <- rand.Intn(1000) // Send a random value into the channel
		// Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default
		// Source. It panics if n <= 0.
	}

}
