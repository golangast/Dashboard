package Server

import (
	"log"
	"net/http"

	Page "github.com/golangast/Dashboard/Dashboards/Dashhandlers/Createpage"
	Handlers "github.com/golangast/Dashboard/Handlers"
	"github.com/rs/cors"
)

//starts the server
func Serv() {

	mux := http.NewServeMux()
	//handlers
	mux.HandleFunc("/search", Handlers.Search)
	mux.HandleFunc("/createpage", Page.Createpage)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
