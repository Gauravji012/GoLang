package main

import "fmt"

// closure -> func ke close ho jane pr bhi variable ka present rhna --> closure
// for ex: count variable is still present after returning the value beccause it is initailize / declear outside of this function

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment())
}
