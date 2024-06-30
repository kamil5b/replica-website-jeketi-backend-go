package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type NonJkt48Event struct {
	ID        uuid.UUID
	Title     string
	Date      time.Time
	URL       sql.NullString
	CreatedAt time.Time
}

type News struct {
	ID        uuid.UUID
	Title     string
	Content   string
	CreatedAt time.Time
}
type OfcContent struct {
	ID        uuid.UUID
	Title     string
	Remark    string //LUMINOR, WHITE ALBUM, OFC ALBUM, etc
	URL       string
	CreatedAt time.Time
}
type HomePageContent struct {
	ID           uuid.UUID
	Title        string
	URL          sql.NullString
	Remark       string //VIDEOS, RELEASE, etc
	Description  string
	ThumbnailURL sql.NullString
	CreatedAt    time.Time
}
