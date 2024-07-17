package manager_routes

import (
	manager_crud_controller "github.com/RydKrm/golang_API_build/handlers/manager_controller"
	"github.com/gin-gonic/gin"
)

func ManagerCrudRouter(router *gin.Engine){
	crud := router.Group("/api/manager");

	crud.POST("/register", manager_crud_controller.ManagerRegister);
	crud.POST("/login", manager_crud_controller.ManagerLogin)
	crud.GET("/single/:id", manager_crud_controller.GetSingleManager)
	crud.GET("/all", manager_crud_controller.GetAllManager)
	crud.PATCH("/updateProfile/:id", manager_crud_controller.ManagerUpdateProfile)
	crud.PATCH("/updateStatus/:id", manager_crud_controller.ManagerUpdateStatus)

}