package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/J-Obog/pomodoro/data"
	"github.com/J-Obog/pomodoro/models"
	"github.com/gorilla/mux"
)


func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	jti := r.Context().Value("jti").(string)
	uid, _ := strconv.ParseUint(jti, 10, 64) 

	var task models.Task

	if e := json.NewDecoder(r.Body).Decode(&task); e != nil {
		w.WriteHeader(500)
	} else {
		task.UserID = uid

		if e := data.DB.Create(&task).Error; e != nil {
			w.WriteHeader(500)
		} else {
			if res, e := json.Marshal(task); e != nil {
				w.WriteHeader(500)
			} else {
				w.Write(res)
			}
		}
	}
}	

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	jti := r.Context().Value("jti")
	var tasks []models.Task

	if e := data.DB.Find(&tasks, "user_id = ?", jti).Error; e != nil {
		w.WriteHeader(500)
	} else {
		if res, e := json.Marshal(tasks); e != nil {
			w.WriteHeader(500)
		} else {
			w.Write(res)
		}
	}
}	

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]	
	var task models.Task

	if e := data.DB.Delete(&task, id).Error; e != nil {
		w.WriteHeader(500)
	} else {
		if res, e := json.Marshal(task); e != nil {
			w.WriteHeader(500)
		} else {
			w.Write(res)
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var body map[string]interface{}
	var task models.Task

	if e := json.NewDecoder(r.Body).Decode(&body); e != nil {
		w.WriteHeader(500)
		return 
	} else {
		if e := data.DB.First(&task, id).Updates(body).Error; e != nil {
			w.WriteHeader(500)
		} else {
			if res, e := json.Marshal(task); e != nil {
				w.WriteHeader(500) 
			} else {
				w.Write(res)
			}
		}
	}
}