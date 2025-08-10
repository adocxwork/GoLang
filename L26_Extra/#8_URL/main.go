package main

import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Can not parse URL :", err)
		return
	}
	fmt.Printf("Type of URL : %T\n",parsedUrl)

	fmt.Println("Scheme :", parsedUrl.Scheme)
	fmt.Println("Host :", parsedUrl.Host)
	fmt.Println("Path :", parsedUrl.Path)
	fmt.Println("RawQuery :", parsedUrl.RawQuery)

	// Modifying URL components
	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=aditya"

	// Construct a URL string from URL object
	newURL := parsedUrl.String()
	fmt.Println("New URL :",newURL)
}