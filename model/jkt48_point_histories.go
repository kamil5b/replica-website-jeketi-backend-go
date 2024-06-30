package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Jkt48PointHistory struct {
	ID              uuid.UUID
	BalanceID       uuid.UUID
	UserID          uuid.UUID
	Point           uint32
	BonusPoint      uint32
	CashFlow        string // DEBIT/CREDIT
	Amount          uint8
	TotalPoint      uint64
	TotalBonusPoint uint64
	Remark          string // Top up, Video Call, Handshake, etc
	Status          string // SUCCESS, DECLINE, PENDING
	PaymentID       uuid.UUID
	ExpiredAt       time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (p *Jkt48PointHistory) AfterUpdate(tx *gorm.DB) error {
	balance := new(UserJkt48PointBalance)
	err := tx.Where("id = ?", p.BalanceID).First(balance).Error
	if err != nil {
		return err
	}
	point := int64(p.TotalPoint)
	bonus := int64(p.TotalBonusPoint)
	if p.CashFlow == "CREDIT" {
		point *= -1
		bonus *= -1
	}
	balance.PointBalance += point
	balance.BonusPointBalance += bonus
	return tx.Save(balance).Error
}
