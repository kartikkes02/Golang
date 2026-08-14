package main

import "fmt"

// pointer -> variable which stores the address of another variable.

func ChangeNum(num int) {
	num = 10
	fmt.Println("Inside ChangeNum function:", num)
}

// dereference -> accessing the value of a variable using its address.
func ChangeNumPointer(num *int) {
	*num = 10
	fmt.Println("Inside ChangeNumPointer function:", *num)
}

func main() {
	fmt.Println("Pointers")
	nums := 5
	ChangeNum(nums)
	fmt.Println("After changeNum in main:", nums)

	ChangeNumPointer(&nums)
	fmt.Println("After changeNumPointer in main:", nums)
}
