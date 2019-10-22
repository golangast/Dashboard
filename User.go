package main

import "fmt"

type err error

//User for Dashboard.
type User struct {
	Name  string
	Email string
}

//Creates a User.
func CreateUser() (error User) {
	fmt.Println("user created")

	u := User{Name: "jill", Email: "jill@yahoo.com"}

	return u
}
