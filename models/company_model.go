package models

type Company struct{
	ID uint `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:varchar(600);not null" json:"description"`
	Status bool `gorm:"default:true" json:"status"`
}