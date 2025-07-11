package main

import "fmt"

func main() {
	// iterating over data structures
	
	nums := []int{6, 7, 8}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	for idx, num := range nums {
		fmt.Println(idx, num)
	}

	m := map[string]string {"adi" : "Aditya", "ayu" : "Ayush"}
	for key, val := range m {
		fmt.Println(key, val)
	}

	for key := range m {
		fmt.Println(key)
	}

	for k, v := range "golang" {
		fmt.Println(k, v, string(v))
	}
}
