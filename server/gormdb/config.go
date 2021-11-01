package gormdb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=root port=5432 dbname=gorm"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err.Error())
	}

	return db
}