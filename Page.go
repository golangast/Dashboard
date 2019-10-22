package main

import "fmt"

//Page is a Page
type Page struct {
	ID int
}

//Creates Page method
func CreatePage() (Page, error) {
	fmt.Println("create page")

	p := Page{ID: 2}
	return err, p
}
