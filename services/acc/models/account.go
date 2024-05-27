package models

import (
	"time"
)

type Account struct {
	ID            uint      `json:"id" db:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name          string    `json:"name"`
	NIK           string    `json:"nik"`
	PhoneNumber   string    `json:"phone_number"`
	Pin           string    `json:"pin"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type ReqGetAccountNumber struct {
	NIK         string
	PhoneNumber string
}
