package Createpage

import (
	"encoding/json"
	"fmt"
	"net/http"

	DBs "github.com/golangast/Dashboard/Dashboards/DB"
	Headers "github.com/golangast/Dashboard/Dashboards/Headers"
)

type Page struct {
	Id     int    `json:"id,omitempty"`
	Userid int    `json:"userid,omitempty"`
	Kind   string `json:"kind,omitempty"`
	Url    int    `json:"url,omitempty"`
}

func Createpage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating page")
	//user for data.

	var page []Page
	Headers.Header(w, r)
	rows := DBs.Getall()

	var id int
	var userid int
	var kind string
	var url string
	var i = 0
	//searching rows
	for rows.Next() {

		err := rows.Scan(&id, &userid, &kind, &url)
		if err != nil {
			fmt.Println(err)
		} else {
			i++
			fmt.Println("scan ", i)
		}

		page := append(page, Page{Id: id, Userid: userid, Kind: kind, Url: url})
		fmt.Println("before marshal ", page)
		json.NewEncoder(w).Encode(page)

	}
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")

	return
}

type Pager interface {
	Creating()
}

func (p Page) Creating() {

}
