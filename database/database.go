package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB 


func SetupDatabaseConnection() *sql.DB {

	// create a connection variable 
	database, err := sql.Open("mysql", "go_user:gouser1234@(127.0.0.1:3306)/go_api?parseTime=true");

	// check for database connected or not 
	if err != nil {
		log.Fatalf("Error occur %v\n", err)
	}

	testPing := database.Ping()

	if testPing != nil {
		log.Fatalf("Error connecting database")
	}

	return database;

}


func CloseDatabaseConnection(database *sql.DB) {
	if err := database.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
}