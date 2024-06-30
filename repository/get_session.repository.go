package repository

import (
	"replica-website-jeketi-backend-go/model"

	guuid "github.com/google/uuid"

	"gorm.io/gorm"
)

func GetSession(db *gorm.DB, id guuid.UUID) (*model.Session, error) {
	session := new(model.Session)
	err := db.Select("id = ?", id).
		First(session).Error
	if err != nil {
		return nil, err
	}
	return session, nil
}
