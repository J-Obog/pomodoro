package main

import (
	"log"
	"net/http"
	"os"

	"github.com/J-Obog/pomodoro/auth"
	rcache "github.com/J-Obog/pomodoro/cache"
	"github.com/J-Obog/pomodoro/db"
	"github.com/J-Obog/pomodoro/task"
	"github.com/J-Obog/pomodoro/user"
	apputils "github.com/J-Obog/pomodoro/utils"
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
		Host: os.Getenv("POSTGRES_HOST"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port: os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
	})

	//connect to cache
	rcache.Connect(&rcache.CacheConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
		Database: os.Getenv("REDIS_DB"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	//create and configure main router
	router := mux.NewRouter().StrictSlash(true)
	router.Use(apputils.RequestLogger)

	//configure sub routers
	task.AddRoutes(router.PathPrefix("/api/task").Subrouter())
	auth.AddRoutes(router.PathPrefix("/api/auth").Subrouter())
	user.AddRoutes(router.PathPrefix("/api/user").Subrouter())
	
	//spin up server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("APP_PORT"), apputils.CORS(router))) //router wrapped with CORS middleware
}