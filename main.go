package main

import (
	"github.com/Djuanzz/gorest/config"
	"github.com/Djuanzz/gorest/docs"
	"github.com/Djuanzz/gorest/routes"
	"github.com/Djuanzz/gorest/utils"
	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	// --- FOR ENVIRONMENT VARIABLES
	enviroment := utils.GetEnv("ENVIRONMENT", "development")

	if enviroment == "development" {
		err := godotenv.Load()
		if err != nil {
			panic("Failed to load .env file!")
		}
	}

	docs.SwaggerInfo.Title = "Go REST API"
	docs.SwaggerInfo.Description = "This is a sample server for a RESTful service written in Go."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDb()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	// --- ROUTES
	r := routes.SetupRouter(db)
	r.Run("localhost:8080")
}
