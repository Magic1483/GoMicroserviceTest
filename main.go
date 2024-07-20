package main

// postgres tutor - https://www.postgresqltutorial.com/postgresql-getting-started/install-postgresql-linux/
// http.net 	  - https://pkg.go.dev/net/http#hdr-Servers
// https://www.sohamkamani.com/golang/working-with-kafka/

import (
	"context"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)



type server struct {
	db 		*datascore
	router 	*httprouter.Router
	kafka_ctx *context.Context
}



func main() {
	log.Println("Listen on port 8000")

	server := &server{}
	server.setupDB()
	server.setupRoutes()
	server.SetupKafka()

	log.Fatal(http.ListenAndServe(":8000",server.router))
	}



