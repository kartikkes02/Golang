package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Variadic Functions")
	result := sum(1, 2, 3)

	nums1 := []int{1, 2, 3, 4}
	result1 := sum(nums1...)
	fmt.Println(result)
	fmt.Println(result1)
}
