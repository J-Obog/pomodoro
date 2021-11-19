package routes

import (
	"github.com/J-Obog/pomodoro/apputils"
	"github.com/J-Obog/pomodoro/controllers"
	"github.com/gorilla/mux"
)


func InitUserRouter(r *mux.Router) {
	r.StrictSlash(true)
	r.Use(apputils.JWTAuthMiddleware)

	r.HandleFunc("/", controllers.GetUser).Methods("GET")
	r.HandleFunc("/metrics", controllers.GetUserMetrics).Methods("GET")
}