package main

import (
	"example.com/oauth2-go/client"
	"example.com/oauth2-go/database"
	"example.com/oauth2-go/token"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/client", client.Save)
	r.GET("/client/:id", client.GetByClientId)

	r.POST("/oauth2/token", token.Create)
	r.GET("/oauth2/token/:access_token", token.Find)

	database.ConnectDatabase()

	r.Run()
}
