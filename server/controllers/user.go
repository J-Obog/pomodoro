package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/J-Obog/pomodoro/apputils"
	"github.com/J-Obog/pomodoro/data"
	"github.com/J-Obog/pomodoro/models"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
	jti := apputils.GetTokenJTI(r)
	var user models.User

	if e := data.DB.First(&user, jti).Error; e != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": user.ID,
		"email": user.Email,
		"joined_at": user.JoinedAt,
	})
}

func GetUserMetrics(w http.ResponseWriter, r *http.Request) {
	jti := apputils.GetTokenJTI(r)
	var tasks []models.Task
	totalTasks := 0
	tasksCompleted := 0
	taskCompletionRate := 0.0

	if e := data.DB.Where(&models.Task{UserID: jti}).Find(&tasks).Error; e != nil {
		w.WriteHeader(500) 
		return 
	} else {
		totalTasks = len(tasks)
	}

	for _, v := range tasks {
		if(v.CompletedAt != nil) {
			tasksCompleted++
		}
	}

	if(totalTasks != 0) {
		taskCompletionRate = (float64(tasksCompleted) / float64(totalTasks)) * 100
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"metrics": []map[string]interface{}{
			{"metric": "Total Tasks Created", "value": totalTasks},
			{"metric": "Total Tasks Completed", "value": tasksCompleted},
			{"metric": "Task Completion Rate", "value": fmt.Sprintf("%.f%%", taskCompletionRate)}, 
		},
	})
}

func GetTasksByDay(w http.ResponseWriter, r *http.Request) {
	jti := apputils.GetTokenJTI(r) 
	var tasks []models.Task
	var groups = map[string]int{}
	var intervalMapper = map[string]int{"1w": 7}
	interval := r.URL.Query().Get("interval")
	
	if _, ok := intervalMapper[interval]; !ok {
		w.WriteHeader(400) 
		return
	}

	if e := data.DB.Where(&models.Task{UserID: jti}).Find(&tasks).Error; e != nil {
		w.WriteHeader(500) 
		return
	}
	
	for _, v := range tasks {
		b := time.Now().Add(time.Duration(-24 * intervalMapper[interval]) * time.Hour)

		if(v.CompletedAt != nil && (v.CompletedAt.Equal(b) || v.CompletedAt.After(b))) {
			k := v.CompletedAt.Format("01/02/2006")

			if _, ok := groups[k]; ok {
				groups[k]++
			} else {
				groups[k] = 1
			}
		} 
	}

	json.NewEncoder(w).Encode(groups)
}