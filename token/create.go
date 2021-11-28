package token

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/oauth2-go/database/models"
	"example.com/oauth2-go/database/repository"
	"example.com/oauth2-go/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type basicHeaders struct {
	AuthorizationBasic string `header:"Authorization"`
}

type clientCredentials struct {
	Scopes    string `json:"scopes" form:"scopes"`
	ExpiresIn int    `json:"expires_in" form:"expires_in"`
}

func Create(c *gin.Context) {
	header := basicHeaders{}
	clientCredentials := clientCredentials{}

	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}

	if err := c.ShouldBind(&clientCredentials); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	clientId, secret, err := getBasicClientIdAndSecretFromHeader(header)

	client, err := repository.FindByClientId(clientId)

	if err != nil {
		log.Println("Não achou o client")
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Client!"})
		return
	}

	h := sha256.Sum256([]byte(secret))
	hash := b64.StdEncoding.EncodeToString(h[:])

	if client.Secret != hash {
		log.Println(client.Secret + " <> " + secret)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Client!"})
		return
	}

	token := generateAccessToken(clientCredentials, client)

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

func getBasicClientIdAndSecretFromHeader(header basicHeaders) (string, string, error) {
	encodedClientAndPassword := strings.Replace(header.AuthorizationBasic, "Basic ", "", -1)
	clientAndPassword, _ := b64.URLEncoding.DecodeString(encodedClientAndPassword)

	separator := strings.Index(string(clientAndPassword), ":")
	clientId := clientAndPassword[:separator]
	secret := clientAndPassword[separator+1:]

	return string(clientId), string(secret), nil
}

func generateAccessToken(request clientCredentials, client models.Client) models.Token {
	timein := time.Now().Local().Add(time.Second * time.Duration(600))
	if request.ExpiresIn != 0 {
		timein = time.Now().Local().Add(time.Second * time.Duration(request.ExpiresIn))
	}

	token := models.Token{
		AccessToken:  uuid.New().String() + uuid.New().String(),
		RefreshToken: uuid.New().String() + uuid.New().String(),
		TokenType:    "bearer",
		ExpiresIn:    timein,
		Scopes:       request.Scopes,
		ClientId:     client.ID,
	}

	repository.SaveToken(token)

	return token
}
