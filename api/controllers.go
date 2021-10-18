package api

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/pomoGOro/gormdb"
)


func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	db := gormdb.Connect()
	db.Create(&t)

	w.WriteHeader(200)
	w.Write([]byte("Task successfully created"))
}	

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	
}