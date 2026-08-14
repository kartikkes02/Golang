package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Investment calculator: ") // fmt.Println() is used to print the string "Investment calculator: " to the console.
	fmt.Print("Enter Investment Amount: ")
	var investmentAmount float64
	fmt.Scan(&investmentAmount)
	// fmt.Scan() is used to read input from the user and store it in the variable investmentAmount.
	// The & operator is used to get the address of the variable, which is required by fmt.Scan() to store the input value.

	fmt.Print("Enter Expected Return Rate: ")
	var expectedReturnRate float64
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Number of Years: ")
	var years float64
	fmt.Scan(&years)

	// const investmentAmount float64 = 8000
	// investmentAmount = 10000   // we cannot change the value of a constant, it will cause a compile-time error
	// var expectedReturnRate float64 = 5
	// var years float64 = 10.1
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Print("Future Value: ", futureValue)
}
