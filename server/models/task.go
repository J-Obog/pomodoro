package models

type Task struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed"`
}