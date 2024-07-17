package models

import "github.com/jinzhu/gorm"

type Program struct {
	gorm.Model
	Name        string   `json:"name" gorm:"type:varchar(100);not null"`
	Description string   `json:"description" gorm:"type:varchar(255);not null"`
	Country     []string `json:"country" gorm:"type:json;not null"`
	CompanyID   uint     `json:"companyId" gorm:"column:company_id;not null"`
	Company     Company  `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CompanyID"`
	Status      bool     `json:"status" gorm:"type:boolean;default:true"`
}
