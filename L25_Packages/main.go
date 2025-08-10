package main

import (
	"fmt"

	"github.com/adocxwork/GoLang/auth"
	"github.com/adocxwork/GoLang/user"
)

// go mod init <package name>
// go get <external package name>
// go mod tidy -> fixes many things.., installs packages which are being used but not being installed

func main() {
	auth.LoginWithCredentials("adocxwork", "secret")
	session := auth.GetSession()
	fmt.Println("Session : ", session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "John Doe",
	}

	fmt.Println(user.Email, user.Name)
}
