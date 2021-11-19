package controllers

import (
	"encoding/json"
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
	var groups = map[string]int{}
	tasks_completed := 0
	groupDays := 7

	if e := data.DB.Where(&models.Task{UserID: jti}).Find(&tasks).Error; e != nil {
		w.WriteHeader(500) 
		return 
	}

	for _, v := range tasks {
		if(v.CompletedAt != nil) {
			tasks_completed++
			
			//logic for grouping task by day
			b := time.Now().Add(time.Duration(-24 * groupDays) * time.Hour)
			if(v.CompletedAt.Equal(b) || v.CompletedAt.After(b)) {
				k := v.CompletedAt.Format("Jan-02-2006")
				if _, ok := groups[k]; ok {
					groups[k]++
				} else {
					groups[k] = 1
				}
			} 
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"metrics": []map[string]interface{}{
			{"metric": "Total Tasks Created", "value": len(tasks)},
			{"metric": "Total Tasks Completed", "value": tasks_completed},
			{"metric": "Task Completion Rate", "value": tasks_completed/len(tasks)}, 
		},
		"tasks_by_day": groups, 
	})
}
