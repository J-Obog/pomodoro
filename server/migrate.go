package main

import (
	"github.com/J-Obog/pomodoro/api"
	"github.com/J-Obog/pomodoro/gormdb"
)

func MakeMigrations() {
	db := gormdb.Connect()
	db.AutoMigrate(&api.Task{})
}