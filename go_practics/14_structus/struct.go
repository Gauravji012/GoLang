package main

// structure -> combination of different data types
// in go we create structure at the place of classess
// here structure is treated as blueprint of the class/struct
// in go we use struct for object oriented programming

import (
	"fmt"
	"time"
)

/*
// decalear  + create instance + provide relation between fun & struct + know about zero value when we do not intialize any filed of the structure.
//constructor
// embadding of thee 2 structures

type order struct {
	id        string
	status    string
	amount    float32
	createdAt time.Time // nanosecond precision
}

// receiver type function --> kisi struct ko function se connect krne ke liye receiver type ka concept use krte h
func (o *order) changeStatus(stat string) { // we use * because of call by reference concept & want to change actual value
	o.status = stat // here struct perform deferenceing internally
}

func (o order) getAmount() float32 { // here we can use without * because at this place we do not modiffy any data so we can use without it. here we just read the data so dont required pointer(*).
	return o.amount
}

func main() {
	//if we does not set any field of struct, default value is zeero value for that field.
	// int=> 0, float=> 0, string=> "", bool=> false
	myOrder := order{
		id:     "1",
		status: "received",
		amount: 50.00,
	}
	myOrder.createdAt = time.Now()
	fmt.Println("My order is :: ", myOrder, myOrder.status)

	myOrder.changeStatus("confirmed")
	fmt.Println("My order is :: ", myOrder, myOrder.getAmount())

}
*/

// now we teach about constructor hack because directly there is no method of constructor.
// to create constructor --> we always start it with new keyword with any name --> it is just a convension

/*
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time // nanosecond precision
}

// trying to create constructor
func newOrder(id string, amt float32, st string) *order {
	//initial setup goes here ....
	myOrder := order {
		id : id,
		amount : amt,
		status : st,
	}
	return &myOrder // we return pointer of the struct instead of just directly return struct ---> convention hai not an compulsary
}

func main() {

// can create a struct which only used once
	lan := struct {
		name string
		isGood bool
	} {"gaurav", true}

	fmt.Println(lan)
	myOrd := newOrder("1", 50.30, "confirmed")
	fmt.Println(myOrd)



	// we can embade the structure as well

}
*/

// embadding the structure

type cust struct {
	name  string
	phone string
}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	cust
}

func newOrder(id string, amt float32, st string, name string, phone string) order {
	newCust := cust{
		name:  name,
		phone: phone,
	}
	myOrder := order{
		id:     id,
		amount: amt,
		status: st,
		cust:   newCust,
	}
	return myOrder
}

func main() {
	myOrd := newOrder("1", 50.30, "confirmed", "gaurav", "098776078596")
	fmt.Println(myOrd, myOrd.cust, myOrd.cust.name)
}
