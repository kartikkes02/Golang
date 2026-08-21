// Select is a Go program that demonstrates the use of channels and the select statement.
// It creates multiple channels and uses select to wait on multiple channel operations,
// allowing for concurrent programming.
package main

import "fmt"

func main() {
	fmt.Println("Select")
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "Hello, world!"
	}()

	for i := 0; i < 2; i++ {
		// Retaining local channel/select implementation
		// The select statement is used to wait on multiple channel operations
		select {
		case msg1 := <-chan1:
			fmt.Println("Received from chan1:", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received from chan2:", msg2)
		}
	}
}
