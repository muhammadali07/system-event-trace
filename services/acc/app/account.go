package app

import (
	"fmt"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/sirupsen/logrus"
)

func (a *AccountApp) CreateAccount(req *models.Account) (err error) {
	err = a.repo.InsertNewAccount(req)
	if err != nil {
		err = fmt.Errorf("failed to create account")
		a.log.WithFields(logrus.Fields{
			"error": err.Error(),
			"nama":  req.Nama,
			"nik":   req.Nik,
			"no_hp": req.NoHp,
			"pin":   req.Pin,
		}).Warn(err.Error())
	} else {
		_, err := a.SendMessageToKafka(*req)
		if err != nil {
			a.log.WithFields(logrus.Fields{
				"error":   err.Error(),
				"payload": req,
			}).Warn(err.Error())
		}

		a.log.WithFields(logrus.Fields{
			"nama":  req.Nama,
			"nik":   req.Nik,
			"no_hp": req.NoHp,
			"pin":   req.Pin,
		}).Info("create account success")
	}
	return
}
