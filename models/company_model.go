package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name;not null"`
	Description string `json:"description" gorm:"column:description"`
	Status      bool   `json:"status" gorm:"default:true"`
}
