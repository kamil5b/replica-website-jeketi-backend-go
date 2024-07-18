package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Jkt48PointHistory struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	BalanceID       uuid.UUID `json:"balanceId"`
	UserID          uuid.UUID `json:"userId"`
	Point           uint32    `json:"point"`
	BonusPoint      uint32    `json:"bonusPoint"`
	CashFlow        string    `json:"cashFlow"` // DEBIT/CREDIT
	Amount          uint8     `json:"amount"`
	TotalPoint      uint64    `json:"totalPoint"`
	TotalBonusPoint uint64    `json:"totalBonusPoint"`
	Remark          string    `json:"remark"` // Top up, Video Call, Handshake, etc
	Status          string    `json:"status"` // SUCCESS, DECLINE, PENDING
	PaymentID       uuid.UUID `json:"paymentId"`
	ExpiredAt       time.Time `json:"expiredAt"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
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
