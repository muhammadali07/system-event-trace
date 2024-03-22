package models

type TransactionDepositWithdraw struct {
	AccountNumber string  `json:"account_number"`
	Amount        float64 `json:"amount"`
}
