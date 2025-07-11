package main

import (
	"fmt"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id       string
	amount   float32
	status   string
	customer
}

func main() {
	newCustormer := customer{
		name:  "john",
		phone: "123456789",
	}
	newOrder := order{
		id:       "1",
		amount:   30,
		status:   "received",
		customer: newCustormer,   // struct embedding
		// customer: customer {
		// 	name: "john",
		// 	phone: "123456789",
		// }
	}

	fmt.Println(newOrder)
	newOrder.customer.name = "robin"
	fmt.Println(newOrder)
}
