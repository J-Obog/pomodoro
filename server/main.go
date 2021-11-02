package main

import (
	"log"
	"net/http"

	"github.com/J-Obog/pomodoro/auth"
	"github.com/J-Obog/pomodoro/db"
	"github.com/J-Obog/pomodoro/task"
	"github.com/J-Obog/pomodoro/user"
	"github.com/gorilla/mux"
)


func main() {
	//connect to database
	db.Connect()

	//create and configure main router
	router := mux.NewRouter().StrictSlash(true)
	router.Use(RequestLogger)

	//configure sub routers
	task.AddRoutes(router.PathPrefix("/api/task").Subrouter())
	auth.AddRoutes(router.PathPrefix("/api/auth").Subrouter())
	user.AddRoutes(router.PathPrefix("/api/user").Subrouter())
	
	//spin up server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", CORS(router))) //router wrapped with CORS middleware
}