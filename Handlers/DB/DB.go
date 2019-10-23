package DB

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func DB() *sql.Rows {
	db, err := sql.Open("sqlite3", "../dash2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("open ")
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ping ")
	}
	// query
	rows, err := db.Query("SELECT * FROM dashes")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SELECT * FROM dashes ")
		return rows
	}
	return rows
}
