package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning Web Service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer res.Body.Close() // closing connection/response
	fmt.Printf("Type of response : %T\n", res)
	// fmt.Println("Response :", res)
	
	// Reading response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println("Response :", string(data))
}