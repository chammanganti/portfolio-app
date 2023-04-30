package config

import (
	"worker/internal/libs/constant"

	"os"
	"strconv"
)

// Config
type Config struct {
	ALLOW_WORKER   string
	API_URL        string
	GRPC_SERVER    string
	REDIS_ADDR     string
	REDIS_DB       int
	REDIS_PASSWORD string
}

// Returns the config
func GetConfig() *Config {
	redisDB := getEnv("REDIS_DB", "0")
	rDB, err := strconv.Atoi(redisDB)
	if err != nil {
		panic(err)
	}

	return &Config{
		ALLOW_WORKER:   getEnv(constant.ALLOW_WORKER, "false"),
		API_URL:        getEnv(constant.API_URL, "http://127.0.0.1:8000"),
		GRPC_SERVER:    getEnv(constant.GRPC_SERVER, "localhost:9000"),
		REDIS_ADDR:     getEnv(constant.REDIS_ADDR, "127.0.0.1:6379"),
		REDIS_DB:       rDB,
		REDIS_PASSWORD: getEnv(constant.REDIS_PASSWORD, ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
