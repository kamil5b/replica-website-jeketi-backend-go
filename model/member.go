package model

import (
	"time"

	"github.com/gojp/kana"
	"github.com/google/uuid"
)

type Member struct {
	ID           uuid.UUID
	Name         string
	JapaneseName string
	CallName     string
	BloodType    string
	Horoscope    string
	Birthday     time.Time
	Height       uint8
	Status       string //ACTIVE, INACTIVE, GRADUATED, RESIGN, TERMINATED
	Generation   uint8
}

func (m *Member) New(name, callName, bloodType string, birthday time.Time, generation, height uint8) {
	m.Name = name
	m.Birthday = birthday
	m.CallName = callName
	m.Generation = generation
	m.Height = height
	m.BloodType = bloodType

	m.JapaneseName = kana.RomajiToKatakana(m.CallName)
	//TODO: add horoscope

}
