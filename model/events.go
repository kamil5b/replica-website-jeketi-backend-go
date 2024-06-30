package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID             uuid.UUID
	Title          string
	Date           time.Time
	PricePerTicket uint32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
