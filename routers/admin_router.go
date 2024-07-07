package routes

import (
	"github.com/RydKrm/golang_API_build/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/api/admin")
	{
		adminRoutes.GET("/", handlers.GetAdmins)
		adminRoutes.POST("/", handlers.CreateAdmin)
		// Add more admin routes here
	}
}
