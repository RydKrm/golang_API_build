package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB 


func SetupDatabaseConnection() *sql.DB {

	// create a connection variable 
	db, err := sql.Open("mysql", "go_user:gouser1234@(127.0.0.1:3306)/go_api?parseTime=true");

	// check for database connected or not 
	if err != nil {
		log.Fatalf("Error occur %v\n", err)
	}

	testPing := db.Ping()

	if testPing != nil {
		log.Fatalf("Error connecting database")
	}

	return db;

}


func CloseDatabaseConnection(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
}