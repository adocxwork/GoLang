package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating map
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"])

	// if key is not found, it returns zero value
	fmt.Println(m["phone"])

	m2 := make(map[string]int)
	m2["age"] = 18
	m2["price"] = 50
	fmt.Println(m2["age"])
	fmt.Println(m2["phone"])
	fmt.Println("Length of map :", len(m2))
	fmt.Println(m2)
	delete(m2, "price")
	fmt.Println(m2)
	
	// empty the map
	clear(m2)
	fmt.Println(m2)
	fmt.Println()

	// creating map
	m3 := map[string]int {"price":3, "age":50}
	fmt.Println(m3)
	fmt.Println()

	// check if the element is present or not in map
	val, isPresent := m3["price"]
	fmt.Println("price :", val)
	if isPresent {
		fmt.Println("present")
	} else {
		fmt.Println("not present")
	}

	fmt.Println()
	// compare 2 maps
	ma1 := map[string]int {"price":3, "age":50}
	ma2 := map[string]int {"price":3, "age":50}
	// fmt.Println(ma1 == ma2) // invalid
	fmt.Println(maps.Equal(ma1, ma2)) // valid
}