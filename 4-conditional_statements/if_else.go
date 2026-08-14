package main

import "fmt"

func main() {
	var age float64
	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	if age <= 10 {
		fmt.Print("Person is younger")
	} else if age <= 18 {
		fmt.Print("Person is adult")
	} else {
		fmt.Print("Person is old")
	}
}

// go doesn't have ternary, we have to use normal if else
