package app

import (
	"errors"
	"fmt"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (a *AccountApp) CashDeposit(req models.CashDeposit) (response models.ResponseApp, err error) {

	// get account balance now

	err = a.repo.InsertCashDeposito(req)
	if err != nil {
		err = fmt.Errorf("failed to create account")
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": req,
		}).Warn(err.Error())
	}

	payload := models.ReqSendingKafka{}
	resKafka, err := a.SendMessageToKafka(payload)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": req,
		}).Warn(err.Error())
	}

	a.log.WithFields(logrus.Fields{
		"payload":  payload,
		"response": resKafka,
	})

	a.log.WithFields(logrus.Fields{
		"payload":       req,
		"sending_kafka": resKafka,
	}).Info("cash deposit success")
	return
}

func (a *AccountApp) GetAccountBalance(req string) (response models.ResponseApp, err error) {
	balance, err := a.repo.GetAccountBalance(req)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("account number not found")
		a.log.WithFields(logrus.Fields{
			"account_number": req,
		}).Warn(err.Error())
		return
	} else if err != nil {
		err = fmt.Errorf("failed to get data account number")
		a.log.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Warn(err.Error())
		return
	}

	response.RespCode = "00"
	response.RespMsg = "Success"
	response.Respdata = balance

	return
}
