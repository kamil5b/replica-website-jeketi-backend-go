package model

import (
	"time"

	"github.com/google/uuid"
)

type ShowAndEventUserHistory struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID         uuid.UUID `json:"userId"`
	ShowAndEventID uuid.UUID `json:"showAndEventId"`
	TicketType     string    `json:"ticketType"` // GEN, OFC, VIP
	ShowType       string    `json:"showType"`   // THEATRE, OFC_EVENT, MINI_SHOW
	Time           time.Time `json:"time"`
	Status         string    `json:"status"` // LOSE, WIN, PENDING
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
