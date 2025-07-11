package main
import "fmt"

func main() {
	// default values
	// int 0, string "", bool false
	
	var nums[4] int
	nums[0] = 98
	fmt.Println(nums[0])
	fmt.Println(nums)
	fmt.Println(len(nums))

	// declare & initialize
	arr := [3]int {4,5,7}
	fmt.Println(arr)

	// 2D arrays
	arr2 := [2][3]int {{1,2,3}, {4,5,6}}
	fmt.Println(arr2)

	var arr3 = [2][3]int {{1,2,3}, {4,5,6}}
	fmt.Println(arr3)

	// fixed size
	// memory optimization
	// O(1) access time
}