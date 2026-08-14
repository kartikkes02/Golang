package main

import "fmt"

// for -> only constructs in go for looping

func main() {

	// for loop
	for i := 0; i < 5; i++ {
		// fmt.Print(i)
		fmt.Println(i)
	}

	// for (range)
	for i := range 3 {
		fmt.Print(i)
	}

	// while loop
	i := 1
	for i < 4 {
		fmt.Print(i)
		i++

		// infinite loop
		// for {
		// 	fmt.Print("1")
		// }

	}
}
