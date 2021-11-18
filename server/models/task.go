package models

import "time"

type Task struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	Title     string `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	CompletedAt time.Time `json:"complete_at"`
}