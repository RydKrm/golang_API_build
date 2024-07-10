package models

import "github.com/jinzhu/gorm"

type Company struct{
	gorm.Model
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:varchar(600);not null" json:"description"`
	Status bool `gorm:"default:true" json:"status"`
}  