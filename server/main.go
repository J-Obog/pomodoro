package main

import (
	"log"
	"net/http"

	"github.com/J-Obog/pomodoro/authsvc"
	"github.com/J-Obog/pomodoro/db"
	"github.com/J-Obog/pomodoro/tasksvc"
	"github.com/J-Obog/pomodoro/usersvc"
	"github.com/gorilla/mux"
)


func main() {
	//connect to database
	db.Connect()

	//create and configure main router
	router := mux.NewRouter().StrictSlash(true)
	router.Use(RequestLogger)

	//configure sub routers
	tasksvc.AddRoutes(router.PathPrefix("/api/task").Subrouter())
	authsvc.AddRoutes(router.PathPrefix("/api/auth").Subrouter())
	usersvc.AddRoutes(router.PathPrefix("/api/user").Subrouter())
	
	//spin up server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", CORS(router))) //router wrapped with CORS middleware
}