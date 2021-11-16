package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/pomodoro/apputils"
	"github.com/J-Obog/pomodoro/data"
	"github.com/J-Obog/pomodoro/models"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
	jti := apputils.GetTokenJTI(r)
	
	var user models.User

	if e := data.DB.First(&user, jti).Error; e != nil {
		w.WriteHeader(500)
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": user.ID,
			"email": user.Email,
		})
	}
}
