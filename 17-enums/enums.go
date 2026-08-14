package main

import "fmt"

// enumerated type -> a way to define a set of named constants.
type OrderStatus int

const (
	Received OrderStatus = iota // untyped integer constant
	Pending
	Completed
)

func changeOrderStatus(status string) {
	fmt.Println("Order status changed to: ", status)
}

func changeOrderStatuss(status OrderStatus) {
	fmt.Println("Order status changed to: ", status)
}
func main() {
	// Enums -> a way to define a set of named constants.
	// In Go, we can use the iota identifier to create enums.
	// iota is a predeclared identifier that represents successive untyped integer constants.
	// It is reset to 0 whenever the word const appears in the source and increments after each const specification in the same block.
	fmt.Println("Enums")
	changeOrderStatus("Received")
	changeOrderStatus("Pending")
	changeOrderStatus("Completed")
	changeOrderStatuss(Received)
	changeOrderStatuss(Pending)
	changeOrderStatuss(Completed)
}
