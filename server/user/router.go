package user

import (
	apputil "github.com/J-Obog/pomodoro/util"
	"github.com/gorilla/mux"
)


func AddRoutes(r *mux.Router) {
	r.StrictSlash(true)
	r.Use(apputil.JwtRequired)

	r.HandleFunc("/", GetUser).Methods("GET")
}