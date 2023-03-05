package store

import (
	"fmt"

	"api/internal/config"
	model "api/internal/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New database
func NewDatabase() (*gorm.DB, error) {
	log.Info("setting up the db connection")

	config := config.GetConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DB_HOST,
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
		config.DB_SSL_MODE,
		config.DB_TIMEZONE)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Project{}, &model.ProjectStatus{}); err != nil {
		return nil, err
	}

	return db, nil
}
