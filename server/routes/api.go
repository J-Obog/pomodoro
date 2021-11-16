package routes

import (
	"github.com/gorilla/mux"
)


func InitApiRouter(r *mux.Router) {
	r.StrictSlash(true)
	InitAuthRouter(r.PathPrefix("/auth").Subrouter())
	InitTaskRouter(r.PathPrefix("/task").Subrouter())
	InitUserRouter(r.PathPrefix("/user").Subrouter())
}