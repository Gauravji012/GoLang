package main

import "fmt"

// func printSlice(nums []int) {
// 	fmt.Println(nums)
// }

// func printStringSlice(names []string) {
// 	fmt.Println(names)
// }

// we solve this code duplication issue with the help of generics
// func printSlice[T any](items []T){ // we can use interface{} instead of any

// func printSlice[T interface{}](items []T){
// 	for _, item := range(items) {
// 		fmt.Println(item)
// 	}
// } // but still it is not a good practice to accept any type of data.
// suppose if we need only 2 types of data i.e. integer & string then --->

// restrict with some limited types of data values
// func printSlice[T int | string | bool](items []T){
func printSlice[T comparable, V string](items []T, name V) { // we can also use comaprable insteadd of just few data types. this is an extension of previously defined data values/types.
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// now generics on structs ====>>>

func main() {
	// 	// nums := []int{1,2,3,4}
	// 	// printSlice(nums)
	// 	// names := []string{"gaurav", "khandelwal"}
	// 	// printStringSlice(names) /// at this time we need to create 2 seperate function for just printiing the value of the slices. & we are just duplicating the code which is not a good practice.

	// 	// nums := []int{1,2,3,4}
	// 	// printSlice(nums)
	// 	// names := []string{"gaurav", "khandelwal"}
	// 	// printSlice(names)
	vals := []bool{true, false, true}
	printSlice(vals, "gaurava")

}

// // LIFO
// type stack[T any] struct {
// 	elements []T
// }

// func main() {
// 	myStack := stack[string]{
// 		elements: []string{"golang"},
// 	}
// 	myStack1 := stack[int]{
// 		elements: []int{1,2,3,4},
// 	}
// 	fmt.Println(myStack, myStack1)
// }
