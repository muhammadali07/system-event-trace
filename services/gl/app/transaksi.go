package app

import (
	"fmt"

	"github.com/muhammadali07/service-grap-go-api/services/gl/models"
	"github.com/sirupsen/logrus"
)

func (a *GLApp) CreateTransaction(req *models.Transaksi) (err error) {
	err = a.repo.InsertActivity(req)
	if err != nil {
		err = fmt.Errorf("failed to create transaction")
		a.log.WithFields(logrus.Fields{
			"error":                 err.Error(),
			"tgl_trx":               req.TanggalTransaksi,
			"nomor_rekening_kredit": req.NoRekeningKredit,
			"nomor_rekening_debit":  req.NoRekeningDebit,
			"nominal_debit":         req.NominalDebit,
			"nominal_kredit":        req.NominalKredit,
		}).Warn(err.Error())
	} else {
		a.log.WithFields(logrus.Fields{
			"error":                 err.Error(),
			"tgl_trx":               req.TanggalTransaksi,
			"nomor_rekening_kredit": req.NoRekeningKredit,
			"nomor_rekening_debit":  req.NoRekeningDebit,
			"nominal_debit":         req.NominalDebit,
			"nominal_kredit":        req.NominalKredit,
		}).Info("create transaction success")
	}
	return
}
