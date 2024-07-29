package middleware

import (
	"replica-website-jeketi-backend-go/service"

	"github.com/gofiber/fiber/v2"
)

func Authenticated(c *fiber.Ctx) error {
	user, err := service.GetUserAuth(c)
	if err != nil {
		return err
	}
	c.Locals("user", user)
	return c.Next()
}
