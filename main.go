package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RydKrm/golang_API_build/database"
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

	database.SetupDatabaseConnection()
	defer database.CloseDatabaseConnection()

	// if os.Getenv("GIN_MODE") != "debug" {
    // 	gin.SetMode(gin.ReleaseMode)
	// }



	router := gin.Default()  

	// how to define setup a router

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"status":true, "message":"Server working"})
	})

	// routes.RegisterAdminRoutes(router)

	// adminRouter := router.Group("/api/admin")
	// {
	// 	registerAdminRoutes(adminRouter)
	// }

	// RegisterAdminRouter(router.Group("/api/admin"))
	
	


	port := os.Getenv("PORT")

	if port == ""{
		port = "5000"
	}

	if err := router.Run(":"+port); err != nil{
		log.Fatalf("Failed to run server : %v \n", err)
	}

}