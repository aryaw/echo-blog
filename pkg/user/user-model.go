package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID            uint64 `gorm:"not null;primary_key"`
	FirstName     string `gorm:"size:255;not null"`
	LastName      string `gorm:"size:255;not null"`
	Email         string `gorm:"size:255;not null;uniqueIndex"`
	Password      string `gorm:"size:255;not null"`
	RememberToken string `gorm:"size:255"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
