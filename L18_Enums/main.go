package main

import "fmt"

// custom types
// type myType string

// enumerated types
// type OrderStatus int
type OrderStatus string

const (
	// Received OrderStatus = iota // helps generate consecutive values
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing status to", status)
}

func main() {
	changeOrderStatus(Prepared)
}

// we implement enums in go using const & custom types
