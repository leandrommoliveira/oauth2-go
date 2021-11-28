package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	ID           int `gorm:"primaryKey;autoIncrement"`
	AccessToken  string
	RefreshToken string
	TokenType    string
	Scopes       string
	ExpiresIn    time.Time
	TokenParent  int
	ClientId     string `gorm:"size:100"`
	Client       Client `gorm:"foreignKey:ClientId"`
}
