package user

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/pomodoro/db"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
	jti := r.Context().Value("jti")
	var user User

	if e := db.DB.First(&user, jti).Error; e != nil {
		w.WriteHeader(500)
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": user.ID,
			"email": user.Email,
		})
	}
}
