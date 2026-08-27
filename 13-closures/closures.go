package main

import "fmt"

closure -> function which is defined inside another function and it can access the variables
// of outer function.
func counter() func() int {
	var count = 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	fmt.Println("Closures")
	inc := counter()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	// fmt.Println(counter())
}
