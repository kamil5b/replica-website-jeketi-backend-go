package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title          string    `json:"title"`
	EventType      string    `json:"eventType"` // 2SHOT, HANDSHAKE, MEET_AND_GREET, VIDEO_CALL
	Date           time.Time `json:"date"`
	PricePerTicket uint32    `json:"pricePerTicket"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`

	LineUp []EventLineUp `json:"lineUp" gorm:"foreignKey:EventID"`
}
