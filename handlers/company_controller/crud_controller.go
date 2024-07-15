package company_controller

import (
	"github.com/RydKrm/golang_API_build/database"
	"github.com/RydKrm/golang_API_build/models"
	"github.com/RydKrm/golang_API_build/utils/response"
	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context){
	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil{
		response.NegativeResponse(c, "Company name and description field required");
		return
	}

	// Create company in the database
	if err := database.DB.Create(&company).Error; err != nil {
		response.NegativeResponse(c, "Failed to create company")
		return
	}

	response.PositiveResponse(c, "New Company created", gin.H{"company":company})

}