package main

import "fmt"

// iterating over data structures
func main() {
	// nums := []int{1,3,4,5}
	// sum := 0
	// for i, num := range nums {
	// 	fmt.Println(i, num)
	// 	sum += num
	// }
	// fmt.Println("sum is :: ", sum)

	//suppose we have a map
	// m := map[string]string{"fname":"gaurav", "lname":"khandelwal"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	// for key := range m {
	// 	fmt.Println(key)
	// }

	//suppose we have string

	for i, c := range "gaurav" {
		fmt.Println(i, string(c)) // it is usually store in unicode format so we need to perform type casting here
	}

}
