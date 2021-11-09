package task

import (
	apputil "github.com/J-Obog/pomodoro/util"
	"github.com/gorilla/mux"
)


func AddRoutes(r *mux.Router) {
	r.StrictSlash(true)
	r.Use(apputil.JwtRequired)
	
	r.HandleFunc("/", CreateNewTask).Methods("POST")
	r.HandleFunc("/", GetAllTasks).Methods("GET")
	r.HandleFunc("/{id}", RemoveTask).Methods("DELETE")
	r.HandleFunc("/{id}", UpdateTask).Methods("PUT")
}