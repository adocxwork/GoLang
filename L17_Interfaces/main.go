package main

import "fmt"

// Interface
type paymenter interface {
	// pay(amount float32) bool
	pay(amount float32)
}

type payment struct{
	// gateway stripe
	// gateway razorpay // we'll need to change this again & again; soln --> interfaces
	gateway paymenter
}

// Open close principle
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)

	// stripePaymentGW := stripe{}
	// stripePaymentGW.pay(amount)

	p.gateway.pay(amount)
}

// -------------------------------------------------

type razorpay struct{}
func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct{}
func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe", amount)
}

type paypal struct {}
func (p paypal) pay(amount float32)  {
	fmt.Println("Making payment using paypal", amount)
}

type fakePayment struct {}
func (f fakePayment) pay(amount float32)  {
	fmt.Println("Making payment using fakePayment", amount)
}

func main() {
	// stripePaymentGW := stripe{}
	// razorpayPaymentGW := razorpay{}
	paypalGW := paypal{}
	// fakeGW := fakePayment{}

	newPayment := payment{
		gateway: paypalGW,
	}
	newPayment.makePayment(100)
}


// No need to write "implements" keyword for implementing an interface, if the method signature matches,
// it will automatically understand that, a struct is implementing an interface