package models

type DataField struct{
	Name string `json:"name" gorm:"type:varchar(50);not null;trim;"`
	Description string `json:"description" gorm:"type:varchar(255);not null;"`
	Company Company `json:"company" gorm:"constraint:OnDelete:SET NULL;OnUpdate:CASCADE;foreignKey:CompanyID"`
	DataType string `json:"dataType" gorm:"type:varchar(30);not null;"`
	Primary string `json:"primary" gorm:"default:false"`
	Secondary string `json:"secondary" gorm:"default:false"`
	NotFoundMessage string `json:"notFoundMessage" gorm:"type:varchar(255)"`
	Status bool `json:"status" gorm:"default:true"`
}