package tasksvc

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/pomodoro/db"
	"github.com/gorilla/mux"
)


func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if e := json.NewDecoder(r.Body).Decode(&task); e != nil {
		w.WriteHeader(500)
		return
	}

	//add new tasks to db
	if e := db.DB.Create(&task).Error; e != nil {
		w.WriteHeader(500)
		return
	}

	if res, e := json.Marshal(task); e != nil {
		w.WriteHeader(500)
		return
	} else {
		w.Write(res)
	}
}	

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	//get all tasks from db
	if e := db.DB.Find(&tasks).Error; e != nil {
		w.WriteHeader(500)
		return
	}

	if res, e := json.Marshal(tasks); e != nil {
		w.WriteHeader(500)
		return
	} else {
		w.Write(res)
	}
}	

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	
	var task Task

	// delete task with associated id
	if e := db.DB.First(&task, id).Delete(&task).Error; e != nil {
		w.WriteHeader(500)
		return 
	}

	if res, e := json.Marshal(task); e != nil {
		w.WriteHeader(500)
		return 
	} else {
		w.Write(res)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var body map[string]interface{}

	if e := json.NewDecoder(r.Body).Decode(&body); e != nil {
		w.WriteHeader(500)
		return 
	}

	var task Task

	//update task with associated id
	if e := db.DB.First(&task,id).Updates(body).Error; e != nil {
		w.WriteHeader(500)
		return
	}

	if res, e := json.Marshal(task); e != nil {
		w.WriteHeader(500)
		return 
	} else {
		w.Write(res)
	}
}