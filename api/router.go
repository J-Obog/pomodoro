package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("<h1>Bar</h1>"))
	}).Methods("GET")
}