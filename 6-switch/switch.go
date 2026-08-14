package main

import (
	"fmt"
	"time"
)

func main() {
	// switch
	fmt.Print("Enter the no.: ")
	var j float64
	fmt.Scan(&j)

	switch j {
	case 1:
		fmt.Print("1")

	case 2:
		fmt.Print("2")

	case 3:
		fmt.Print("3")

	default:
		fmt.Print("error")
	}
	fmt.Println()

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Print("Weekend")

	default:
		fmt.Print("Workday")
	}
}
