package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	rcache "github.com/J-Obog/pomodoro/cache"
	"github.com/J-Obog/pomodoro/db"
	"github.com/J-Obog/pomodoro/user"
	apputil "github.com/J-Obog/pomodoro/util"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)


func LogUserIn(w http.ResponseWriter, r *http.Request) {
	var user user.User

	if e := json.NewDecoder(r.Body).Decode(&user); e != nil {
		w.WriteHeader(500) 
		return
	}

	pass := user.Password

	if e := db.DB.Where("email = ?", user.Email).First(&user).Error; e != nil {
		w.WriteHeader(401) 
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Account with email does not exist"})
	} else {
		if e := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); e != nil {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "Email and password do not match"})
		} else {
			//give user token
			accessToken := apputil.CreateAuthToken(time.Now().Add(1*time.Hour).Unix(), user.ID)
			refreshToken := apputil.CreateAuthToken(time.Now().AddDate(0, 0, 30).Unix(), user.ID)

			if accessToken == "" || refreshToken == "" {
				w.WriteHeader(500)
			} else {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"access_token": accessToken,
					"refresh_token": refreshToken,
				})
			}
		}
	}
}

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var user user.User
	
	if e := json.NewDecoder(r.Body).Decode(&user); e != nil {
		w.WriteHeader(500)
		return
	}


	if e := validator.New().Struct(user); e != nil {
		w.WriteHeader(401) 
	} else {
		if hash, e := bcrypt.GenerateFromPassword([]byte(user.Password), 10); e != nil {
			w.WriteHeader(500)
		} else {
			user.Password = string(hash)
			
			if e := db.DB.Create(&user).Error; e != nil {
				w.WriteHeader(500)
			} else {
				json.NewEncoder(w).Encode(map[string]interface{}{"message": "Registration successful"})
			}
		}
	}
}

func LogUserOut(w http.ResponseWriter, r *http.Request) {
	jti := r.Context().Value("jti")

	if _, e := rcache.RS.SetEX(rcache.CTX, fmt.Sprintf("token-%s", jti), "", 1*time.Hour).Result(); e != nil {
		w.WriteHeader(500)
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Logout successful"})
	}
}