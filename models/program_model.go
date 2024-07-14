package models

type Program struct {
	Name string `json:"name" gorm:"type:varchar(100);not null"`
	Description string `json:"description" gorm:"type:varchar(255);not null;"`
	Country []string `json:"country" gorm:"type:json;not null;"`
	CompanyId Company `json:"companyId" gorm:"constraint:OnUpdate:CASCADE,onDelete:RESTRICT; foreignKey:CompanyID;"`
	Status bool `json:"status" gorm:"type:boolean;default:true;"`
}