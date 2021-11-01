package tasksvc

import (
	"github.com/gorilla/mux"
)


func AddRoutes(r *mux.Router) {
	r.HandleFunc("/", CreateNewTask).Methods("POST")
	r.HandleFunc("/", GetAllTasks).Methods("GET")
	r.HandleFunc("/{id}", RemoveTask).Methods("DELETE")
	r.HandleFunc("/{id}", UpdateTask).Methods("PUT")
}