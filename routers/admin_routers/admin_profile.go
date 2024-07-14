package admin_routers

import (
	"github.com/RydKrm/golang_API_build/handlers/admin_controller"
	"github.com/gin-gonic/gin"
)


func RegisterAdminProfileRouter(router *gin.RouterGroup){
	router.POST("/register", admin_controller.Register);
	router.POST("/login", admin_controller.Login)
}