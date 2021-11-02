package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Obog/pomodoro/db"
	"github.com/J-Obog/pomodoro/user"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func LogUserIn(w http.ResponseWriter, r *http.Request) {
	var user user.User

	if e := json.NewDecoder(r.Body).Decode(&user); e != nil {
		w.WriteHeader(500) 
		return
	}

	if e := db.DB.Where("email = ?", user.Email).First(&user).Error; e != nil {
		w.WriteHeader(401) 
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Account with email does not exist"})
		return
	}
}

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var user user.User

	if e := validator.New().Struct(user); e != nil {
		w.WriteHeader(401) 
		fmt.Println(e)
		return
	}

	if e := json.NewDecoder(r.Body).Decode(&user); e != nil {
		w.WriteHeader(500)
		return
	}
	
	if hash, e := bcrypt.GenerateFromPassword([]byte(user.Password), 10); e != nil {
		w.WriteHeader(500)
		return
	} else {
		user.Password = string(hash)
	}

	if e := db.DB.Create(&user).Error; e != nil {
		w.WriteHeader(500)
		return
	}

	if res, e := json.Marshal(user); e != nil {
		w.WriteHeader(500)
		return
	} else {
		w.Write(res)
	}
}

func LogUserOut(w http.ResponseWriter, r *http.Request) {

}