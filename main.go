package main

import (
	"fmt"
	"kms/config"
	"kms/controllers"
	"kms/db"
	"kms/repositories"
	"kms/router"
	"kms/services"
	"log"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Configuration has not been loaded")
	}

	db.ConnectDatabase()

	r := router.SetupRouter()
	fmt.Println("starting server")

	articleRepo := repositories.NewMongoArticleRepositories()
	articleService := services.NewArticleService(articleRepo)
	controllers.NewArticleController(articleService)

	r.Run(":8080")

	db.DisConnectDatabase()
}
