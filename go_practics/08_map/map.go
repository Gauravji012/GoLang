package main

import (
	"fmt"
	"maps"
)

/*
declaration

geet an element
setting an element

maps package

*/
// maps -> hash / object / dictionary
func main() {

	// m := make(map[string]string)

	// //setting an element
	// m["name"] = "gorav"
	// m["lang"] = "backend"

	// //get aan element
	// fmt.Println(m["name"], m["lang"], m["phone"])
	// IMP:::
	// m["phone"] returns zero value it means
	// if we have int -> 0 / string = "" / bool = false ----- these are zero values
	// m := make(map[string]int)
	// m["age1"] = 27
	// m["age2"] = 28
	// m["age3"] = 29
	// m["age4"] = 233
	// fmt.Println(m["age"], m["pon"]) // here output is 27 0
	// fmt.Println(len(m))
	//  delete(m, "age2")
	//  fmt.Println(m)
	//  clear(m)
	//  fmt.Println(m)

	//direct creation of map without make function
	// m := map[string]int{"price":40, "phones":3}
	// fmt.Println(m)

	// //how to find element exist or not
	// v, ok := m["phones"]
	// if ok {
	// 	fmt.Println("value is all ok", v)
	// } else {
	// 	fmt.Println("value is not ok", v)
	// }

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}

	fmt.Println(maps.Equal(m1, m2))
}
