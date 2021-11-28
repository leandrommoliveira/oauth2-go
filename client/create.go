package client

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"example.com/oauth2-go/database"
	"example.com/oauth2-go/database/models"
	"example.com/oauth2-go/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Save(c *gin.Context) {
	var clientJson request.Client
	if err := c.ShouldBindJSON(&clientJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientId := uuid.New().String()
	secret := uuid.New().String()

	h := sha256.Sum256([]byte(secret))
	hash := base64.StdEncoding.EncodeToString(h[:])

	database.DB.Create(&models.Client{
		ID:          clientId,
		Name:        clientJson.Name,
		Type:        clientJson.Type,
		RedirectUri: clientJson.RedirectUri,
		Secret:      hash,
	})

	c.JSON(200, gin.H{
		"client_id":     clientId,
		"client_secret": secret,
	})
}
