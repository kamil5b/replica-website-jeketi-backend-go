package request

import (
	"time"

	guuid "github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email            string     `json:"email"`
	Phone            string     `json:"phone"`
	Password         string     `json:"password"`
	Name             string     `json:"name"`
	Birthday         time.Time  `json:"birthday"`
	Gender           string     `json:"gender"`
	IdentityNo       string     `json:"-"` // NO. KTP, PASSPORT, etc (HALF-ENCRYPTED)
	IdentityType     string     `json:"identityType"`
	Poscode          string     `json:"poscode"`
	Address          string     `json:"address"`
	City             string     `json:"city"`
	OshimenID        guuid.UUID `json:"oshimenId"`
	ProfilePictureID guuid.UUID `json:"profilePictureId"`
}
