package handler

import (
	"api/internal/libs/validator"
	model "api/internal/models"
	repository "api/internal/repositories"
	service "api/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Project status handler interface
type ProjectStatusHandlerInterface interface {
	BaseHandlerInterface
}

// Project status handler
type ProjectStatusHandler struct {
	Repository repository.ProjectStatusRepositoryInterface
	Service    service.ProjectStatusServiceInterface
}

// New project status handler
func NewProjectStatusHandler(db *gorm.DB) ProjectStatusHandlerInterface {
	repository := repository.NewProjectStatusRepository(db)
	service := service.NewProjectStatusService(repository)
	return &ProjectStatusHandler{
		Repository: repository,
		Service:    service,
	}
}

// Gets all project statuses
func (p ProjectStatusHandler) All(c *fiber.Ctx) error {
	projectStatuses, err := p.Repository.FindAll()
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(projectStatuses)
}

// Gets a project status by ID
func (p ProjectStatusHandler) Get(c *fiber.Ctx) error {
	id, err := GetIDParam(c)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	projectStatus, err := p.Repository.Find(id)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(projectStatus)
}

// Creates a project status
func (p ProjectStatusHandler) Create(c *fiber.Ctx) error {
	var projectStatus model.ProjectStatus
	if err := c.BodyParser(&projectStatus); err != nil {
		return RespondWithError(c, fiber.StatusUnprocessableEntity, err)
	}

	errors := validator.ValidateStruct(model.ProjectStatusRequest{
		Name:      projectStatus.Name,
		IsHealthy: projectStatus.IsHealthy,
		ProjectID: projectStatus.ProjectID,
	})
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	projectStatus, err := p.Service.Create(projectStatus)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusCreated).JSON(projectStatus)
}

// Updates a project status
func (p ProjectStatusHandler) Update(c *fiber.Ctx) error {
	id, err := GetIDParam(c)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	var projectStatus model.ProjectStatus
	if err := c.BodyParser(&projectStatus); err != nil {
		return RespondWithError(c, fiber.StatusUnprocessableEntity, err)
	}

	errors := validator.ValidateStruct(model.ProjectStatusRequest{
		Name:      projectStatus.Name,
		IsHealthy: projectStatus.IsHealthy,
		ProjectID: projectStatus.ProjectID,
	})
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	updatedProjectStatus, err := p.Service.Update(id, projectStatus)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(updatedProjectStatus)
}

// Deletes a project status
func (p ProjectStatusHandler) Delete(c *fiber.Ctx) error {
	id, err := GetIDParam(c)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	if err := p.Service.Delete(id); err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
