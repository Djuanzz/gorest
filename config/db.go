package config

import (
	"github.com/Djuanzz/gorest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	username := "root"
	password := "root"
	host := "@tcp(localhost:3306)"
	dbName := "go_rest"

	dsn := username + ":" + password + host + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Movie{}, &models.AgeRatingCategory{}, &models.User{})

	return db
}
