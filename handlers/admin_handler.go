package handlers

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