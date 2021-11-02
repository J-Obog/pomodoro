package main

import (
	"log"
	"net/http"
	"os"

	"github.com/J-Obog/pomodoro/auth"
	"github.com/J-Obog/pomodoro/db"
	"github.com/J-Obog/pomodoro/task"
	"github.com/J-Obog/pomodoro/user"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func main() {
	if os.Getenv("GO_ENV") != "PROD" {
		if e := godotenv.Load(".env"); e != nil {
			log.Fatal("Failed to initialize dotenv configuration")
		}
	}

	//connect to database
	db.Connect(&db.DBConfig{
		Host: os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port: os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DBNAME"),
	})

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