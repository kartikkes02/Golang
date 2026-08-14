package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	fmt.Println("Maps")

	// creating map
	// n := make(map[type]type)	declaration
	m := make(map[string]string)

	// setting an element
	m["name"] = "kartik"
	m["address"] = "Prayag"

	// get an element
	fmt.Println("Name: ", m["name"], "Phone: ", m["phone"], "Address: ", m["address"])

	// if key doesn't exists in the map it returns zero.
	n := make(map[string]int)
	n["age"] = 10
	fmt.Println("Age: ", n["age"])
	fmt.Println(n["phone"]) // 0

	// Length
	fmt.Println("Length of m: ", len(m))

	// delete an element
	delete(m, "address")
	fmt.Println("Delete(address) m: ", m)

	// clear all the element.
	clear(m)
	fmt.Println("Clear m: ", m)

	// Another method
	c := map[string]int{"state": 7}
	fmt.Println(c)

	// check the that elements is in the map or not
	v, ok := c["state"]
	fmt.Println("Value of State: ", v) // Returns Value
	if ok {
		fmt.Println("All Ok")
	} else {
		fmt.Println("Not Ok")
	}

	// checks the two maps are equal or not
	m3 := map[string]int{"name": 3, "class": 8}
	m4 := map[string]int{"name": 3, "class": 9}

	fmt.Println(maps.Equal(m3, m4))
}
