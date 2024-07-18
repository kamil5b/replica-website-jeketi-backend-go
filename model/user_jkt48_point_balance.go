package model

import "github.com/google/uuid"

type UserJkt48PointBalance struct {
	ID                uuid.UUID           `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID            uuid.UUID           `json:"userId"`
	PointBalance      int64               `json:"pointBalance"`
	BonusPointBalance int64               `json:"bonusPointBalance"`
	Histories         []Jkt48PointHistory `json:"histories" gorm:"foreignKey:BalanceID;->"`
}
