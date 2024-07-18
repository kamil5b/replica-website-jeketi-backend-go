package model

import (
	"time"

	"github.com/google/uuid"
)

type ShowAndEvent struct {
	ID       uuid.UUID            `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title    string               `json:"title"`
	Date     time.Time            `json:"date"`
	Price    uint32               `json:"price"`
	Location string               `json:"location"`
	Quota    uint16               `json:"quota"`
	LineUp   []ShowAndEventLineUp `json:"lineUp" gorm:"foreignKey:ShowID;->"`

	CreatedAt time.Time `json:"-" gorm:"default:now()"`
	UpdatedAt time.Time `json:"-" gorm:"default:now()"`
}
