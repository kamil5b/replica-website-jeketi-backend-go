package handlers

import (
	"replica-website-jeketi-backend-go/constant"
	"replica-website-jeketi-backend-go/database"
	"replica-website-jeketi-backend-go/model"
	"replica-website-jeketi-backend-go/request"
	"replica-website-jeketi-backend-go/service"
	"replica-website-jeketi-backend-go/util"

	"github.com/gofiber/fiber/v2"
)

type User model.User
type Session model.Session
type Product model.Product

func Login(c *fiber.Ctx) error {
	json := new(request.LoginRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(constant.BindingError)
	}

	token, err := service.Login(c, *json)
	if err != nil {
		return nil
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"token":   token,
	})
}

func Logout(c *fiber.Ctx) error {
	err := service.Logout(c)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func Register(c *fiber.Ctx) error {
	json := new(request.RegisterRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(constant.BindingError)
	}

	token, err := service.Register(c, *json)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"token":   token,
	})
}

func GetUserInfo(c *fiber.Ctx) error {
	user := c.Locals("user").(User)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB
	json := new(request.DeleteUserRequest)
	user := c.Locals("user").(User)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	if !util.ComparePasswords(user.Password, []byte(json.Password)) {
		return c.JSON(fiber.Map{
			"code":    401,
			"message": "Invalid Password",
		})
	}
	db.Model(&user).Association("Sessions").Delete()
	db.Model(&user).Association("Products").Delete()
	db.Delete(&user)
	c.ClearCookie("sessionid")
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func ChangePassword(c *fiber.Ctx) error {
	db := database.DB
	user := c.Locals("user").(User)
	json := new(request.ChangePasswordRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	if !util.ComparePasswords(user.Password, []byte(json.Password)) {
		return c.JSON(fiber.Map{
			"code":    401,
			"message": "Invalid Password",
		})
	}
	user.Password = util.HashAndSalt([]byte(json.NewPassword))
	db.Save(&user)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}
