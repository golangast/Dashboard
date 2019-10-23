package Server

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	Handlers "github.com/golangast/Dashboard/Handlers"
)

//starts the server
func Serv() {

	mux := http.NewServeMux()
	//handlers
	mux.HandleFunc("/search", Handlers.Search)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
