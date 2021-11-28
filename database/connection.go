package database

import (
	"example.com/oauth2-go/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:my-secret-pw@tcp(172.17.0.2:3306)/oauth2?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/oauth2?charset=utf8mb4&parseTime=True&loc=Local"
	//database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&models.Client{})
	database.AutoMigrate(&models.Token{})

	DB = database
}
