package models

import "time"

type ReqSendingKafka struct {
	Topic string `json:"topic"`
	Data  any    `json:"data"`
}

type RespSendingKafka struct {
	RespCode string `json:"resp_code"`
	RespMsg  string `json:"resp_msg"`
	RespData string `json:"resp_data"`
}

type JournalKafka struct {
	TransactionDate     time.Time `json:"transaction_date"`
	AccountNumberCredit string    `json:"account_number_credit"`
	AmountCredit        float64   `json:"amount_credit"`
	TransactionType     string    `json:"transaction_type"`
	AccountNumberDebit  string    `json:"account_number_debit"`
	AmountDebit         float64   `json:"amount_debit"`
	TransactionDesc     string    `json:"transaction_desc"`
}

type KliringKafka struct {
	AccountNumberSource      string    `json:"account_number_source"`
	AccountNumberDestination string    `json:"account_number_destination"`
	AmountKliring            float64   `json:"amount_kliring"`
	TransactionDate          time.Time `json:"transaction_date"`
	TransactionType          string    `json:"transaction_type"`
	TransactionDesc          string    `json:"transaction_desc"`
}
