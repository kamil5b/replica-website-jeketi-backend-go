package service

import (
	"replica-website-jeketi-backend-go/constant"
	"replica-website-jeketi-backend-go/database"
	"replica-website-jeketi-backend-go/model"
	"replica-website-jeketi-backend-go/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetUserAuth(c *fiber.Ctx) (*model.User, error) {
	db := database.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	key, _ := database.Redis.Get("blacklist-token")
	if len(key) > 0 {
		return nil, c.JSON(constant.UnauthorizedError)
	}

	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	id, ok := claims["id"]
	if !ok {
		return nil, c.JSON(constant.InternalServerError)
	}
	email, ok := claims["email"]
	if !ok {
		return nil, c.JSON(constant.InternalServerError)
	}
	role, ok := claims["role"]
	if !ok {
		return nil, c.JSON(constant.InternalServerError)
	}
	userIdString := id.(string)
	userID, err := uuid.Parse(userIdString)
	if err != nil {
		return nil, c.JSON(constant.InternalServerError)
	}
	user, err := repository.GetUserAuth(tx, userID, email.(string), role.(string))
	if err != nil {
		return nil, c.JSON(constant.InternalServerError)
	}
	return user, nil
}
