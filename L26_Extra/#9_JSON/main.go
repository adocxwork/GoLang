package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{Name: "Aditya", Age: 19, IsAdult: true}
	fmt.Println("Person data :", person)

	// Convert person into JSON : Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil  {
		fmt.Println("Error Marshalling :", err)
		return
	}
	fmt.Println("Person data in JSON :", string(jsonData))

	// Decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error Unmarshalling :", err)
		return
	}
	fmt.Println("Person data after Unmarshalling :", personData)
}