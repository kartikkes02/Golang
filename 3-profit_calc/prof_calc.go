package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("Profit Calculator")
	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)

	EarningsBeforeTax := revenue - expenses
	EarningsAfterTax := EarningsBeforeTax * (1 - taxRate/100)
	ratio := EarningsAfterTax / revenue

	// sprintf is used to format the output as a string, which can be useful for displaying the results in a specific format or for further processing.
	// fmt.Sprintf is used to format the output as a string, which can be useful for displaying the results in a specific format or for further processing.

	formattedEBT := fmt.Sprintf("EBT: %.0f \n", EarningsBeforeTax)
	formattedEAT := fmt.Sprintf("EAT: %.0f \n", EarningsAfterTax)

	fmt.Printf("Earnings Before Tax: %.0f\n", EarningsBeforeTax)
	fmt.Printf("Net Profit: %f\n", EarningsAfterTax)
	fmt.Printf("Profit Margin: %.2f\n", ratio)

	fmt.Printf("Earnings Before Tax: %.0f\nNet Profit: %f\n", EarningsBeforeTax, EarningsAfterTax)
	// fmt.Println("Using Sprintf")
	fmt.Print(formattedEBT, formattedEAT)
	fmt.Println("line break")

	// line break through `` is used to create a multi-line string, which can be useful for formatting the output in a more readable way.

	fmt.Printf(`Earnings Before Tax: %.0f\n
	line break
	Net Profit: %f\n`, EarningsBeforeTax, EarningsAfterTax)
}
