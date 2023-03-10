package config

import (
	"api/internal/libs/constant"
	"os"
)

// Config
type Config struct {
	ADDR          string
	ALLOW_WORKER  string
	DB_HOST       string
	DB_PORT       string
	DB_USERNAME   string
	DB_PASSWORD   string
	DB_NAME       string
	DB_SSL_MODE   string
	DB_TIMEZONE   string
	JWT_AT_SECRET string
	JWT_SS        string
}

// Returns the config
func GetConfig() *Config {
	return &Config{
		ADDR:          getEnv(constant.ADDR, ":8000"),
		ALLOW_WORKER:  getEnv(constant.ALLOW_WORKER, "false"),
		DB_HOST:       getEnv(constant.DB_HOST, "127.0.0.1"),
		DB_PORT:       getEnv(constant.DB_PORT, "5432"),
		DB_USERNAME:   getEnv(constant.DB_USERNAME, "postgres"),
		DB_PASSWORD:   getEnv(constant.DB_PASSWORD, "postgres"),
		DB_NAME:       getEnv(constant.DB_NAME, "portfolio"),
		DB_SSL_MODE:   getEnv(constant.DB_SSL_MODE, "disable"),
		DB_TIMEZONE:   getEnv(constant.DB_TIMEZONE, "Asia/Manila"),
		JWT_AT_SECRET: getEnv(constant.JWT_AT_SECRET, "superstrongsecretkey"),
		JWT_SS:        getEnv(constant.JWT_SS, "hiddensecret"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
