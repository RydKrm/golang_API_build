package handlers

import (
	"net/http"

	"github.com/RydKrm/golang_API_build/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AdminInput struct {
	Name        string `json:"name" building:"required"`
	Email       string `json:"email" building:"required"`
	PhoneNumber string `json:"phoneNumber" building:"required"`
}

func AdminRegister(c *gin.Context){
	var input AdminInput;

	// check that in coming data is json or not 
	// if not json consider it as a new view 

	if err := c.ShouldBindBodyWithJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}

	db := c.MustGet("db").(*gorm.DB)

	var existingAdmin models.Admin 

	if db.Where("email = ? OR phoneNumber = ?",input.Email, input.PhoneNumber).First(&existingAdmin).RecordNotFound(){
		password := "123456";
		admin := models.Admin{
			Name : input.Name,
			Email : input.Email,
			PhoneNumber: input.PhoneNumber,
			Password: password,
		}
		db.Create(&admin)
		c.JSON(http.StatusCreated, gin.H{"success":true, "message":"Admin created success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"success":false, "message":"Admin not created"})
	}

}

func AdminLogin(c *gin.Context){
	var input struct{
		Email string `json:"email" building:"required"`
		Password string `json:"password" building:"required"`
	}

	if err:= c.ShouldBindBodyWithJSON(&input); err !=nil{
		c.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	var admin models.Admin

	if db.Where("email = ?", input.Email).First(&admin).RecordNotFound(){
		c.JSON(http.StatusNotFound, gin.H{"status":false,"message":"Admin not found"});
		return
	}

	if !admin.CheckPassword(input.Password){
		c.JSON(http.StatusForbidden, gin.H{"status":false,"message":"Password did not match"})
		return
	}

}