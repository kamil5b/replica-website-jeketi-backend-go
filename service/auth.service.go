package service

import (
	"errors"
	"replica-website-jeketi-backend-go/constant"
	"replica-website-jeketi-backend-go/database"
	"replica-website-jeketi-backend-go/model"
	"replica-website-jeketi-backend-go/repository"
	"replica-website-jeketi-backend-go/request"
	"replica-website-jeketi-backend-go/util"
	"strconv"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx, req request.LoginRequest) (*string, error) {
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
	token, err := util.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, c.JSON(constant.InternalServerError)
	}
	return &token, nil
}

func Logout(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	err := blacklistToken(token)
	if err != nil {
		return c.JSON(constant.InternalServerError)
	}
	return nil
}

func Register(c *fiber.Ctx, req request.RegisterRequest) (*string, error) {
	db := database.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	password := util.HashAndSalt([]byte(req.Password))
	err := checkmail.ValidateFormat(req.Email)
	if err != nil {
		return nil, c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Email Address",
		})
	}
	userID := guuid.New()
	user := model.User{
		ID:       userID,
		Email:    req.Email,
		Password: password,
		Name:     req.Name,
		Role:     "FAN",
		Properties: &model.UserProperty{
			ID:           guuid.New(),
			UserID:       userID,
			Birthday:     req.Birthday,
			Gender:       req.Gender,
			IdentityNo:   req.IdentityNo,
			IdentityType: req.IdentityType,
			Poscode:      req.Poscode,
			Address:      req.Address,
			City:         req.City,
			OshimenID:    req.OshimenID,
		},
	}

	// TODO: UPLOAD USER ATTACHMENTS: Profile picture URL
	// TODO: GENERATE BARCODE

	err = repository.SaveUser(tx, &user)
	if err != nil {
		return nil, c.JSON(constant.InternalServerError)
	}
	token, err := util.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, c.JSON(constant.InternalServerError)
	}
	return &token, nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}

func blacklistToken(tokenString string) error {
	tokenExpiryHour := util.GetEnv("TOKEN_EXPIRY_HOUR")
	expiryHour, err := strconv.Atoi(tokenExpiryHour)
	if err != nil {
		return err
	}
	return database.Redis.Set("blacklist-token", []byte(tokenString), time.Hour*time.Duration(expiryHour))
}
