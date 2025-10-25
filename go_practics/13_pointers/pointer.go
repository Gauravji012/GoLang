package main

import "fmt"

// pointer -> it is kind of variable, used to store the address of the variable

/*

//call by value where parameters copy is created inside the called function 

func changeNum(num int) {
	num = 5
	fmt.Println("Inside change is :: ", num)
}

func main()  {
	num := 1
	changeNum(num)
	fmt.Println("outside / after change is :: ", num)
}
*/


// call by reference

func changeNum(num *int) {
	*num = 5
	fmt.Println("Inside change is :: ", num, *num)
}

func main()  {
	num := 1
	changeNum(&num)
	fmt.Println("outside / after change is :: ", &num, num)
}