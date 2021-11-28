package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ID          string
	Name        string
	Secret      string
	RedirectUri string
	Type        string
}
