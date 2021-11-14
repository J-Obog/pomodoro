package task

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/J-Obog/pomodoro/db"
	"github.com/gorilla/mux"
)


func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	jti := r.Context().Value("jti").(string)
	uid, _ := strconv.ParseUint(jti, 10, 64) 

	var task Task

	if e := json.NewDecoder(r.Body).Decode(&task); e != nil {
		w.WriteHeader(500)
	} else {
		task.UserID = uid

		if e := db.DB.Create(&task).Error; e != nil {
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
	var tasks []Task

	if e := db.DB.Find(&tasks, "user_id = ?", jti).Error; e != nil {
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
	var task Task

	if e := db.DB.Delete(&task, id).Error; e != nil {
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
	var task Task

	if e := json.NewDecoder(r.Body).Decode(&body); e != nil {
		w.WriteHeader(500)
		return 
	} else {
		if e := db.DB.First(&task, id).Updates(body).Error; e != nil {
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