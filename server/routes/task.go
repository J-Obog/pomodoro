package routes

import (
	"github.com/J-Obog/pomodoro/apputils"
	"github.com/J-Obog/pomodoro/controllers"
	"github.com/gorilla/mux"
)


func InitTaskRouter(r *mux.Router) {
	r.StrictSlash(true)
	r.Use(apputils.JWTAuthMiddleware)
	
	r.HandleFunc("/", controllers.CreateNewTask).Methods("POST")
	r.HandleFunc("/", controllers.GetAllTasks).Methods("GET")
	r.HandleFunc("/{id}", controllers.RemoveTask).Methods("DELETE")
	r.HandleFunc("/{id}", controllers.UpdateTask).Methods("PUT")
}