package main

import (
	"fmt"
	"sync"
)

// WaitGroups are used to wait for a collection of goroutines to finish executing.
func task(id int, r *sync.WaitGroup) {
	defer r.Done() // Signal that this goroutine has finished
	// defer is used to ensure that Done() is called when the function exits,
	// even if it exits due to an error or panic.
	fmt.Println("Task done: ", id)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("WaitGroups in Go!")
	for i := 0; i < 7; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine

		go task(i, &wg) // Pass the WaitGroup pointer to the goroutine
	}
	wg.Wait() // Wait for all goroutines to finish
}
