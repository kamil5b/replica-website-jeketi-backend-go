package model

import (
	"time"

	"github.com/google/uuid"
)

type EventLineUp struct {
	ID       uuid.UUID
	MemberID uuid.UUID
	EventID  uuid.UUID
	Quota    uint16
	Time     time.Time
	Session  uint8
}
