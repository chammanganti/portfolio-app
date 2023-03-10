package main

import (
	"api/internal/config"
	"api/internal/router"
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

	awsConfig, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("failed loading the aws config")
		return err
	}

	router.SetupRoutes(fiber, awsConfig, db)

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
