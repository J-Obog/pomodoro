package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/J-Obog/pomodoro/db"
	"github.com/J-Obog/pomodoro/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func CreateAuthToken(isRefresh bool, sub uint) string {
	var exp int64

	if isRefresh {
		exp = time.Now().AddDate(0, 0, 30).Unix() //refresh expires in 30 days
	} else {
		exp = time.Now().Add(60 * time.Minute).Unix() //access expires in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id: fmt.Sprint(sub),
		ExpiresAt: exp,
		IssuedAt: time.Now().Unix(),
	})

	tokenStr, e := token.SignedString([]byte("Sssshhh!!!"))
	
	if e != nil { return "" }
	return tokenStr
}


func LogUserIn(w http.ResponseWriter, r *http.Request) {
	var user user.User

	if e := json.NewDecoder(r.Body).Decode(&user); e != nil {
		w.WriteHeader(500) 
		return
	}

	pass := user.Password

	// check if there is a user with login email
	if e := db.DB.Where("email = ?", user.Email).First(&user).Error; e != nil {
		w.WriteHeader(401) 
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Account with email does not exist"})
		return
	}

	// check if user password matches with login password
	if e := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); e != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Email and password do not match"})
	} else {
		//give user token
		accessToken := CreateAuthToken(false, user.ID)
		refreshToken := CreateAuthToken(true, user.ID)

		if accessToken == "" || refreshToken == "" {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200) 
			json.NewEncoder(w).Encode(map[string]interface{}{
				"access_token": accessToken,
				"refresh_token": refreshToken,
			})
		}
	}
}

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var user user.User
	
	if e := json.NewDecoder(r.Body).Decode(&user); e != nil {
		w.WriteHeader(500)
		return
	}

	// validate user input
	if e := validator.New().Struct(user); e != nil {
		w.WriteHeader(401) 
		return
	}
	
	// hash user password
	if hash, e := bcrypt.GenerateFromPassword([]byte(user.Password), 10); e != nil {
		w.WriteHeader(500)
		return
	} else {
		user.Password = string(hash)
	}

	// create and add user to database
	if e := db.DB.Create(&user).Error; e != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Registration successful"})
	}
}

func LogUserOut(w http.ResponseWriter, r *http.Request) {

}