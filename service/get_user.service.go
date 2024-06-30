package service

import (
	"replica-website-jeketi-backend-go/database"
	"replica-website-jeketi-backend-go/model"
	"replica-website-jeketi-backend-go/repository"

	guuid "github.com/google/uuid"
)

func GetUserAuth(sessionid guuid.UUID) (*model.User, error) {
	db := database.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	session, err := repository.GetSession(tx, sessionid)
	if err != nil {
		return nil, err
	}

	user, err := repository.GetUserByID(tx, session.UserRefer)
	if err != nil {
		return nil, err
	}
	return user, nil
}
