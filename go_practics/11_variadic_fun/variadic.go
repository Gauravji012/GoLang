package main

import (
	"fmt"
)

func sum(num ...interface{}) int {
	s := 0
	for _, val := range num {
		// if v, ok:= val.(int); ok {/
		if v, ok := val.(int); ok {
			s += v
			// fmt.Println(v, ok)
		} else {
			continue
		}
	}
	return s
}

// varriadic function -> we can pass any no of same type argument & perform some operation over it.

// if we need any type of value it means ***** we should use interface{} to accept any type of value
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3))
	fmt.Println("hey ", sum(1, 2, 3, "gaurav"))

	// we can also pass slice array
	nums := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...))
}
