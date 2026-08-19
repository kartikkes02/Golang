package main

import "fmt"

// open close principles
// Open for extension: You should be able to add new functionality.
// Closed for modification: You should not need to change existing, tested code to add that functionality.

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	// gateway  razorpay // stripe
	// gateways stripe
	// gateway  paymenter
	gateways paymenter
}

func (p payment) makePayment(amounts float32) {
	// razorPayPaymentGw := razorpay{}
	// razorPayPaymentGw.pay(amounts)
	// p.gateway.pay(amounts)
	p.gateways.pay(amounts)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using Razorpay")
}

type stripe struct{}

func (s stripe) pay(amountss float32) {
	fmt.Println("Making payment using stripe: ", amountss)
}

func main() {
	// Interfaces -> a way to define a set of methods that a type must implement.
	// An interface type is defined by a set of methods.
	// A type implements an interface by implementing its methods.
	// A type can implement multiple interfaces.

	fmt.Println("Interfaces")
	// we make stripePayment here we have also run without this also, but
	// in the real world used this because if this method doesn't used so error comes.
	stripePayment := razorpay{}
	// here if we use razorpay{} then it will give error because we have not implemented the
	// pay method for razorpay.

	// and after used interface paymenter we used here razorpay{} and it will work because we have
	//  implemented the pay method for razorpay.
	newPayment := payment{
		gateways: stripePayment,
	}
	newPayment.makePayment(100.50)
}

// Interface Main Code

// package main

// import "fmt"

// type paymenter interface {
// 	pay(amount float32)
// }

// type payment struct {
// 	gateways paymenter
// }

// func (p payment) makePayment(amounts float32) {
// 	p.gateways.pay(amounts)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	fmt.Println("Payment of", amount, "made using Razorpay")
// }

// type stripe struct{}

// func (s stripe) pay(amountss float32) {
// 	fmt.Println("Making payment using stripe: ", amountss)
// }

// func main() {
// 	// we make stripePayment here we have also run without this also, but
// 	// in the real world used this because if this method doesn't used so error comes.
// 	stripePayment := razorpay{}
// 	// here if we use razorpay{} then it will give error because we have not implemented the
// 	// pay method for razorpay.

// 	// and after used interface paymenter we used here razorpay{} and it will work because we have
// 	//  implemented the pay method for razorpay.

// 	newPayment := payment{
// 		gateways: stripePayment,
// 	}
// 	newPayment.makePayment(100.50)
// }
