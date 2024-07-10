package comapanycontroller

import (
	"net/http"

	"github.com/RydKrm/golang_API_build/database"
	"github.com/RydKrm/golang_API_build/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CompanyCreateRequest struct{
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status bool `json:"status" binding:"-"`
}

func CreateCompany(c *gin.Context){
	var company CompanyCreateRequest

	if err := c.ShouldBindJSON(&company); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"status":false,"message":"Company name and description required"})
		return;
	}

	if err := database.DB.Create(&company).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to create company"})
	}  else {
		c.JSON(http.StatusCreated, gin.H{"status":true, "message":"Company created"})
	}

}

type CompanyUpdateRequest struct{
	Name string `json:"name" binding:"-"`
	Description string `json:"description" binding:"-"`
}

func UpdateCompany(c *gin.Context){
	var company CompanyUpdateRequest

	if err := c.ShouldBindJSON(&company); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"status":false,"message":"Field data required"})
		return
	}

	fieldId := c.Param("fieldId")

	if err := database.DB.Where("ID = ?",fieldId).Updates(&company); err == nil{
		c.JSON(http.StatusOK, gin.H{"status":true,"message":"Company updated"})
	} else {
		c.JSON(http.StatusNotModified, gin.H{"status":false,"message":"Company not updated"})
	}

}

func GetSingleCompany(c *gin.Context){
	var company models.Company

	// get the company 
	// convert this to json 
	if err := c.ShouldBindJSON(&company); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"status":false, "message":"Company Data required"})
		return;
	}

	// get the field id from params 
	fieldId := c.Param("fieldId")

	// fetch the data from database

	if err := database.DB.First(&company,fieldId).Error; err!=nil{
		if err== gorm.ErrRecordNotFound {
			c.JSON(400,gin.H{"status":false,"message":"company not found"})
		} else {
			c.JSON(http.StatusInternalServerError,gin.H{"status":false, "message":"Error on fatching data"})
		}
		return
	}

	c.JSON(200,gin.H{"status":true, "message":"Single company found", "company":company})
}