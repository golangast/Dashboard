package main

import (
	"fmt"
	"log"
	

)

func main() {

	fmt.Println("....starting")
	Dashboard()

	P, err := Page.CreatePage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(P)

	U, err := User.CreateUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(U)
}
func Dashboard() {
	fmt.Println("this is dashboard")
}
