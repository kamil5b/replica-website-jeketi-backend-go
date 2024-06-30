package model

import "github.com/google/uuid"

type ShowAndEventLineUp struct {
	ShowID   uuid.UUID
	MemberID uuid.UUID
}
