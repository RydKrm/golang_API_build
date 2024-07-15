package admin_routers

import (
	"github.com/RydKrm/golang_API_build/handlers/admin_controller"
	"github.com/gin-gonic/gin"
)

func AdminProfileRouter(router *gin.Engine) {
    profile := router.Group("/api/admin")

    profile.POST("/register", admin_controller.Register)
    profile.POST("/login", admin_controller.Login)
    profile.PATCH("/update/:id", admin_controller.UpdateProfile)
    profile.PATCH("/updatePassword/:id", admin_controller.UpdatePassword)
    profile.PATCH("/updateStatus/:id", admin_controller.UpdateStatus)
    profile.GET("/single/:id", admin_controller.GetSingleAdmin)
    profile.GET("/allAdmin", admin_controller.GetAllAdmin)

}
