package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/RydKrm/golang_API_build/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/RydKrm/golang_API_build/database"
)

func Register(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "All fields required"})
		return
	}

	var count int64
	database.DB.Model(&models.Admin{}).Where("email = ? OR phone_number = ?", admin.Email, admin.PhoneNumber).Count(&count)
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "Admin already exists with phone number or email"})
		return
	}

	password := "generated_password"  // Generate a password as per your requirement
	admin.Password = password

	if err := database.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin created successfully"})
}

func Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email and password required"})
		return
	}

	var admin models.Admin
	if err := database.DB.Where("email = ?", request.Email).First(&admin).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Admin not found"})
		return
	}

	if !admin.ComparePassword(request.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect password"})
		return
	}

	token, err := admin.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"role":  admin.Role,
		"id":    admin.ID,
		"message": "Login success",
	})
}

func UpdateProfile(c *gin.Context) {
	id := c.Param("id")
	var admin models.Admin
	if err := database.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Admin not found"})
		return
	}

	var updateData struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "At least one field required"})
		return
	}

	if updateData.Email != "" {
		var count int64
		database.DB.Model(&models.Admin{}).Where("email = ? AND id != ?", updateData.Email, id).Count(&count)
		if count > 0 {
			c.JSON(http.StatusConflict, gin.H{"message": "Email already used"})
			return
		}
	}

	database.DB.Model(&admin).Updates(models.Admin{
		Name:        updateData.Name,
		Email:       updateData.Email,
		PhoneNumber: updateData.PhoneNumber,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
}

func UpdatePassword(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Old and new password required"})
		return
	}

	var admin models.Admin
	if err := database.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Admin not found"})
		return
	}

	if !admin.ComparePassword(request.OldPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Old password did not match"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash new password"})
		return
	}

	admin.Password = string(hashedPassword)
	if err := database.DB.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})
}
