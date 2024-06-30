package model

import "github.com/google/uuid"

type UserJkt48PointBalance struct {
	ID                uuid.UUID
	UserID            uuid.UUID
	PointBalance      int64
	BonusPointBalance int64
}
