package service

import (
	"api/internal/libs/constant"
	model "api/internal/models"
	repository "api/internal/repositories"
	"errors"
	"fmt"
)

// Project service interface
type ProjectServiceInterface interface {
	Create(project model.Project) (model.Project, error)
	Update(ID uint, updatedProject model.Project) (model.Project, error)
	Delete(ID uint) error
}

// Project service
type ProjectService struct {
	Repository repository.ProjectRepositoryInterface
}

// New project service
func NewProjectService(repository repository.ProjectRepositoryInterface) ProjectServiceInterface {
	return &ProjectService{
		Repository: repository,
	}
}

// Creates a project
func (p ProjectService) Create(project model.Project) (model.Project, error) {
	if existing, _ := p.Repository.UFindByName(project.Name); existing.ID != 0 {
		return model.Project{}, errors.New(constant.ALREADY_EXISTS)
	}

	return p.Repository.Create(project)
}

// Updates a project
func (p ProjectService) Update(ID uint, updatedProject model.Project) (model.Project, error) {
	project, err := p.Repository.Find(ID)
	if err != nil {
		return model.Project{}, err
	}

	if project.Name != updatedProject.Name {
		if p, _ := p.Repository.UFindByName(updatedProject.Name); p.ID != 0 {
			return model.Project{}, fmt.Errorf(constant.RECORD_ALREADY_EXISTS, updatedProject.Name)
		}
	}

	return p.Repository.Update(project, updatedProject)
}

// Deletes a project
func (p ProjectService) Delete(ID uint) error {
	project, err := p.Repository.Find(ID)
	if err != nil {
		return err
	}

	return p.Repository.Delete(project)
}
