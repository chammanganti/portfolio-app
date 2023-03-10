package main

import (
	"api/internal/config"
	store "api/internal/store/database"
	"api/internal/worker"
	"context"
	"strconv"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	log "github.com/sirupsen/logrus"
)

type Worker struct{}

func (w Worker) Run() error {
	config := config.GetConfig()
	if val, _ := strconv.ParseBool(config.ALLOW_WORKER); !val {
		log.Warn("worker is not allowed to run")
		return nil
	}

	db, err := store.NewDatabase()
	if err != nil {
		log.Fatal("failed setting up the database")
		return err
	}

	awsConfig, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("failed loading the aws config")
	}

	worker := worker.NewWorker(awsConfig, db)

	log.Info("worker is running")
	for range time.Tick(time.Second * 5) {
		// TODO: cache the projects with a TTL of 30s
		// TODO: cache the project statuses with a TTL of 30s

		worker.UpdateEC2Statuses()
	}

	return nil
}

func main() {
	worker := Worker{}
	if err := worker.Run(); err != nil {
		log.Fatal(err)
	}
}
