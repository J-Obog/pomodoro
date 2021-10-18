package api

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/pomoGOro/gormdb"
)


func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	//add new tasks to db
	db := gormdb.Connect()
	db.Model(&Task{}).Create(&task)

	w.WriteHeader(200)
	w.Write([]byte("Task successfully created"))
}	

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	db := gormdb.Connect()
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
	
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	
}