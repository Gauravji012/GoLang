package main

import (
	"fmt"
	// "time"
)

func main() {
	//basic switch
	// var i = 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	//multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it is a weekend")
	// default:
	// 	fmt.Println("it is a workday")
	// }

	//type switch statement
	//function jo h go mein class vision bhi hota h ---> func mein parameter pass kr skte ho
	// --> func ko khi se call kr skte package main
	// --> func ko kisi variable mein store bhhi kr skte h

	whoAmI := func(i interface{}) { //interface is an generic type i.e. we can use i for any data type like int, string, bool etc.
		//basically we can pass any value because of interface
		switch t := i.(type) {
		case int:
			fmt.Println("it is an integer")
		case string:
			fmt.Println("it is a string")
		default:
			fmt.Println("it is default option", t)
		}
	}
	whoAmI(55.7)

}
