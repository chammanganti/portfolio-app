package service

import (
	"worker/internal/libs/constant"
	model "worker/internal/models"
	repository "worker/internal/repositories"

	"fmt"
)

// Project status service interface
type ProjectStatusServiceInterface interface {
	Update(projectStatusId string, updatedProjectStatus model.ProjectStatus) (model.ProjectStatus, error)
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

// Updates a project status
func (p ProjectStatusService) Update(projectStatusId string, updatedProjectStatus model.ProjectStatus) (model.ProjectStatus, error) {
	projectStatus, err := p.Repository.Find(projectStatusId)
	if err != nil {
		return model.ProjectStatus{}, err
	}

	if projectStatus.Name != updatedProjectStatus.Name {
		if existing, _ := p.Repository.FindByNameProject(updatedProjectStatus.Name, projectStatus.ProjectId); existing.ProjectStatusId != "" {
			return model.ProjectStatus{}, fmt.Errorf(constant.RECORD_ALREADY_EXISTS, updatedProjectStatus.Name)
		}
	}

	return p.Repository.Update(projectStatus, updatedProjectStatus)
}
