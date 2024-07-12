package model

import (
	"database/sql"
	"time"

	guuid "github.com/google/uuid"
)

type User struct {
	ID              guuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email           string           `json:"email"`
	Phone           string           `json:"phone"`
	Password        string           `json:"-"` // FULLY ENCRYPT
	Name            string           `json:"name"`
	Role            string           `json:"role"` // FAN/MEMBER/STAFF
	Properties      UserProperty     `json:"properties" gorm:"foreignKey:UserID"`
	Sessions        []Session        `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	UserAttachments []UserAttachment `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE"`

	CreatedAt time.Time `json:"-" `
	UpdatedAt time.Time `json:"-"`
}

type UserProperty struct {
	ID     guuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID guuid.UUID `json:"-"`

	MemberNo          int64        `json:"memberNo" gorm:"uniqueIndex"`
	OfcNo             string       `json:"ofcNo" gorm:"uniqueIndex"`
	OfcExpiryDate     sql.NullTime `json:"ofcExpiryDate"`
	Absences          int32        `json:"absences"`
	Birthday          time.Time    `json:"birthday"`
	Gender            string       `json:"gender"`
	IdentityNo        string       `json:"-"` // NO. KTP, PASSPORT, etc (HALF-ENCRYPTED)
	IdentityType      string       `json:"identityType"`
	Poscode           string       `json:"poscode"`
	Address           string       `json:"address"`
	City              string       `json:"city"`
	ProfilePictureURL string       `json:"profilePicture"`
	Barcode           string       `json:"barcode"`

	Status    string       `json:"status"` // ACTIVE / BLACKLIST / INACTIVE / SUSPENDED
	HeldUntil sql.NullTime `json:"-"`      // IF THERE IS SUSPENSION
}
