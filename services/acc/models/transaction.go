package models

type TransactionDepositWithdraw struct {
	AccountNumber string  `json:"account_number"`
	Amount        float64 `json:"amount"`
}

type TransactionKliring struct {
	AccountNumberSource      string  `json:"account_number_source"`
	AccountNumberDestination string  `json:"account_number_destination"`
	AmountKliring            float64 `json:"amount_kliring"`
	BalanceSource            float64 `json:"balance_source"`
	BalanceDestination       float64 `json:"balance_destination"`
	DescTransaction          string  `json:"desc_transaction"`
}
