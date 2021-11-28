package repository

import (
	"example.com/oauth2-go/database"
	"example.com/oauth2-go/database/models"
)

func SaveToken(token models.Token) {
	database.DB.Create(&token)
}

func FindTokenByAccessToken(accessToken string) (models.Token, error) {
	var token models.Token

	if err := database.DB.Where("access_token = ?", accessToken).First(&token).Error; err != nil {
		return token, err
	}

	return token, nil
}
