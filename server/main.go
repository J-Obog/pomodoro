package main

import (
	"log"
	"net/http"

	"github.com/J-Obog/pomodoro/mware"
	"github.com/J-Obog/pomodoro/tasksvc"
	"github.com/gorilla/mux"
)


func main() {
	//create and configure main router
	router := mux.NewRouter().StrictSlash(true)
	router.Use(mware.ReqLogger)

	//configure api task routes
	tasksvc.AddRoutes(router.PathPrefix("/api").Subrouter())
	
	//spin up server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", mware.CORS(router))) //router wrapped with CORS middleware
}