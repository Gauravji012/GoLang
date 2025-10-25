package main

import "fmt"

/*
slices dynamic array
most useful construct in go
useful method present with slices

how to declare array using slices ::
1. var nums [] int
2. var nums = make([]int, 2)
3. var nums = make([]int, 0)
4. nums := []int {}

copy function
slice operator
slice package
2D slice
*/
func main() {
	// var nums [] int
	// nums = append(nums, 1)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums)) //--- if we want our array should not be nil that means we can use make function to create it

	//make(type, size, capacity)
	// var nums = make([]int, 2)
	// var nums = make([]int, 0) /// if we want to create empty slice then we put 0 at the place of size
	// fmt.Println(nums, "\n", nums==nil, "\n", len(nums), cap(nums))
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// // nums = append(nums, 1)

	// fmt.Println(nums, "\n", nums==nil, "\n", len(nums), cap(nums))

	// 	nums := []int {}
	//  fmt.Println(nums, "\n", nums==nil, "\n", len(nums), cap(nums))
	// 	nums = append(nums, 1)
	// 	fmt.Println(nums, "\n", nums==nil, "\n", len(nums), cap(nums))

	///// copy function

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 1)
	// var nums2 = make([]int, len(nums))
	// //copy function
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	//// slice operator

	// var nums = []int {1,2,3,4,5,6,7,8}
	// //nums[from : to]
	// fmt.Println(nums[0:4])
	// fmt.Println(nums[:4])
	// fmt.Println(nums[4:])

	// slice package

	// var nums1 = []int {1,2,3}
	// var nums2 = []int {1,2,3}

	// fmt.Println(slices.Equal(nums1, nums2))



	// 2D slice 
	var nums = [][]int {{1,3,4}, {5,4,6}}
	fmt.Println(nums)
}
