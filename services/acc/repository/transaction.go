package repository

import (
	"fmt"
	"time"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/sirupsen/logrus"
)

func (r *Accountepository) InsertCashDeposito(req models.CashDeposit) (err error) {
	err = r.db.Create(req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error":      err.Error(),
			"payload":    req,
			"created_at": time.Time.GoString(time.Now()),
		}).Error("create new account data failed")
	}
	return
}

func (r *Accountepository) GetAccountBalance(req string) (response float64, err error) {
	var account models.Account
	err = r.db.First(&account, req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"request": req,
		}).Error("query account balance data failed")

		remark := "data with account number not found"
		err = fmt.Errorf(remark)
		return
	}
	return
}
