package main

import (
	"fmt"
	"time"
)

// Goroutines -> a lightweight thread of execution.
// Goroutines are managed by the Go runtime and are multiplexed onto a smaller number of OS threads.
// Goroutines are cheaper than threads and can be created in large numbers.
// Goroutines are used to perform concurrent tasks in Go.

func task(id int) {
	fmt.Println("Task doing: ", id)
}
func main() {

	fmt.Println("Goroutines in Go!")
	// here doesn't give order of execution, due to the nature of goroutines,
	// they are executed concurrently and the order of execution is not guaranteed.
	for i := 0; i < 7; i++ {
		go task(i)
	}
	time.Sleep(time.Second * 2) // wait for all goroutines to finish
	// here we use time.Sleep beacuse we don't have any way to wait for the goroutines to finish, 
	// so we use time.Sleep to wait for a second before exiting the program.

}
