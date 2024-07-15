package company_routes

import (
	"github.com/RydKrm/golang_API_build/handlers/company_controller"
	"github.com/gin-gonic/gin"
)

func CompanyCrudRouter(router *gin.Engine){
	crud := router.Group("/api/company")

	crud.POST("/create", company_controller.CreateCompany)


}