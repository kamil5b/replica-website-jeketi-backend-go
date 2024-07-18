package model

import (
	"time"

	"github.com/google/uuid"
)

type EventUserHistory struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID        uuid.UUID `json:"userId"`
	EventLineUpID uuid.UUID `json:"eventLineUpId"`
	TicketType    string    `json:"ticketType"` // GEN, OFC, VIP
	Status        string    `json:"status"`     // DONE, BOOKED
	CreatedAt     time.Time `json:"createdAt" gorm:"default:now()"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"default:now()"`
}
