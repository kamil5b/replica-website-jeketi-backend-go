package repository

import (
	"replica-website-jeketi-backend-go/model"

	guuid "github.com/google/uuid"

	"gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, id guuid.UUID) (*model.User, error) {
	user := new(model.User)
	err := db.Select("id = ?", id).
		First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	user := new(model.User)
	err := db.Select("email = ?", email).
		First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
