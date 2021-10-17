package gormdb

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}