package main

//in go we do not have any enum keyword to make an enum
// enumerated types
// to make this(enums) we use --> custom types --> const & type
// how to make custom types --->>  "type myString string"

import "fmt"

// func changeOrderStauts(status string) {
// 	fmt.Println("changing order status to ", status)
// }

func changeOrderStauts(status OrderStatus) {
	fmt.Println("changing order status to ", status)
}

// here we are converting our string value to integer using iota i.e. it is automatically incerasing by 1
// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

// now we try to accept string instead of int
type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "preapred"
	Delivered             = "delivered"
)

func main() {
	changeOrderStauts(Prepared)
}
