package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type NonJkt48Event struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title     string         `json:"title"`
	Date      time.Time      `json:"date"`
	URL       sql.NullString `json:"url"`
	CreatedAt time.Time      `json:"createdAt"`
}

type News struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
type OfcContent struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title     string    `json:"title"`
	Remark    string    `json:"remark"` //LUMINOR, WHITE ALBUM, OFC ALBUM, etc
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
}
type HomePageContent struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title        string         `json:"title"`
	URL          sql.NullString `json:"url"`
	Remark       string         `json:"remark"` //VIDEOS, RELEASE, etc
	Description  string         `json:"description"`
	ThumbnailURL sql.NullString `json:"thumbnailUrl"`
	CreatedAt    time.Time      `json:"createdAt"`
}
