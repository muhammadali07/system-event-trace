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

func (a *AccountApp) TransferKliring(req models.TransactionKliring) (response float64, err error) {
	if req.AmountKliring < 0 {
		err = fmt.Errorf("amount kliring failed")
		return
	}

	_, err = a.repo.GetvalidateAccount(req.AccountNumberDestination)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("account number not found")
		a.log.WithFields(logrus.Fields{
			"account_number": req,
		}).Warn(err.Error())
		return
	}

	// journal to kafka
	payload := models.ReqSendingKafka{
		Topic: "transfer_kliring",
		Data: models.KliringKafka{
			TransactionDate:          time.Now(),
			AccountNumberSource:      req.AccountNumberSource,
			AccountNumberDestination: req.AccountNumberDestination,
			AmountKliring:            req.AmountKliring,
			TransactionType:          "K",
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

	// get account balance account number source
	var acc []string
	acc = append(acc, req.AccountNumberSource)
	acc = append(acc, req.AccountNumberDestination)

	var balance []float64
	for _, v := range acc {
		resBalance, errors := a.GetAccountBalance(v)
		if errors != nil {
			err = fmt.Errorf(errors.Error())
			return
		}
		balance = append(balance, resBalance.Respdata.(float64))
	}

	// // Store balance for later use
	payloadKliring := models.TransactionKliring{
		AccountNumberSource:      req.AccountNumberSource,
		AccountNumberDestination: req.AccountNumberDestination,
		AmountKliring:            req.AmountKliring,
		BalanceSource:            balance[0],
		BalanceDestination:       balance[1],
	}

	err = a.TransactionTransferKliring(payloadKliring)
	if err != nil {
		err = fmt.Errorf("failed to transfer kliring")
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": req,
		}).Warn(err.Error())
		return
	}

	// response = initialBalance

	a.log.WithFields(logrus.Fields{
		"payload":       req,
		"sending_kafka": resKafka,
	}).Info("transfer kliring success")

	return
}

func (a *AccountApp) TransactionTransferKliring(req models.TransactionKliring) (err error) {

	if req.BalanceSource < req.AmountKliring {
		err = fmt.Errorf("balance account not enough")
		return
	}
	// update balance account source
	initialBalanceSource := req.BalanceSource - req.AmountKliring
	payloadSource := models.TransactionDepositWithdraw{
		AccountNumber: req.AccountNumberSource,
		Amount:        initialBalanceSource,
	}
	err = a.repo.TransactionCashWithDraw(payloadSource)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}

	// update balance account destination
	initialBalanceDestination := req.BalanceDestination + req.AmountKliring
	payloadDestination := models.TransactionDepositWithdraw{
		AccountNumber: req.AccountNumberSource,
		Amount:        initialBalanceDestination,
	}
	err = a.repo.TransactionCashWithDraw(payloadDestination)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}

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
