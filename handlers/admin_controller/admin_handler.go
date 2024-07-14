package admin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAdmins(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message":"Get all admin"})
}

func CreateAdmin(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{"message":"Create admin"})
}

func GetAdmin(c *gin.Context){
	
}

func GetAllCounselorPasseword(c *gin.Context){}

func GetAllManagerPassword(c *gin.Context){}

func GetAllAdminPassword(c *gin.Context){}

