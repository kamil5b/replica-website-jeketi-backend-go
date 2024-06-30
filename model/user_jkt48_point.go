package model

import "github.com/google/uuid"

type UserJkt48Point struct {
	ID                uuid.UUID
	UserID            uuid.UUID
	PointBalance      uint64
	BonusPointBalance uint64
}
