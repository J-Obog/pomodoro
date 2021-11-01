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
	db := db.Connect()
	db.Create(&task)

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
	db := db.Connect()
	db.Find(&tasks)

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
	db := db.Connect()
	// delete task with associated id
	db.First(&task, id) 
	db.Delete(&task)

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

	db := db.Connect()
	//update task with associated id
	db.First(&task,id)
	db.Model(&task).Updates(body)

	res, err := json.Marshal(task)

	if err != nil {
		w.WriteHeader(500)
		return 
	}
		
	w.WriteHeader(200)	
	w.Write(res)
}