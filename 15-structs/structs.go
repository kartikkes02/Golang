package main

import (
	"fmt"
	"time"
)

// struct -> collection of fields.
// Struct are used to group together data of different types into a single type.
// Structs are useful for creating complex data structures that can represent real-world entities.
// order struct -> collection of fields which are ordered.
type orders struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time // nanosecond precision timestamp
}

// method -> function with a receiver type.
// A method is a function that is associated with a specific type, and it can be called on instances
// of that type.
func (o *orders) changeStatus(status string) {
	o.status = status
	// we can't give the *o beacuse they automatically dereference the pointer and we can access the
	// value of the struct.
}

// constructor -> function that returns a new instance of a struct.
func newOrderss(id string, amount float64, status string) *orders {
	orderss := orders{
		id:     id,
		amount: amount,
		status: status,
	}
	return &orderss
}

// Structure embedding -> a way to create a new struct that includes the fields of another struct.
type orderDetails struct {
	customerName string
}
type orders4 struct {
	id           string
	amount       float64
	status       string
	createdAt    time.Time // nanosecond precision timestamp
	orderDetails           // embedding the orderDetails struct
}

func main() {
	fmt.Println("Structs")

	myOrder := orders{
		id:     "123",
		amount: 100.50,
		status: "pending",
	}

	fmt.Println("Old Order: ", myOrder)
	myOrder.amount = 200.75
	myOrder.createdAt = time.Now()
	fmt.Println("Updated Order: ", myOrder)
	fmt.Println("Order ID: ", myOrder.id)
	myOrder.changeStatus("completed")
	fmt.Println("Updated Order Status: ", myOrder.status)

	// if we don't set any value for the fields of the struct, they will be initialized with their
	// zero values.
	// int -> 0, float64 -> 0.0, string -> "", bool -> false, pointer -> nil, slice -> nil,
	// map -> nil, struct -> zero value of its fields.

	myOrder2 := newOrderss("1", 0, "recived")
	fmt.Println("Empty Order: ", myOrder2)

	// Another way to create the struct
	myOrd := struct {
		id     string
		amount float64
	}{"2", 0.0}

	fmt.Println("Anonymous Struct: ", myOrd)

	fmt.Println("Structs Embedding")

	myCustomer := orderDetails{
		customerName: "John Doe",
	}

	myOrder3 := orders4{
		id:           "123",
		amount:       100.50,
		status:       "pending",
		createdAt:    time.Now(),
		orderDetails: myCustomer,

		// another method
		// 	orderDetails: orderDetails{
		// 		customerName: "John Doe",
		//  }
	}
	fmt.Println("Embedded Order: ", myOrder3)
	fmt.Println("Customer Details: ", myOrder3.orderDetails) // we can access the fields of the embedded struct directly.
}
