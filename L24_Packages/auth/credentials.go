package auth

import "fmt"

// func name -> starts with capital letter -> Higher scope
func LoginWithCredentials(username string, passwords string) {
	fmt.Println("Loging user :", username, passwords)
}