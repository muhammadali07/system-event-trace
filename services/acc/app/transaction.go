package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (a *AccountApp) CashDeposit(req models.TransactionDepositWithdraw) (response float64, err error) {
	// journal to kafka
	payload := models.ReqSendingKafka{
		Topic: "cash_deposit",
		Data: models.JournalKafka{
			TransactionDate:     time.Now(),
			AccountNumberCredit: req.AccountNumber,
			AmountCredit:        req.Amount,
			TransactionType:     "C",
		},
	}
	resKafka, err := a.SendMessageToKafka(payload)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": payload,
		}).Warn("sending message to kafka failed")
		remark := "cash deposit failed"
		err = fmt.Errorf(remark)
		return
	}

	// get account balance now
	balance, err := a.GetAccountBalance(req.AccountNumber)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}

	// Store balance for later use
	initialBalance := balance.Respdata.(float64)

	req.Amount += initialBalance

	err = a.repo.TransactionCashWithDraw(req)
	if err != nil {
		err = fmt.Errorf("failed to create account")
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": req,
		}).Warn(err.Error())
		return
	}

	response = initialBalance + req.Amount

	a.log.WithFields(logrus.Fields{
		"payload":       req,
		"sending_kafka": resKafka,
	}).Info("cash deposit success")

	return
}

func (a *AccountApp) CashWithDraw(req models.TransactionDepositWithdraw) (response float64, err error) {
	// journal to kafka
	payload := models.ReqSendingKafka{
		Topic: "cash_withdraw",
		Data: models.JournalKafka{
			TransactionDate:     time.Now(),
			AccountNumberCredit: req.AccountNumber,
			AmountCredit:        req.Amount,
			TransactionType:     "C",
		},
	}
	resKafka, err := a.SendMessageToKafka(payload)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": payload,
		}).Warn("sending message to kafka failed")
		remark := "cash withdraw failed"
		err = fmt.Errorf(remark)
		return
	}

	// get account balance now
	balance, err := a.GetAccountBalance(req.AccountNumber)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}

	// Store balance for later use
	initialBalance := balance.Respdata.(float64)
	initialBalance -= req.Amount

	payloadCashWithDraw := models.TransactionDepositWithdraw{
		AccountNumber: req.AccountNumber,
		Amount:        initialBalance,
	}

	err = a.repo.TransactionCashWithDraw(payloadCashWithDraw)
	if err != nil {
		err = fmt.Errorf("failed to create account")
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": req,
		}).Warn(err.Error())
		return
	}

	response = initialBalance + req.Amount

	a.log.WithFields(logrus.Fields{
		"payload":       req,
		"sending_kafka": resKafka,
	}).Info("cash withdraw success")

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
