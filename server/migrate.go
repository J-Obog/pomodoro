package main

import (
	"github.com/J-Obog/pomoGOro/api"
	"github.com/J-Obog/pomoGOro/gormdb"
)

func MakeMigrations() {
	db := gormdb.Connect()
	db.AutoMigrate(&api.Task{})
}