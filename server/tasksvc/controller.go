package tasksvc

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/pomodoro/db"
	"github.com/gorilla/mux"
)


func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	//add new tasks to db
	db.DB.Create(&task)

	res, err := json.Marshal(task)
	
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(res)
}	

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	//get all tasks from db
	db.DB.Find(&tasks)

	res, err := json.Marshal(tasks)
	
	if err != nil {
		w.WriteHeader(500)
		return 
	}

	w.WriteHeader(200)
	w.Write(res)
}	

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	
	var task Task

	// delete task with associated id
	db.DB.First(&task, id) 
	db.DB.Delete(&task)

	res, err := json.Marshal(task)

	if err != nil {
		w.WriteHeader(500)
		return 
	}

	w.WriteHeader(200)	
	w.Write(res)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	
	if err != nil {
		w.WriteHeader(500)
		return 
	}

	var task Task

	//update task with associated id
	db.DB.First(&task,id)
	db.DB.Model(&task).Updates(body)

	res, err := json.Marshal(task)

	if err != nil {
		w.WriteHeader(500)
		return 
	}
		
	w.WriteHeader(200)	
	w.Write(res)
}