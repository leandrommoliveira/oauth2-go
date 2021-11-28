package token

import (
	"net/http"
	"strconv"
	"time"

	"example.com/oauth2-go/database/repository"
	"example.com/oauth2-go/response"
	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	token, err := repository.FindTokenByAccessToken(c.Param("access_token"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	expiresIn := int(token.ExpiresIn.Unix() - time.Now().Unix())

	response := response.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		Scopes:       token.Scopes,
		ExpiresIn:    strconv.Itoa(expiresIn),
	}

	c.JSON(http.StatusOK, response)
}
