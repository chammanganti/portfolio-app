package service

import (
	"api/internal/libs/constant"
	model "api/internal/models"
	repository "api/internal/repositories"
	"errors"
	"fmt"
)

// Project status service interface
type ProjectStatusServiceInterface interface {
	Create(projectStatus model.ProjectStatus) (model.ProjectStatus, error)
	Update(ID uint, updatedProjectStatus model.ProjectStatus) (model.ProjectStatus, error)
	Delete(ID uint) error
}

// Project status service
type ProjectStatusService struct {
	Repository repository.ProjectStatusRepositoryInterface
}

// New project status service
func NewProjectStatusService(repository repository.ProjectStatusRepositoryInterface) ProjectStatusServiceInterface {
	return &ProjectStatusService{
		Repository: repository,
	}
}

// Creates a project status
func (p ProjectStatusService) Create(projectStatus model.ProjectStatus) (model.ProjectStatus, error) {
	if existing, _ := p.Repository.UFindByNameProject(projectStatus.Name, projectStatus.ProjectID); existing.ID != 0 {
		return model.ProjectStatus{}, errors.New(constant.ALREADY_EXISTS)
	}

	return p.Repository.Create(projectStatus)
}

// Updates a project status
func (p ProjectStatusService) Update(ID uint, updatedProjectStatus model.ProjectStatus) (model.ProjectStatus, error) {
	projectStatus, err := p.Repository.Find(ID)
	if err != nil {
		return model.ProjectStatus{}, err
	}

	if projectStatus.Name != updatedProjectStatus.Name {
		if existing, _ := p.Repository.UFindByNameProject(updatedProjectStatus.Name, projectStatus.ProjectID); existing.ID != 0 {
			return model.ProjectStatus{}, fmt.Errorf(constant.RECORD_ALREADY_EXISTS, updatedProjectStatus.Name)
		}
	}

	return p.Repository.Update(projectStatus, updatedProjectStatus)
}

// Deletes a project status
func (p ProjectStatusService) Delete(ID uint) error {
	projectStatus, err := p.Repository.Find(ID)
	if err != nil {
		return err
	}

	return p.Repository.Delete(projectStatus)
}
