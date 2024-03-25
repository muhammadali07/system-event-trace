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
	err = a.SendMessageToKafka(payload)
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

	a.log.WithFields(logrus.Fields{
		"account_number":      req.AccountNumber,
		"balance_account":     balance.Respdata.(float64),
		"amount_case_deposit": req.Amount,
	}).Info("request payload cash deposit")

	// Store balance for later use
	initialBalance := balance.Respdata.(float64)

	req.Amount += initialBalance

	response = req.Amount

	err = a.repo.TransactionCasDeposithWithDraw(req)
	if err != nil {
		err = fmt.Errorf("failed to create account")
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": req,
		}).Warn(err.Error())
		return
	}

	a.log.WithFields(logrus.Fields{
		"payload":         req,
		"balance_account": balance,
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
	err = a.SendMessageToKafka(payload)
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

	err = a.repo.TransactionCasDeposithWithDraw(payloadCashWithDraw)
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
		"payload": req,
	}).Info("cash withdraw success")

	return
}

func (a *AccountApp) TransferKliring(req models.TransactionKliring) (response float64, err error) {
	if req.AmountKliring < 0 {
		err = fmt.Errorf("amount kliring failed")
		return
	}

	resValidationAccNumber, err := a.repo.GetvalidateAccount(req.AccountNumberDestination)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("account number destinaation not found")
			a.log.WithFields(logrus.Fields{
				"account_number": req,
			}).Warn(err.Error())
			return
		} else {
			err = fmt.Errorf(err.Error())
			return
		}
	}

	if resValidationAccNumber.Status != "A" {
		err = fmt.Errorf("account number destination is not active")
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
			TransactionDesc:          req.DescTransaction,
		},
	}

	err = a.SendMessageToKafka(payload)
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
	account := []string{req.AccountNumberSource, req.AccountNumberDestination}
	var balance []float64
	for _, v := range account {
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
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": req,
		}).Warn(err.Error())
		err = fmt.Errorf(err.Error())
		return
	}

	response = balance[0] - req.AmountKliring

	a.log.WithFields(logrus.Fields{
		"payload": req,
	}).Info("transfer kliring success")

	return
}

func (a *AccountApp) TransactionTransferKliring(req models.TransactionKliring) (err error) {

	a.log.WithFields(logrus.Fields{
		"request": req,
	}).Info("request to transfer transaction")

	if req.BalanceSource < req.AmountKliring {
		err = fmt.Errorf("balance acount number source is not enough")
		return
	}
	// update balance account source
	initialBalanceSource := req.BalanceSource - req.AmountKliring
	payloadSource := models.TransactionDepositWithdraw{
		AccountNumber: req.AccountNumberSource,
		Amount:        initialBalanceSource,
	}
	err = a.repo.TransactionCasDeposithWithDraw(payloadSource)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}

	// update balance account destination
	initialBalanceDestination := req.BalanceDestination + req.AmountKliring
	payloadDestination := models.TransactionDepositWithdraw{
		AccountNumber: req.AccountNumberDestination,
		Amount:        initialBalanceDestination,
	}
	err = a.repo.TransactionCasDeposithWithDraw(payloadDestination)
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
