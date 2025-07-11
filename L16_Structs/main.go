package main

import (
	"fmt"
	"time"
)

type order struct {
	id       string
	amount   float32
	status   string
	createAt time.Time
}

// constructor for struct
func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// if u don't set any values, default value is zero value
	// int -> 0, float -> 0.0, string -> "", bool -> false

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 54.3,
	// 	status: "received",
	// }
	myOrder := newOrder("1", 54.3, "received")


	// Adding a new field
	myOrder.createAt = time.Now()
	fmt.Println(myOrder)

	fmt.Println("Status :", myOrder.status)
	myOrder.changeStatus("Confirmed")
	fmt.Println("Status :", myOrder.status)
	fmt.Println("Amount :", myOrder.getAmount())

	// one time use only struct without name
	language := struct {
		name string
		isGood bool
	} { "golang", true }
	fmt.Println(language)
}
