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
	//cc := context.Background()
	//handlers
	mux.HandleFunc("/search", Handlers.Search)

	//context handlers
	//contextedMux := Context.AddContext(cc, mux)
	//log.Fatal(http.ListenAndServe(":8081", contextedMux))
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
