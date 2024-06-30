package model

import (
	"time"

	"github.com/google/uuid"
)

type ShowAndEvent struct {
	ID        uuid.UUID
	Title     string
	Date      time.Time
	Price     uint32
	Location  string
	Quota     uint16
	CreatedAt time.Time
	UpdatedAt time.Time
}
