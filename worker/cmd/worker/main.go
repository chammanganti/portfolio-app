package main

import (
	"worker/internal/config"
	"worker/internal/grpc"
	redis "worker/internal/store/redis"
	"worker/internal/worker"

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

	grpc, err := grpc.NewGRPC()
	if err != nil {
		log.Fatal("failed setting up the grpc connection")
	}
	defer grpc.Close()

	rdb := redis.NewRedis(context.Background(), config.REDIS_ADDR, config.REDIS_PASSWORD, config.REDIS_DB)

	awsConfig, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("failed loading the aws config")
	}

	worker := worker.NewWorker(config, awsConfig, grpc, rdb)

	log.Info("worker is running")

	cacheChanSize := 2
	cacheChan := make(chan string, cacheChanSize)
	for range time.Tick(time.Second * 5) {
		cacheChan <- worker.CacheProjects()
		cacheChan <- worker.CacheProjectStatuses()

		for range make([]int, cacheChanSize) {
			if v := <-cacheChan; v != "" {
				log.Info(v)
			}
		}

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
