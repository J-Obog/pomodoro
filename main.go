package main

import (
	"log"
	"net/http"

	"github.com/J-Obog/pomoGOro/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	//configure api task routes
	api.AddRoutes(router.PathPrefix("/api/tasks").Subrouter())

	//spin up server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}