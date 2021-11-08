package auth

import (
	"github.com/gorilla/mux"
)


func AddRoutes(r *mux.Router) {
	r.StrictSlash(true)
	//r.Use(apputil.TestMiddleware)
	r.HandleFunc("/login", LogUserIn).Methods("POST")
	r.HandleFunc("/register", RegisterNewUser).Methods("POST")
	r.HandleFunc("/logout", LogUserOut).Methods("DELETE")
}