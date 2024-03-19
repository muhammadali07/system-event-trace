package models

import (
	"time"
)

type Account struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Nama      string    `json:"nama"`
	Nik       string    `json:"nik"`
	NoHp      string    `json:"no_hp"`
	Pin       string    `json:"pin"`
	CreatedAt time.Time `json:"created_at"`
}
