package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host string
	Username string
	Password string
	Port string
	Database string
}


var DB *gorm.DB

func Connect(cfg *DBConfig) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", 
						cfg.Host, 
						cfg.Username, 
						cfg.Password, 
						cfg.Port, 
						cfg.Database)
						
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to make connection to database")
	}

	DB = db
}