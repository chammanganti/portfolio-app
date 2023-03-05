package handler

import (
	"api/internal/libs/validator"
	model "api/internal/models"
	repository "api/internal/repositories"
	service "api/internal/services"
	"strings"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Project handler interface
type ProjectHandlerInterface interface {
	BaseHandlerInterface
	GetByName(c *fiber.Ctx) error
}

// Project handler
type ProjectHandler struct {
	Repository repository.ProjectRepositoryInterface
	Service    service.ProjectServiceInterface
}

// New project handler
func NewProjectHandler(db *gorm.DB) ProjectHandlerInterface {
	repository := repository.NewProjectRepository(db)
	service := service.NewProjectService(repository)
	return &ProjectHandler{
		Repository: repository,
		Service:    service,
	}
}

// Gets all projects
func (p ProjectHandler) All(c *fiber.Ctx) error {
	projects, err := p.Repository.FindAll()
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(projects)
}

// Gets a project by ID
func (p ProjectHandler) Get(c *fiber.Ctx) error {
	id, err := GetIDParam(c)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	project, err := p.Repository.Find(id)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(project)
}

// Gets a project by name
func (p ProjectHandler) GetByName(c *fiber.Ctx) error {
	var project model.Project
	var err error

	name := c.Params("name")
	name = strings.ReplaceAll(name, "%20", " ")
	if err != nil {
		return RespondWithError(c, fiber.StatusBadRequest, err)
	}
	isUnscoped := c.Query("unscoped")
	log.Info(name)

	if isUnscoped == "true" {
		project, err = p.Repository.UFindByName(name)
	} else {
		project, err = p.Repository.FindByName(name)
	}
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(project)
}

// Creates a project
func (p ProjectHandler) Create(c *fiber.Ctx) error {
	var project model.Project
	if err := c.BodyParser(&project); err != nil {
		return RespondWithError(c, fiber.StatusUnprocessableEntity, err)
	}

	errors := validator.ValidateStruct(model.ProjectRequest{
		Name:             project.Name,
		Description:      project.Description,
		URL:              project.URL,
		GithubRepository: project.GithubRepository,
	})
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	newProject, err := p.Service.Create(project)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusCreated).JSON(newProject)
}

// Updates a project
func (p ProjectHandler) Update(c *fiber.Ctx) error {
	id, err := GetIDParam(c)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	var project model.Project
	if err := c.BodyParser(&project); err != nil {
		return RespondWithError(c, fiber.StatusUnprocessableEntity, err)
	}

	errors := validator.ValidateStruct(model.ProjectRequest{
		Name:             project.Name,
		Description:      project.Description,
		URL:              project.URL,
		GithubRepository: project.GithubRepository,
	})
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	updatedProject, err := p.Service.Update(id, project)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(updatedProject)
}

// Deletes a project
func (p ProjectHandler) Delete(c *fiber.Ctx) error {
	id, err := GetIDParam(c)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	if err := p.Service.Delete(id); err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
