package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	// Create a custom logger
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io.Writer
		logger.Config{
			SlowThreshold: time.Second, // Log SQL queries slower than 1 second
			LogLevel:      logger.Info, // Log all SQL queries
			Colorful:      true,        // Enable colorful output
		},
	)

	dsn := "host=**** user=* password=**** dbname=******* port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // Use the custom logger
	})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = db
	log.Println("Database connection established")
}
