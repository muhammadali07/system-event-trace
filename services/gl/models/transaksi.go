package models

import (
	"database/sql"
	"time"
)

type Transaksi struct {
	ID               uint            `json:"id" gorm:"primarykey"`
	TanggalTransaksi sql.NullTime    `json:"tanggal_transaksi"`
	NoRekeningKredit sql.NullString  `json:"no_rekening_kredit"`
	NoRekeningDebit  sql.NullString  `json:"no_rekening_debit"`
	NominalKredit    sql.NullFloat64 `json:"nominal_kredit"`
	NominalDebit     sql.NullFloat64 `json:"nominal_debit"`
	CreatedAt        time.Time       `json:"created_at"`
}
