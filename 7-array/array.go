package main

import "fmt"

func main() {
	// zeroed values
	// int -> 0, string -> "", bool -> FALSE
	var nums [3]int
	fmt.Println(nums)
	nums[1] = 3
	fmt.Println(nums)
	fmt.Println("Length of nums: ", len(nums))

	var nums1 [4]string
	fmt.Println(nums1)
	nums1[1] = "golang"
	fmt.Println(nums1)

	var nums2 [4]bool
	fmt.Println(nums2)
	nums2[1] = true
	nums2[3] = true
	fmt.Println(nums2)

	// to declare in single line
	nums3 := [3]int{2, 8}
	fmt.Println(nums3)

	// 2d array
	nums4 := [2][2]int{{2, 6}, {8, 9}}
	fmt.Println(nums4)

	// fixed size - i.e predictable
	// Memory optimization
	// Constant time access
}
