package middleware

import (
	"replica-website-jeketi-backend-go/model"
	"replica-website-jeketi-backend-go/service"

	"github.com/gofiber/fiber/v2"
)

func Authenticated(c *fiber.Ctx) error {
	json := new(model.Session)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Session Format",
		})
	}
	user, err := service.GetUserAuth(json.Sessionid)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "404: not found",
		})
	}
	c.Locals("user", user)
	return c.Next()
}
