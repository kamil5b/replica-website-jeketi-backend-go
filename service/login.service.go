package service

import (
	"errors"
	"replica-website-jeketi-backend-go/constant"
	"replica-website-jeketi-backend-go/database"
	"replica-website-jeketi-backend-go/model"
	"replica-website-jeketi-backend-go/repository"
	"replica-website-jeketi-backend-go/request"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx, req request.LoginRequest) (*model.Session, error) {

	db := database.DB
	user, err := repository.GetUserByEmail(db, req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, c.JSON(constant.DataNotFoundError)
	}
	if err != nil {
		return nil, c.JSON(constant.InternalServerError)
	}
	if !comparePasswords(user.Password, []byte(req.Password)) {
		return nil, c.JSON(constant.InvalidPasswordError)
	}
	session := &model.Session{
		UserRefer: user.ID,
		Expires:   SessionExpires(),
		Sessionid: uuid.New(),
	}
	err = repository.CreateSession(db, session)
	if err != nil {
		return nil, c.JSON(constant.InternalServerError)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "sessionid",
		Expires:  SessionExpires(),
		Value:    session.Sessionid.String(),
		HTTPOnly: true,
	})
	return session, nil
}

// Universal date the Session Will Expire
func SessionExpires() time.Time {
	return time.Now().Add(5 * 24 * time.Hour)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}
