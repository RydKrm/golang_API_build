package models

import "github.com/jinzhu/gorm"


type MessageTemplate struct {
	Initial struct {
		MessageBody string `json:"messageBody" gorm:"column:message_body;not null"`
	} `json:"initial"`

	Repeated []struct {
		MessageBody string `json:"messageBody" gorm:"column:message_body;not null"`
		HourBefore string `json:"hour_before" gorm:"column:hour_before;not null"` 
	} `json:"repeated"`

	OnTheDay struct {
		MessageBody string `json:"messageBody" gorm:"column:message_body;not null"`
		MessageSentTime string `json:"messageSentTime" gorm:"column:message_sent_time;not null"`
	}
}

type CampaignData struct {
	DataFieldName string `json:"dataFieldName" gorm:"column:data_field_name;not null"`
	FieldType string `json:"fieldType" gorm:"column:fieldType;not null"`
	TagType string `json:"tagType" gorm:"column:tag_type;not null"`
	FieldId string `json:"fieldId" gorm:"column:field_id;"`
	Value string `json:"value" gorm:"column:value"`
	Max string `json:"max" gorm:"column:max"`
	Min string `json:"min" gorm:"column:min"`
}

type Campaign struct {
	gorm.Model

	CampaignType string `json:"campaignType" gorm:"column:campaign_type;not null"`
	CampaignName string `json:"campaignName" gorm:"column:campaign_name;not null"`
	// Company uint `json:"company" gorm:"column:company;not null"`;
	StartAppointmentAfter string `json:"startAppointmentAfter" gorm:"start_appointment_after;not null;default:1"`
	MessageTemplate MessageTemplate `json:"messageTemplate" gorm:"embedded;embeddedPrefix:message_template_"`
	CampaignData []CampaignData `json:"campaignData" gorm:"foreignKey:CampaignID"`

	Program Program `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:ProgramID"`
    Company Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CompanyID"`
}