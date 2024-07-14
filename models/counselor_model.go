package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Counselor struct{
	gorm.Model
	Name string `json:"name" gorm:"column:name;trim;not null"`;
	Email string `json:"email" gorm:"column:email;unique;not null;"`
	Password string `json:"-" gorm:"column:email;not null"`
	PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(20);not null"`
	ProgramList []Program `json:"programList" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ProgramID"`
	WhatsApp string `json:"whatsApp" gorm:"column:whats_app;trim"`
	Status bool `json:"status" gorm:"default:true"`
	TeamLead bool `json:"teamLead" gorm:"column:team_lead;default:false"`
	Role string `json:"role" gorm:"column:role;default:counselor"`
	Online bool `json:"online" gorm:"column:online;default:false"`
	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL foreignKey:CompanyID"`
	Manager Manager `json:"manager" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL foreignKey:ManagerID"`
}

func (counselor *Counselor) BeforeSave()(err error){
	counselor.Password, err = hashPassword(counselor.Password)
	if err != nil{
		return err;
	} 
	return nil;
}

func (counselor *Counselor) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(counselor.Password),[]byte(password))
	return err == nil;
}

func (counselor *Counselor) GenerateToken(password string) (string, error){
	claims := jwt.MapClaims{}
	claims["id"] = counselor.ID;
	claims["role"] = counselor.Role;
	claims["exp"] = time.Now().Add(time.Hour*24).Unix();

	token := jwt.NewWithClaims(jwt.SigningMethodES256,claims);
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}