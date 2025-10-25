package main

import "fmt"

//lets understand using payment gateway
// type payment struct {}
// func (p payment) makePayment(amount float32) {
// 	// we call an api of the gatteway e.g. razorpay
// 	razorpayPaymentGW := razorpay{}
// 	razorpayPaymentGW.pay(amount)

// }

// type razorpay struct{}
// func (r razorpay) pay(amount float32) {
// 	// actual logic
// 	fmt.Println("making payment using razorpay", amount)
// }

// func main() {
// 	newPayment := payment{}
// 	newPayment.makePayment(100)

// }

//supoose our new requirement is we need to switch stripe payment gateway instead of razorpay so how can you do that??
/*
type payment struct {}
func (p payment) makePayment(amount float32) { // this method voilates the open cclosed principle rule. we modify the existing code logic insteaad of doing extension of the code logic. so wee need to think it again .
	// we call an api of the gatteway e.g. razorpay
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)
	stripePaymentGW := stripe{}
	stripePaymentGW.pay(amount)
}
type razorpay struct{}
func (r razorpay) pay(amount float32) {
	// actual logic
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}
func (s stripe)pay(amount float32) {
	fmt.Println("making payment using stripe ", amount)
}
func main() {
	newPayment := payment{}
	newPayment.makePayment(100)
	// newPayment_stripe := stripe{}
	// newPayment_stripe.pay(200)	 /// modify inside the payment struct
} ///// but in this code we voilate the open closed principle which tells us open for extension & closed for modification.

// so now how we approach this issue??

*/

// now understand the conccept of interfaces

type paymenter interface {
	pay(amount float32)
}
type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("makeing payment using paymenter interface", amount)
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("makeing payment using fake gateway for testing purpose")
}

// /suppose we need to provide new payment gateway i.e. paypal
type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("paypal is using for making an payment", amount)
}
func main() {
	razorpayPaymentGW := fakepayment{}
	newPayment := payment{
		gateway: razorpayPaymentGW,
	}
	paypalPaymentGW := paypal{}
	paypalPaymentGW.pay(800)
	newPayment.makePayment(400)
}
