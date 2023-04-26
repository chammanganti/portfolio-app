package repository

import (
	model "worker/internal/models"

	"gorm.io/gorm"
)

// Project repository interface
type ProjectRepositoryInterface interface {
	FindAll() ([]model.Project, error)
}

// Project repository
type ProjectRepository struct {
	DB *gorm.DB
}

// New project repository
func NewProjectRepository(db *gorm.DB) ProjectRepositoryInterface {
	return &ProjectRepository{
		DB: db,
	}
}

// Finds all projects
func (p ProjectRepository) FindAll() ([]model.Project, error) {
	var projects []model.Project
	err := p.DB.Find(&projects).Error
	return projects, err
}
