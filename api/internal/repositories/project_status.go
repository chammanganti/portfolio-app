package repository

import (
	model "api/internal/models"

	"gorm.io/gorm"
)

// Project status repository interface
type ProjectStatusRepositoryInterface interface {
	FindAll() ([]model.ProjectStatus, error)
	Find(ID uint) (model.ProjectStatus, error)
	FindByName(name string) (model.ProjectStatus, error)
	FindByNameProject(name string, projectID uint) (model.ProjectStatus, error)
	UFindByNameProject(name string, projectID uint) (model.ProjectStatus, error)
	Create(newProjectStatus model.ProjectStatus) (model.ProjectStatus, error)
	Update(projectStatus, updatedProjectStatus model.ProjectStatus) (model.ProjectStatus, error)
	Delete(projectStatus model.ProjectStatus) error
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
func (p ProjectRepositoryStatus) Find(ID uint) (model.ProjectStatus, error) {
	var projectStatus model.ProjectStatus
	err := p.DB.First(&projectStatus, ID).Error
	return projectStatus, err
}

// Finds a project status by name
func (p ProjectRepositoryStatus) FindByName(name string) (model.ProjectStatus, error) {
	var projectStatus model.ProjectStatus
	err := p.DB.Where("name = ?", name).First(&projectStatus).Error
	return projectStatus, err
}

// Finds a project status by name and project ID
func (p ProjectRepositoryStatus) FindByNameProject(name string, projectID uint) (model.ProjectStatus, error) {
	var projectStatus model.ProjectStatus
	err := p.DB.Where("name = ?", name).Where("project_id = ?", projectID).First(&projectStatus).Error
	return projectStatus, err
}

// Finds a project status by name and project ID including deleted
func (p ProjectRepositoryStatus) UFindByNameProject(name string, projectID uint) (model.ProjectStatus, error) {
	var projectStatus model.ProjectStatus
	err := p.DB.Unscoped().Where("name = ?", name).Where("project_id = ?", projectID).First(&projectStatus).Error
	return projectStatus, err
}

// Creates a project status
func (p ProjectRepositoryStatus) Create(newProjectStatus model.ProjectStatus) (model.ProjectStatus, error) {
	err := p.DB.Create(&newProjectStatus).Error
	return newProjectStatus, err
}

// Updates a project status
func (p ProjectRepositoryStatus) Update(projectStatus, updatedProjectStatus model.ProjectStatus) (model.ProjectStatus, error) {
	err := p.DB.Model(&projectStatus).Select("*").Omit("ID", "created_at", "deleted_at", "project_id").Updates(&updatedProjectStatus).Error
	return projectStatus, err
}

// Deletes a project status
func (p ProjectRepositoryStatus) Delete(projectStatus model.ProjectStatus) error {
	return p.DB.Delete(&projectStatus).Error
}
