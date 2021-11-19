package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/J-Obog/pomodoro/apputils"
	"github.com/J-Obog/pomodoro/data"
	"github.com/J-Obog/pomodoro/models"
	"github.com/J-Obog/pomodoro/validators"
	"github.com/gorilla/mux"
)


func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	jti := apputils.GetTokenJTI(r)
	body, e := apputils.ParseBody(r)

	if e != nil {
		w.WriteHeader(500)
		return
	}

	if e := validators.ValidateNewTask(body); e != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]interface{}{"errors": e})
		return
	}

	var task = models.Task{ 
		UserID: jti, 
		Title: body["title"].(string),
		CreatedAt: time.Now(),
	}

	if e := data.DB.Create(&task).Error; e != nil {
		w.WriteHeader(500)
		return
	}
		
	json.NewEncoder(w).Encode(task)
}	

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	jti := apputils.GetTokenJTI(r)
	var tasks []models.Task

	if e := data.DB.Find(&tasks, "user_id = ?", jti).Error; e != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}	

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	id, e := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if e != nil {
		w.WriteHeader(402)
		return
	}

	if e := data.DB.Delete(&models.Task{}, id).Error; e != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	body, e := apputils.ParseBody(r)
	id := mux.Vars(r)["id"]
	var task models.Task

	if e != nil {
		w.WriteHeader(500)
		return
	}

	if e := data.DB.First(&task, id).Updates(body).Error; e != nil {
		w.WriteHeader(500)
		return
	} 
	
	json.NewEncoder(w).Encode(task)
}