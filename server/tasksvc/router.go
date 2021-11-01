package tasksvc

import (
	"github.com/gorilla/mux"
)


func AddRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", CreateNewTask).Methods("POST")
	r.HandleFunc("/tasks", GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", RemoveTask).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
}