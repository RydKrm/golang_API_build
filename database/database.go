package database

import (
	"log"

	"github.com/RydKrm/golang_API_build/models"
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
		log.Fatalf("Error occurred while connecting to the database: %v\n", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Auto-migrate the schema
	// database schema is updated automatically based on your Go struct definitions 
	// Auto-migrate models
	err = DB.AutoMigrate(&models.Admin{}, &models.Company{}, &models.Counselor{}, &models.Program{}, &models.Manager{})
	if err != nil {
		log.Fatalf("Error auto-migrating schema: %v\n", err)
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

	log.Println("Database connection closed")
}
