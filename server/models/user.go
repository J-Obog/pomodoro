package models

import "time"

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	JoinedAt time.Time `json:"joined_at"`
}