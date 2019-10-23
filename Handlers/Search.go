package Search

import (
	"encoding/json"
	"fmt"
	"net/http"

	DBs "github.com/golangast/Dashboard/Handlers/DB"
	Handlers "github.com/golangast/Dashboard/Handlers/Headers"
)

//searching db.
func Search(w http.ResponseWriter, r *http.Request) {
	//user for data.
	type User struct {
		Id       int    `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
	}

	var user []User
	Handlers.Header(w, r)
	rows := DBs.DB()

	var id int
	var username string
	var i = 0
	//searching rows
	for rows.Next() {

		err := rows.Scan(&id, &username)
		if err != nil {
			fmt.Println(err)
		} else {
			i++
			fmt.Println("scan ", i)
		}

		user := append(user, User{Id: id, Username: username})
		fmt.Println("before marshal ", user)
		json.NewEncoder(w).Encode(user)

	}
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")

	return
}
