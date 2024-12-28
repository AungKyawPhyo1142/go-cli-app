package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbString := os.Getenv("DATABASE_URL")
	if dbString == "" {
		log.Fatal("Database url is not set.")
	}

	// connect to db
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting database: %v", err)
	}

	log.Print("Database connected successfully")

	DB = db.Debug()

}
