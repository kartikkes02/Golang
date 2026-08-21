// Mutex is a Go program that demonstrates the use of mutexes for synchronizing access to shared
// resources in concurrent programming.
// It's also known as mutual exclusion, which ensures that only one goroutine can access a shared
// resource at a time.
// Race conditions can occur when multiple goroutines access shared resources concurrently,
// leading to unpredictable behavior.
package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex // Mutex to synchronize access to the views field
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock() // Unlock the mutex to allow other goroutines to access the views field
		wg.Done()
	}() // Signal that this goroutine has finished

	// lock does when the modification is does.
	p.mu.Lock() // Lock the mutex to ensure exclusive access to the views field
	p.views += 1

}
func main() {
	var wg sync.WaitGroup
	fmt.Println("Mutex")

	postInc := post{
		views: 0,
	}

	for i := 0; i < 100; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go func() {
			postInc.inc(&wg)
		}()
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println(postInc.views)
}
