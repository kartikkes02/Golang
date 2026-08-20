// for recieving
package main

import (
	"fmt"
)

func sum(results chan int, num1 int, num2 int) {
	numResult := num1 + num2
	results <- numResult // Send the result into the channel
}
func main() {
	result := make(chan int) // Create a channel of type int
	go sum(result, 10, 20)   // Start a goroutine to calculate the sum

	// Receive the result from the channel
	numResults := <-result

	fmt.Println("Result: ", numResults)
}
