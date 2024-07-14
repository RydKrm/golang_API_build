package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OffDay struct{
	gorm.Model
	CompanyID uint `json:"companyID"` 
	Company Company `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"` 
	CounselorID uint `json:"counselorID"`
	Counselor `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	Date time.Time `json:"date" gorm:"not null"`;
	Occasion string `json:"occasion" gorm:"type:varchar(255);not null"`
	OffDayCreatedBY string `json:"offDayCreatedBy"`
	OffDayCreatorRole string `json:"offDayCreatorRole"`
}

func (o *OffDay) BeforeSave(tx *gorm.DB) (err error) {
    o.Date = time.Date(o.Date.Year(), o.Date.Month(), o.Date.Day(), 0, 0, 0, 0, time.UTC)
    return
}