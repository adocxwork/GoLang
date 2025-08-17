package main

import (
	"fmt"

	// "github.com/gorilla/mux"
)

func main() {
	fmt.Println("hi there..")
}

// Adding dependencies
// import "github.com/gorilla/mux"

// Update a dependency
// go get -u github.com/gorilla/mux

// clean un-used modueles + adds missing ones
// go mod tidy

/*

| Tool               | Purpose                                      |
| ------------------ | -------------------------------------------- |
| `go fmt`           | Format code (standard style)                 |
| `go vet`           | Analyze code for bugs or suspicious patterns |
| `golint`           | Suggest improvements (naming, comments)      |
| `go test`          | Run unit tests                               |
| `go test -bench=.` | Run benchmarks                               |
| `go mod init`      | Create Go module for dependency management   |
| `go mod tidy`      | Clean up dependencies                        |

*/