package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Manager struct {
	gorm.Model
	Name     string  `gorm:"type:varchar(255);not null" json:"name"`
	Email    string  `gorm:"type:varchar(255);not null" json:"email"`
	Password string  `gorm:"type:varchar(100);not null" json:"-"`
	PhoneNumber string `gorm:"type:varchar(22);not null" json:"phoneNumber"`
	Role     string  `gorm:"type:varchar(20);default:'manager'" json:"role"`
	Status   bool    `gorm:"default:true" json:"status"`
	CompanyID uint   `json:"companyId"`
	Company  Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"company"`
	
}

func (m *Manager) BeforeSave(tx *gorm.DB) (err error) {
	m.Password, err = hashedPassword(m.Password)
	if err != nil {
		return err
	}
	return nil
}

func hashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (m *Manager) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}

func (m *Manager) GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"id":   m.ID,
		"role": m.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}
