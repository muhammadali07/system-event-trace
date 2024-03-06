package repository

import (
	"github.com/muhammadali07/service-grap-go-api/services/gl/models"
	"github.com/sirupsen/logrus"
)

func (r *GLRepository) InsertActivity(req *models.Transaksi) (err error) {
	err = r.db.Create(req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error":                 err.Error(),
			"tgl_trx":               req.TanggalTransaksi,
			"nomor_rekening_kredit": req.NoRekeningKredit,
			"nomor_rekening_debit":  req.NoRekeningDebit,
			"nominal_debit":         req.NominalDebit,
			"nominal_kredit":        req.NominalKredit,
			// "created_at":            time.TimeString(time.Now()),
		}).Error("create transaction data failed")
	}
	return
}
