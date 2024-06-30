package handlers

import (
	"time"

	"replica-website-jeketi-backend-go/constant"
	"replica-website-jeketi-backend-go/database"
	"replica-website-jeketi-backend-go/model"
	"replica-website-jeketi-backend-go/request"
	"replica-website-jeketi-backend-go/service"
	"replica-website-jeketi-backend-go/util"

	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type User model.User
type Session model.Session
type Product model.Product

func Login(c *fiber.Ctx) error {
	json := new(request.LoginRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(constant.BindingError)
	}

	session, err := service.Login(c, *json)

	if err != nil {
		return nil
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    session,
	})
}

//TODO: CREATE SERVICE LAYER

func Logout(c *fiber.Ctx) error {
	db := database.DB
	json := new(Session)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	session := Session{}
	query := Session{Sessionid: json.Sessionid}
	err := db.First(&session, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Session not found",
		})
	}
	db.Delete(&session)
	c.ClearCookie("sessionid")
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	json := new(request.CreateUserRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	password := util.HashAndSalt([]byte(json.Password))
	err := checkmail.ValidateFormat(json.Email)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Email Address",
		})
	}
	new := User{
		Password: password,
		Email:    json.Email,
		ID:       guuid.New(),
	}
	found := User{}
	query := User{Email: json.Email}
	err = db.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "User already exists",
		})
	}
	db.Create(&new)
	session := Session{UserRefer: new.ID, Sessionid: guuid.New()}
	err = db.Create(&session).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Creation Error",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "sessionid",
		Expires:  time.Now().Add(5 * 24 * time.Hour),
		Value:    session.Sessionid.String(),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    session,
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
