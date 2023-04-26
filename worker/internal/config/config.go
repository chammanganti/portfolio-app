package config

import (
	"worker/internal/libs/constant"

	"os"
	"strconv"
)

// Config
type Config struct {
	ALLOW_WORKER   string
	DB_HOST        string
	DB_PORT        string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_NAME        string
	DB_SSL_MODE    string
	DB_TIMEZONE    string
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
		DB_HOST:        getEnv(constant.DB_HOST, "127.0.0.1"),
		DB_PORT:        getEnv(constant.DB_PORT, "5432"),
		DB_USERNAME:    getEnv(constant.DB_USERNAME, "postgres"),
		DB_PASSWORD:    getEnv(constant.DB_PASSWORD, "postgres"),
		DB_NAME:        getEnv(constant.DB_NAME, "portfolio"),
		DB_SSL_MODE:    getEnv(constant.DB_SSL_MODE, "disable"),
		DB_TIMEZONE:    getEnv(constant.DB_TIMEZONE, "Asia/Manila"),
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
