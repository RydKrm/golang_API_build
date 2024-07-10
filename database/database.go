package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// This variable hold the database connection instance.
var DB *gorm.DB

func SetupDatabaseConnection() {
	dsn := "go_user:gouser1234@(127.0.0.1:3306)/go_api?parseTime=true"
	var err error
	// create a connection with database 
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error occurred: %v\n", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}

func CloseDatabaseConnection() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
}
