package main

import (
	"fmt"

	Page "github.com/golangast/Dashboard"
	User "github.com/golangast/Dashboard"
)

type err error

func main() {

	fmt.Println("....starting")
	Dashboard()

	err, P := Page.CreatePage()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(P)

	err, U := User.CreateUser()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(U)
}
func Dashboard() {
	fmt.Println("this is dashboard")
}
