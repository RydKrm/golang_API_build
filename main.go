package main

import (
	"log"
	"os"

	"github.com/RydKrm/golang_API_build/config"
	routes "github.com/RydKrm/golang_API_build/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// first load the env file
	err := godotenv.Load();

	// check if this file is loaded or not 
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database := config.SetupDatabaseConnection()

	defer config.CloseDatabaseConnection(database)

	// if os.Getenv("GIN_MODE") != "debug" {
    // 	gin.SetMode(gin.ReleaseMode)
	// }


	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	routes.RegisterAdminRoutes(router)


	port := os.Getenv("PORT")

	if port == ""{
		port = "5000"
	}

	if err := router.Run(":"+port); err != nil{
		log.Fatalf("Failed to run server : %v \n", err)
	}

}