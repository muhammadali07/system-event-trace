package models

import "time"

type Journal struct {
	ID                  uint64    `json:"id" db:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	TransactionDate     time.Time `json:"transaction_date" db:"transaction_date"`
	AccountNumberCredit string    `json:"account_number_credit" db:"account_number_credit"`
	AccountNumberDebit  string    `json:"account_number_debit" db:"account_number_debit"`
	AmountCredit        float64   `json:"amount_credit" db:"amount_credit"`
	AmountDebit         float64   `json:"amount_debit" db:"amount_debit"`
	TypeTransaction     string    `json:"type_transaction" db:"type_transaction"`
}
