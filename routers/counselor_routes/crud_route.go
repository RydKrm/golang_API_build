package counselor_routes

import (
	counselor_crud_controller "github.com/RydKrm/golang_API_build/handlers/counselor_controller"
	"github.com/gin-gonic/gin"
)


func CounselorCrudRouter(router *gin.Engine){
	crud := router.Group("/api/counselor");

	crud.POST("/register", counselor_crud_controller.CounselorRegister)
	crud.POST("/login", counselor_crud_controller.CounselorLogin)
	crud.GET("/single/:id", counselor_crud_controller.GetSingleCounselor)
	crud.GET("/all", counselor_crud_controller.GetAllCounselor)
	crud.PATCH("/updateProfile/:id", counselor_crud_controller.CounselorUpdateProfile)
	crud.PATCH("/updatePassword/:id", counselor_crud_controller.CounselorUpdatePassword)
	crud.PATCH("/updateStatus/:id", counselor_crud_controller.CounselorUpdateStatus)
	crud.PATCH("/updateTeamLead/:id", counselor_crud_controller.CounselorUpdateTeamLead)
}