package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", 
						os.Getenv("POSTGRES_HOST"), 
						os.Getenv("POSTGRES_USER"), 
						os.Getenv("POSTGRES_PASSWORD"), 
						os.Getenv("POSTGRES_PORT"), 
						os.Getenv("POSTGRES_DB"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to make connection to database")
	}

	DB = db
}