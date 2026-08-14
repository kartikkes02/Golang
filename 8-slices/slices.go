package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods

func main() {
	fmt.Println("Slices")
	// uninitialized slice in null
	var nums []int
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	fmt.Println("Another Method(using make())")
	var nums1 = make([]int, 2, 5) // make(Type , initialization, capacity)
	nums1[0] = 2
	fmt.Println(nums1)
	fmt.Println(nums1 == nil)
	fmt.Println(len(nums1))
	// capacity -> maximum no. of elements can fit.
	fmt.Println(cap(nums1))

	nums1 = append(nums1, 7)
	fmt.Println("Appending 7 into nums1: ", nums1)
	fmt.Print("Capacity: ")
	fmt.Println(cap(nums1))

	// if we append the elements more than the capacity so, capacity return the double.
	nums1 = append(nums1, 7)
	nums1 = append(nums1, 7)
	nums1 = append(nums1, 8)
	fmt.Print("Updated Capacity: ")
	fmt.Println(cap(nums1))

	fmt.Println("Another method")
	nums3 := []int{}
	nums3 = append(nums3, 2)
	fmt.Println(nums3)

	fmt.Println("Copy Function: ")
	var nums4 = make([]int, 2, 3)
	var nums5 = make([]int, 1)

	nums4 = append(nums4, 3)  // 0 0 3
	nums5 = append(nums5, 7)  // 0 7
	fmt.Println(nums4, nums5) // 0 0 3 , 0 7

	copy(nums4, nums5)        // nums4 copied from nums 5
	fmt.Println(nums4, nums5) // 0 7 3 , 0 7

	// Slicing
	fmt.Println("Slicing: ")
	var nums6 = []int{1, 2, 3, 4}
	fmt.Println(nums6[0:3])
	fmt.Println(nums6[0:])
	fmt.Println(nums6[:3])

	// slice package
	var nums7 = []int{1, 2}
	var nums8 = []int{1, 2}

	fmt.Println(slices.Equal(nums7, nums8))

	var nums9 = [][]int{{1, 2}, {2, 3}}
	fmt.Println(nums9)
}
