package main

// postgres tutor - https://www.postgresqltutorial.com/postgresql-getting-started/install-postgresql-linux/
// http.net 	  - https://pkg.go.dev/net/http#hdr-Servers

import (
	"log"
	"net/http"
	_ "net/http"

	_ "github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)



type server struct {
	db 		*datascore
	router 	*httprouter.Router
}



func main() {
	log.Println("Listen on port 8000")

	server := &server{}
	server.setupDB()
	server.setupRoutes()

	log.Fatal(http.ListenAndServe(":8000",server.router))
	}



