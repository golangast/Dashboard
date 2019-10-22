package main

import (
	"fmt"
	"log"
	Pages "github.com/golangast/Dashboard/Page"
	Users "github.com/golangast/Dashboard/User"
	Dashboards "github.com/golangast/Dashboard/Dashboards"
)
func main() {

	fmt.Println("....starting")
	

	P, err := Pages.CreatePage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(P)

	U, err := Users.CreateUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(U)

	D, err := Dashboards.CreateDashboards()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(D)
}

