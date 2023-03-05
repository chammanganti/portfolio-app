package main

import (
	"api/internal/config"
	handler "api/internal/handlers"
	"api/internal/router"
	service "api/internal/services"
	store "api/internal/store/database"
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type App struct{}

func (a App) Run() error {
	config := config.GetConfig()

	fiber := fiber.New()

	db, err := store.NewDatabase()
	if err != nil {
		log.Fatal("failed setting up the database")
		return err
	}

	router.SetupRoutes(fiber, db)

	awsConfig, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("failed loading the aws config")
	}
	aws := fiber.Group("aws")
	awsService := service.NewAWSService()
	awsHandler := handler.NewAWSHandler(awsConfig, awsService)
	aws.Get("", awsHandler.GetEC2Status)

	if err := fiber.Listen(config.ADDR); err != nil {
		log.Fatal("failed setting up the server")
		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
