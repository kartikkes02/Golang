package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32) {
	razorPay := razorpay{}
	razorPay.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment of ", amount, " made using Razorpay")
}

func main() {
	// Interfaces -> a way to define a set of methods that a type must implement.
	// An interface type is defined by a set of methods.
	// A type implements an interface by implementing its methods.
	// A type can implement multiple interfaces.

	fmt.Println("Interfaces")
	newPayment := payment{}
	newPayment.makePayment(100.50)
}
