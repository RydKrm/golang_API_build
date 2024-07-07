package models

import (
	"database/sql"

	"github.com/RydKrm/golang_API_build/config"
)

var db *sql.DB

func init() {
	db = config.SetupDatabaseConnection()
}

func GetDB() *sql.DB {
	return db
}
