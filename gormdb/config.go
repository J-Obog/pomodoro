package gormdb

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("pomogoro.db"))
	
	if err != nil {
		panic(err.Error())
	}

	return db
}