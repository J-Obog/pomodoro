package controllers

import (
	"encoding/json"
	"net/http"
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
		
	if res, e := json.Marshal(task); e != nil {
		w.WriteHeader(500)
	} else {
		w.Write(res)
	}
}	

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	jti := apputils.GetTokenJTI(r)
	var tasks []models.Task

	if e := data.DB.Find(&tasks, "user_id = ?", jti).Error; e != nil {
		w.WriteHeader(500)
		return
	}

	if res, e := json.Marshal(tasks); e != nil {
			w.WriteHeader(500)
	} else {
			w.Write(res)
	}
}	

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]	
	var task models.Task

	if e := data.DB.Delete(&task, id).Error; e != nil {
		w.WriteHeader(500)
		return
	}

	if res, e := json.Marshal(task); e != nil {
		w.WriteHeader(500)
	} else {
		w.Write(res)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	body, e := apputils.ParseBody(r)
	var task models.Task
	id := mux.Vars(r)["id"]

	if e != nil {
		w.WriteHeader(500)
		return
	}

	if e := data.DB.First(&task, id).Updates(body).Error; e != nil {
		w.WriteHeader(500)
		return
	} 
	
	if res, e := json.Marshal(task); e != nil {
		w.WriteHeader(500) 
	} else {
		w.Write(res)
	}
}