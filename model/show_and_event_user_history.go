package model

import (
	"time"

	"github.com/google/uuid"
)

type ShowAndEventUserHistory struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	ShowAndEventID uuid.UUID
	TicketType     string // GEN, OFC, VIP
	ShowType       string // THEATRE, OFC_EVENT, MINI_SHOW
	Time           time.Time
	Status         string // LOSE, WIN, PENDING
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
