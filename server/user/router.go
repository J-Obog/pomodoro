package user

import (
	"github.com/gorilla/mux"
)


func AddRoutes(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/", nil).Methods("GET")
}