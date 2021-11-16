package main

import (
	"log"
	"net/http"
	"os"

	"github.com/J-Obog/pomodoro/apputils"
	"github.com/J-Obog/pomodoro/data"
	"github.com/J-Obog/pomodoro/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func main() {
	if os.Getenv("GO_ENV") == "dev" {
		if e := godotenv.Load(".env"); e != nil {
			log.Fatal("Failed to initialize dotenv configuration")
		}
	}
	
	data.ConnectDB()
	data.ConnectCache()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(apputils.LoggerMiddleware)
	routes.InitApiRouter(router.PathPrefix("/api").Subrouter())

	log.Println("Server running on port " + os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("APP_PORT"), apputils.CORSMiddleware(router)))
}