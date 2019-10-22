package main

import (
	"fmt"
	"log"
	. "github.com/golangast/Dashboard/Page"
	. "github.com/golangast/Dashboard/User"
)
func main() {

	fmt.Println("....starting")
	Dashboard()

	P, err := CreatePage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(P)

	U, err := CreateUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(U)
}
func Dashboard() {
	fmt.Println("this is dashboard")
}
