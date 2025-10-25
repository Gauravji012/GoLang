package main

import "fmt"


func main() {
	//how to use if / if else if/ if else

	// age := 18

	// if age >= 18 {
	// 	fmt.Println("person is an adult")

	// } else if age >= 12 {
	// 	fmt.Println("person is teenager")
	// } else {
	// 	fmt.Println("person is child")
	// }


	//logical operators ---
	admin := "gk"
	haspermission := false

	if admin == "gk" && haspermission {
		fmt.Println("you are valid user")
	} else {
		fmt.Println("Happy trying")
	}
//we can declare a variable inside if construct 
	if age := 15; age >= 18 {
		fmt.Println("you are an adult")
	} else if age >= 12 {
		fmt.Println("you are an teeenager", age)
	}
	// fmt.Println(age)


	// go does not have ternary operators feature, we will havee to use normal if else to handle this scenario

}
