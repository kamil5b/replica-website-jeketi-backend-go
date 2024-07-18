package model

import (
	"time"

	"github.com/google/uuid"
)

type EventLineUp struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	MemberID  uuid.UUID `json:"memberId"`
	Member    Member    `json:"member" gorm:"->"`
	EventID   uuid.UUID `json:"eventId"`
	Quota     uint16    `json:"quota"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Line      uint16    `json:"line"`
	Session   uint8     `json:"session"`
}
