package model

import (
	"time"

	guuid "github.com/google/uuid"
)

//TODO: Refactor User table

type User struct {
	ID        guuid.UUID `gorm:"primaryKey" json:"-"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Sessions  []Session  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	Products  []Product  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time  `json:"-" `
	UpdatedAt time.Time  `json:"-"`
}
