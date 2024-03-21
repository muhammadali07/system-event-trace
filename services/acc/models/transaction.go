package models

type CashDeposit struct {
	NomorRekening string  `json:"nomor_rekening"`
	Nominal       float64 `json:"nominal"`
}
