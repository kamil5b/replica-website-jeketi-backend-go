package model

import "github.com/google/uuid"

type UserAttachment struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID uuid.UUID `json:"userId"`
	URL    string    `json:"url"`
	Remark string    `json:"remark"` //KTP, PASSPORT, SIM, OFC_CARD, etc
}
