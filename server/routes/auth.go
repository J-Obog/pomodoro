package routes

import (
	"github.com/J-Obog/pomodoro/controllers"
	"github.com/gorilla/mux"
)


func InitAuthRouter(r *mux.Router) {
	r.StrictSlash(true)
	//r.Use(apputil.TestMiddleware)
	r.HandleFunc("/login", controllers.LogUserIn).Methods("POST")
	r.HandleFunc("/register", controllers.RegisterNewUser).Methods("POST")
	r.HandleFunc("/logout", controllers.LogUserOut).Methods("DELETE")
}