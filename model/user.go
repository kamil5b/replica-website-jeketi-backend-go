package model

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              guuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email           string           `json:"email"`
	Phone           string           `json:"phone"`
	Password        string           `json:"-"` // FULLY ENCRYPT
	Name            string           `json:"name"`
	Role            string           `json:"role"` // FAN/MEMBER/STAFF
	Properties      *UserProperty    `json:"properties" gorm:"foreignKey:UserID"`
	UserAttachments []UserAttachment `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE;->" json:"userAttachments"`

	CreatedAt time.Time      `json:"-" gorm:"default:now()"`
	UpdatedAt time.Time      `json:"-" gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type UserProperty struct {
	ID     guuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID guuid.UUID `json:"-"`

	MemberNo          int64                     `json:"memberNo" gorm:"uniqueIndex"`
	OfcNo             string                    `json:"ofcNo" gorm:"uniqueIndex"`
	OfcExpiryDate     sql.NullTime              `json:"ofcExpiryDate"`
	Absences          int32                     `json:"absences"`
	Birthday          time.Time                 `json:"birthday"`
	Gender            string                    `json:"gender"`
	IdentityNo        string                    `json:"-"` // NO. KTP, PASSPORT, etc (HALF-ENCRYPTED)
	IdentityType      string                    `json:"identityType"`
	Poscode           string                    `json:"poscode"`
	Address           string                    `json:"address"`
	City              string                    `json:"city"`
	ProfilePictureURL string                    `json:"profilePicture"`
	Barcode           string                    `json:"barcode"`
	ShowHistory       []ShowAndEventUserHistory `json:"showHistory" gorm:"foreignKey:UserID;->"`
	EventHistory      []EventUserHistory        `json:"eventHistory" gorm:"foreignKey:UserID;->"`
	PointBalance      UserJkt48PointBalance     `json:"pointBalance" gorm:"foreignKey:UserID;->"`
	OshimenID         guuid.UUID                `json:"oshimenId"`
	Oshimen           Member                    `json:"oshimen" gorm:"foreignKey:OshimenID;->"`

	Status    string       `json:"status"` // ACTIVE / BLACKLIST / INACTIVE / SUSPENDED
	HeldUntil sql.NullTime `json:"-"`      // IF THERE IS SUSPENSION
}

func GetUserFromCtx(c *fiber.Ctx) *User {
	user, ok := c.Locals("user").(User)
	if ok {
		return &user
	}
	return nil
}
