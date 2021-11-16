package models

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}