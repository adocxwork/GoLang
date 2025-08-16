package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple, banana, mango"
	fruits := strings.Split(data, ",")
	fmt.Println(fruits)

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str = "    Hello Go!     "
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str1 := "Aditya"
	str2 := "Gupta"
	res := strings.Join([]string{str1, str2}, " ")
	fmt.Println(res)
}