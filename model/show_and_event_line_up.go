package model

import "github.com/google/uuid"

type ShowAndEventLineUp struct {
	ShowID   uuid.UUID `json:"showId"`
	MemberID uuid.UUID `json:"memberId"`
	Member   Member    `gorm:"->"`
}
