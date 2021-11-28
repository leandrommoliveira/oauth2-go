package repository

import (
	"example.com/oauth2-go/database"
	"example.com/oauth2-go/database/models"
)

func FindByClientId(clientId string) (models.Client, error) {
	var client models.Client

	if err := database.DB.Where("ID = ?", clientId).First(&client).Error; err != nil {
		return client, err
	}

	return client, nil
}
