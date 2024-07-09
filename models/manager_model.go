package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Manager struct {
	ID uint `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Email string `gorm:"type:varchar(255);not null" json:"email"`
	Password string `gorm:"type:varchar(100):not null" json:"-"`
	Role string `gorm:"type:varchar(20);default:'manager'" json:"role"`
	Status string `gorm:"type:varchar(20);deafult:true" json:"status"`
	Company Company `gorm:"constraint:OnUpdate:CASECADE,onDelete:SET NULL;" json:"company"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (m *Manager) BeforeSave()(err error){
	m.Password,err = hashedPassword(m.Password)
	if err!= nil {
		return err;
	}
	return nil
}

func hashedPassword(password string)(string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err!=nil{
		return "",err
	}
	return string(hash),nil
}

func (m *Manager) CheckPassword(password string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(m.Password),[]byte (password))
	return err == nil
}

func (m *Manager) GenerateToken()(string, error){
	claims := jwt.MapClaims{}
	claims["id"] = m.ID
	claims["role"] = m.Role
	claims["exp"] = time.Now().Add(time.Hour*72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

}