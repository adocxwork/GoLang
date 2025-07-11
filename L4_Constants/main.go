package main
import "fmt"

// const variables can change their values
const maxVal int = 10

// shorthand syntax are not allowed outside the function
// maxVal2 := 15
func main() {
	fmt.Println(maxVal)

	// const grouping
	const (
		port = 5500
		host = "localhost"
	)
	fmt.Println(port, host)
}