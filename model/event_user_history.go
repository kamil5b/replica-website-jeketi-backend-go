package model

import (
	"time"

	"github.com/google/uuid"
)

type EventUserHistory struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	EventLineUpID uuid.UUID
	TicketType    string // GEN, OFC, VIP
	ShowType      string // 2SHOT, HANDSHAKE, MEET_AND_GREET, VIDEO_CALL
	Time          time.Time
	Status        string // DONE, BOOKED
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
