package client

import (
	"net/http"

	"example.com/oauth2-go/database/repository"
	"example.com/oauth2-go/response"
	"github.com/gin-gonic/gin"
)

// GET /client/:id
// Find a client by id
func GetByClientId(c *gin.Context) {

	client, err := repository.FindByClientId(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	clientResponse := response.Client{Name: client.Name, RedirectUri: client.RedirectUri, Type: client.Type, ID: client.ID}

	c.JSON(http.StatusOK, clientResponse)
}
