package model

import (
	"time"

	"github.com/abakum/gozodiac"
	"github.com/gojp/kana"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name         string    `json:"name"`
	JapaneseName string    `json:"japaneseName"`
	CallName     string    `json:"callName"`
	BloodType    string    `json:"bloodType"`
	Horoscope    string    `json:"horoscope"`
	Birthday     time.Time `json:"birthday"`
	Height       uint8     `json:"height"`
	Status       string    `json:"status"` //ACTIVE, INACTIVE, GRADUATED, RESIGN, TERMINATED
	Generation   uint8     `json:"generation"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *Member) New(name, callName, bloodType string, birthday time.Time, generation, height uint8) {
	m.Name = name
	m.Birthday = birthday
	m.CallName = callName
	m.Generation = generation
	m.Height = height
	m.BloodType = bloodType
	horoscope := gozodiac.GetZodiacSign(birthday)
	m.Horoscope = horoscope[0].String()

	m.JapaneseName = kana.RomajiToKatakana(m.CallName)

}
