package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RydKrm/golang_API_build/database"
	"github.com/RydKrm/golang_API_build/routers/admin_routers"
	"github.com/RydKrm/golang_API_build/routers/company_routes"
	"github.com/RydKrm/golang_API_build/routers/counselor_routes"
	"github.com/RydKrm/golang_API_build/routers/manager_routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// first load the env file
	// this will make the envirment variable globally access
	err := godotenv.Load();

	// check if this file is loaded or not 
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
    
	// make the database connection 
	database.SetupDatabaseConnection()
	// database connection close should be make in the end of the file
	defer database.CloseDatabaseConnection()

    // Uncomment for production to disable Gin's debug mode
	if os.Getenv("GIN_MODE") != "debug" {
    	gin.SetMode(gin.ReleaseMode)
	}
    

	// initialize the router
	router := gin.Default();

	// Middleware setup 
    router.Use(gin.Logger());
    router.Use(gin.Recovery());
    
	 // checker router to check if the server working or not
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":  true,
            "message": "Server working",
        })
    })
    
	// * start adding router
    // Admin routes
    admin_routers.AdminProfileRouter(router)
	company_routes.CompanyCrudRouter(router)
	counselor_routes.CounselorCrudRouter(router)
	manager_routes.ManagerCrudRouter(router)

	// Set trusted proxies if needed
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// find the port number from .env 
	port := os.Getenv("PORT") 

	if port == ""{
		port = "5000"
	}
    

	// finally run the server 
	if err := router.Run(":"+port); err != nil{
		log.Fatalf("Failed to run server : %v \n", err)
	}

}