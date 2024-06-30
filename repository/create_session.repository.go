package repository

import (
	"replica-website-jeketi-backend-go/model"

	"gorm.io/gorm"
)

func CreateSession(db *gorm.DB, session *model.Session) error {
	return db.Create(&session).Error
}
