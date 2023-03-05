package handler

import (
	"api/internal/libs/constant"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Base handler interface
type BaseHandlerInterface interface {
	All(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

// Respond with error
func RespondWithError(c *fiber.Ctx, defaultStatusCode int, err error) error {
	e := strings.ToLower(err.Error())

	return c.Status(getSpecificStatus(defaultStatusCode, e)).JSON(fiber.Map{
		"message": e,
	})
}

// Returns the specific status code
func getSpecificStatus(defaultStatusCode int, err string) int {
	statusCode := defaultStatusCode

	switch err {
	case constant.RECORD_NOT_FOUND:
		statusCode = fiber.StatusNotFound
	}

	if strings.Contains(err, constant.ALREADY_EXISTS) {
		statusCode = fiber.StatusConflict
	}

	return statusCode
}

// Returns ID param of type uint
func GetIDParam(c *fiber.Ctx) (uint, error) {
	id := c.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return uint(i), nil
}
