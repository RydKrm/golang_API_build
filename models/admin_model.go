package models

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
    ID          uint       `gorm:"primary_key" json:"id"`
    Name        string     `gorm:"type:varchar(255);not null" json:"name"`
    Email       string     `gorm:"type:varchar(255);not null" json:"email"`
    Password    string     `gorm:"type:varchar(255);not null" json:"-"`
    PhoneNumber string     `gorm:"type:varchar(29);unique_index;not null" json:"phone_number"`
    Status      bool       `gorm:"default:true" json:"status"`
    Role        string     `gorm:"default:'admin'" json:"role"`
    CreatedAt   time.Time  `json:"create_at"`
    UpdatedAt   time.Time  `json:"update_at"`
}

func (a *Admin) BeforeSave(db *gorm.DB) error {
    hashedPassword, err := hashPassword(a.Password)
    if err != nil {
        return err
    }
    a.Password = hashedPassword
    return nil
}

func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

func (a *Admin) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
    return err == nil
}

func (a *Admin) GenerateToken() (string, error) {
    claims := jwt.MapClaims{}
    claims["id"] = a.ID
    claims["role"] = a.Role
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    // Retrieve the TOKEN_SECRET from environment variables and sign the token
    secret := []byte(os.Getenv("TOKEN_SECRET"))
    if len(secret) == 0 {
        return "", fmt.Errorf("TOKEN_SECRET environment variable not set")
    }
    
    signedToken, err := token.SignedString(secret)
    if err != nil {
        return "", fmt.Errorf("error signing token: %v", err)
    }
    
    return signedToken, nil
}

