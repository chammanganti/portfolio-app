package middleware

import (
	"api/internal/config"
	"api/internal/libs/constant"
	"api/internal/libs/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Auth config
type AuthConfig struct{}

// New auth middleware
func NewAuth(a ...AuthConfig) fiber.Handler {
	conf := config.GetConfig()

	invalidTokenErr := fiber.Map{
		"message": constant.INVALID_TOKEN,
	}

	return func(c *fiber.Ctx) error {
		token, ok := checkBearerToken(c.Get(fiber.HeaderAuthorization))
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(invalidTokenErr)
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		jss := claims["jss"]
		if jss != conf.JWT_SS {
			return c.Status(fiber.StatusUnauthorized).JSON(invalidTokenErr)
		}

		return c.Next()
	}
}

func checkBearerToken(authorization string) (string, bool) {
	authParts := strings.Split(authorization, " ")
	if len(authParts) != 2 {
		return "", false
	}

	if strings.ToLower(authParts[0]) != "bearer" {
		return "", false
	}

	return authParts[1], true
}
