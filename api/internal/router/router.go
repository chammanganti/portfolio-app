package router

import (
	handler "api/internal/handlers"
	middleware "api/internal/router/middlewares"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const v1Prefix = "/api/v1"

// Sets up the routes
func SetupRoutes(f *fiber.App, db *gorm.DB) {
	authMiddleware := middleware.NewAuth(middleware.AuthConfig{})

	projectHandler := handler.NewProjectHandler(db)
	projectStatusHandler := handler.NewProjectStatusHandler(db)

	v1 := f.Group(v1Prefix)

	projects := v1.Group("projects")
	projects.Get("/", projectHandler.All)
	projects.Get("/:id", projectHandler.Get)
	projects.Get("/name/:name", projectHandler.GetByName)
	projects.Post("/", projectHandler.Create)
	projects.Put("/:id", projectHandler.Update)
	projects.Delete("/:id", projectHandler.Delete)

	projectStatuses := v1.Group("project-statuses")
	projectStatuses.Get("/", projectStatusHandler.All)
	projectStatuses.Get("/:id", projectStatusHandler.Get)
	projectStatuses.Post("/", projectStatusHandler.Create)
	projectStatuses.Put("/:id", projectStatusHandler.Update)
	projectStatuses.Delete("/:id", projectStatusHandler.Delete)

	auth := v1.Group("/auth")
	auth.Use(authMiddleware)
	auth.Post("/check", func(c *fiber.Ctx) error {
		token := c.Get(fiber.HeaderAuthorization)
		token = strings.TrimPrefix(token, "Bearer ")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"access_token": token,
		})
	})

	health := v1.Group("/health")
	health.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("not dead")
	})
}
