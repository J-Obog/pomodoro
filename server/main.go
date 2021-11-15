package main

import (
	"log"
	"net/http"
	"os"

	"github.com/J-Obog/pomodoro/auth"
	"github.com/J-Obog/pomodoro/db"
	"github.com/J-Obog/pomodoro/task"
	"github.com/J-Obog/pomodoro/user"
	apputil "github.com/J-Obog/pomodoro/util"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)



func main() {
	if os.Getenv("GO_ENV") == "dev" {
		if e := godotenv.Load(".env"); e != nil {
			log.Fatal("Failed to initialize dotenv configuration")
		}
	}
	
	//connect to postgres db and redis cache
	db.Connect()
	//rcache.Connect()

	//create and configure main router
	router := mux.NewRouter().StrictSlash(true)
	router.Use(apputil.RequestLogger)

	//configure sub routers
	task.AddRoutes(router.PathPrefix("/api/task").Subrouter())
	auth.AddRoutes(router.PathPrefix("/api/auth").Subrouter())
	user.AddRoutes(router.PathPrefix("/api/user").Subrouter())
	
	//spin up server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("APP_PORT"), apputil.CORS(router))) //router wrapped with CORS middleware
}