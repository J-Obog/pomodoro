package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/J-Obog/pomodoro/apputils"
	"github.com/J-Obog/pomodoro/data"
	"github.com/J-Obog/pomodoro/models"
	"github.com/J-Obog/pomodoro/validators"
	"golang.org/x/crypto/bcrypt"
)


func LogUserIn(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if e := json.NewDecoder(r.Body).Decode(&user); e != nil {
		w.WriteHeader(500) 
		return
	}

	pass := user.Password

	if e := data.DB.Where("email = ?", user.Email).First(&user).Error; e != nil {
		w.WriteHeader(401) 
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Account with email does not exist"})
		return
	} 

	if e := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); e != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Email and password do not match"})
		return
	} 

	accessToken := apputils.CreateAuthToken(1, user.ID, "access")
	refreshToken := apputils.CreateAuthToken(24, user.ID, "refresh")

	if accessToken == "" || refreshToken == "" {
		w.WriteHeader(500)
		return
	}
	
	json.NewEncoder(w).Encode(map[string]interface{}{"access_token": accessToken,"refresh_token": refreshToken})
}

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	body, e := apputils.ParseBody(r)
	
	if e != nil {
		w.WriteHeader(500)
		return
	}

	if e := validators.ValidateUserReg(body); e != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]interface{}{"errors": e})
		return
	}

	if hash, e := bcrypt.GenerateFromPassword([]byte(body["password"].(string)), 10); e != nil {
		w.WriteHeader(500)
	} else {
		var user = models.User{ Email: body["email"].(string), Password: string(hash), }

		if e := data.DB.Create(&user).Error; e != nil {
			w.WriteHeader(500)
			return
		}
		
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Registration successful"})
	}
}

func LogUserOut(w http.ResponseWriter, r *http.Request) {
	jwtStr := apputils.GetTokenRaw(r)
	k := fmt.Sprintf("token.%s", jwtStr)

	if _, e := data.RS.SetEX(context.Background(), k, "", 24*time.Hour).Result(); e != nil {
		w.WriteHeader(500)
		return
	}
		
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Logout successful"})
}