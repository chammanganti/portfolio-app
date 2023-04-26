package repository

import (
	model "worker/internal/models"

	"gorm.io/gorm"
)

// Project status repository interface
type ProjectStatusRepositoryInterface interface {
	FindAll() ([]model.ProjectStatus, error)
	Find(projectStatusId string) (model.ProjectStatus, error)
	FindByNameProject(name string, projectId string) (model.ProjectStatus, error)
	Update(projectStatus, updatedProjectStatus model.ProjectStatus) (model.ProjectStatus, error)
}

// Project status repository
type ProjectRepositoryStatus struct {
	DB *gorm.DB
}

// New project status repository
func NewProjectStatusRepository(db *gorm.DB) ProjectStatusRepositoryInterface {
	return &ProjectRepositoryStatus{
		DB: db,
	}
}

// Finds all project statuses
func (p ProjectRepositoryStatus) FindAll() ([]model.ProjectStatus, error) {
	var projectStatuses []model.ProjectStatus
	err := p.DB.Find(&projectStatuses).Error
	return projectStatuses, err
}

// Finds a project status by ID
func (p ProjectRepositoryStatus) Find(projectStatusId string) (model.ProjectStatus, error) {
	var projectStatus model.ProjectStatus
	err := p.DB.First(&projectStatus, projectStatusId).Error
	return projectStatus, err
}

// Finds a project status by name and project ID
func (p ProjectRepositoryStatus) FindByNameProject(name string, projectId string) (model.ProjectStatus, error) {
	var projectStatus model.ProjectStatus
	err := p.DB.Where("name = ?", name).Where("project_id = ?", projectId).First(&projectStatus).Error
	return projectStatus, err
}

// Updates a project status
func (p ProjectRepositoryStatus) Update(projectStatus, updatedProjectStatus model.ProjectStatus) (model.ProjectStatus, error) {
	err := p.DB.Model(&projectStatus).Select("*").Omit("ID", "created_at", "deleted_at", "project_id").Updates(&updatedProjectStatus).Error
	return projectStatus, err
}
