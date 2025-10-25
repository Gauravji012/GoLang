package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add1(a, b int, c float32) float32 { // when 2 parameter having same data type then we can define with the last parameter variable
	return float32(a+b) + c
}

// we can return multiple values in go

func getLang() (string, string, string) {
	return "go", "js", "c"
}

func help(fn func(a int) string) {
	fmt.Println(fn(1))
}

func help1() func(a int) string {
	return func(a int) string {
		return "called with return vvavlue as function"
	}
}

func main() {
	fmt.Println(add(3, 5))
	fmt.Println(add1(3, 5, 2.5))
	fmt.Println(getLang())
	l1, _, l3 := getLang() // with the help of _ we can bypass one / more variable values
	fmt.Println(l1, l3)

	/// we can save the func in a variable &&& we can also return function
	// 1. in a variable

	// fn := func(a int)string {
	// 	return "func called with fn parameter"
	// }
	// help(fn)

	// 2 as a return type
	fn := help1()
	fmt.Println(fn(6))
}
