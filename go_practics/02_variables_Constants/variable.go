package main

import "fmt"

func main()  {
	//string
	var name string = "gaurav"
	var naam = "same gaurav" // type ko infer kr leti h golang



	// var isbool = true
	isbool:=true // is case declar + initialize bhi krna hota h
	fmt.Println(isbool)
	fmt.Println(name)
	fmt.Println(naam)


	const (
		mon = 1
		tues = 2

	)
	fmt.Println(mon,  "---" , tues)

	//cosntants
	const  ji = "hello golang"

	fmt.Println(ji)

}