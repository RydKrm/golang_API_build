package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Email       string `gorm:"type:varchar(255);not null" json:"email"`
	Password    string `gorm:"type:varchar(255);not null" json:"-"`
	PhoneNumber string `gorm:"type:varchar(29);unique_index;not null" json:"phone_number"`
	Status      bool   `gorm:"default:true" json:"status"`
	Role        string    `gorm:"default:'admin'" json:"role"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
}


func (a *Admin) BeforeSave()(err error){
	a.Password, err = hashPassword(a.Password)
	if err != nil {
		return err;
	}
	return nil
}

func hashPassword(password string)(string, error){
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashPassword),nil
}

func (a *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}

func (a *Admin) GenerateToken()(string, error){
	claims := jwt.MapClaims{}
	claims["id"] = a.ID;
	claims["role"] = a.Role;
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}