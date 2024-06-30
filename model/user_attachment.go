package model

import "github.com/google/uuid"

type UserAttachment struct {
	ID     uuid.UUID
	UserID uuid.UUID
	URL    string
	Remark string //KTP, PASSPORT, SIM, OFC_CARD, etc
}
