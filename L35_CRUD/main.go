package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response :", res.Status)
		return
	}
	
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error :", err)
	// 	return
	// }
	// fmt.Println("Data :", string(data))
	
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("Todo :", todo)
	fmt.Println("Title response :", todo.Title)
	fmt.Println("Completed response :", todo.Completed)
}

func performPostRequest() {
	todo := Todo {
		UserID: 23,
		Title: "Aditya",
		Completed: true,
	}

	// Convert Todo struct into JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	// Convert JSON into string
	jsonString := string(jsonData)
	// Convert string into reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	// Send POST request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Response status :", res.Status)
}

func performUpdateRequest() {
	todo := Todo {
		UserID: 2343,
		Title: "Golang",
		Completed: false,
	}

	// Convert Todo struct into JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	// Convert JSON into string
	jsonString := string(jsonData)
	// Convert string into reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Create PUT request
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer res.Body.Close()
	
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Response status :", res.Status)
}

func performDeleteRequest() {
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Create Delete Request
	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status :", res.Status)
}

func main() {
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}
