package model

import (
	"time"

	"github.com/google/uuid"
)

type Jkt48PointHistory struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Point      uint32
	BonusPoint uint32
	CashFlow   string // DEBIT/CREDIT
	Amount     uint8
	Remark     string // Top up, Video Call, Handshake, etc
	Status     string // SUCCESS, DECLINE, PENDING
	PaymentID  uuid.UUID
	ExpiredAt  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
