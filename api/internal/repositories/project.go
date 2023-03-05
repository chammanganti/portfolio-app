package repository

import (
	model "api/internal/models"

	"gorm.io/gorm"
)

// Project repository interface
type ProjectRepositoryInterface interface {
	FindAll() ([]model.Project, error)
	Find(ID uint) (model.Project, error)
	FindByName(name string) (model.Project, error)
	UFindByName(name string) (model.Project, error)
	Create(newProject model.Project) (model.Project, error)
	Update(project, updatedProject model.Project) (model.Project, error)
	Delete(project model.Project) error
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

// Finds a project by ID
func (p ProjectRepository) Find(ID uint) (model.Project, error) {
	var project model.Project
	err := p.DB.First(&project, ID).Error
	return project, err
}

// Finds a project by name
func (p ProjectRepository) FindByName(name string) (model.Project, error) {
	var project model.Project
	err := p.DB.Where("name = ?", name).First(&project).Error
	return project, err
}

// Finds a project by name including deleted
func (p ProjectRepository) UFindByName(name string) (model.Project, error) {
	var project model.Project
	err := p.DB.Unscoped().Where("name = ?", name).First(&project).Error
	return project, err
}

// Creates a project
func (p ProjectRepository) Create(newProject model.Project) (model.Project, error) {
	err := p.DB.Create(&newProject).Error
	return newProject, err
}

// Updates a project
func (p ProjectRepository) Update(project, updatedProject model.Project) (model.Project, error) {
	err := p.DB.Model(&project).Select("*").Omit("ID", "created_at", "deleted_at").Updates(&updatedProject).Error
	return project, err
}

// Deletes a project
func (p ProjectRepository) Delete(project model.Project) error {
	return p.DB.Delete(&project).Error
}
