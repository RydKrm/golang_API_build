package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	Password string `json:"password" gorm:"type:varchar(255);not null"`      
	UserID   uint   `json:"userId" gorm:"not null"`    
	Role     string `json:"role" gorm:"type:varchar(50);not null"`         
}

// Define the roles as a constant for better code readability and maintainability
const (
	AdminRole     = "admin"
	ManagerRole   = "manager"
	CounselorRole = "counselor"
)

func (p *Password) BeforeSave(tx *gorm.DB) (err error) {
	validRoles := []string{AdminRole, ManagerRole, CounselorRole}
	isValidRole := false
	for _, role := range validRoles {
		if p.Role == role {
			isValidRole = true
			break
		}
	}
	if !isValidRole {
		return fmt.Errorf("invalid role: %s", p.Role)
	}
	return
}
