package app

import (
	"fmt"
	"time"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/muhammadali07/service-grap-go-api/services/acc/pkg/utils"
	"github.com/sirupsen/logrus"
)

func (a *AccountApp) CreateAccount(req *models.Account) (response string, err error) {
	// validation account number exist

	valAccountNumber, err := a.GetAccountNumber(models.ReqGetAccountNumber{
		NIK:         req.NIK,
		PhoneNumber: req.PhoneNumber,
	})

	if err != nil {
		remark := fmt.Sprintf("nik or phone_number has already exist with account number %v", valAccountNumber)
		err = fmt.Errorf(remark)
		return
	}

	resGenNomorRekening := utils.GenerateAccountNumber()
	encryptedPin, err := utils.EncryptPin(req.Pin)
	if err != nil {
		err = fmt.Errorf("failed to encrypt pin")
		a.log.WithFields(logrus.Fields{
			"error": err.Error(),
			"pin":   req.Pin,
		}).Warn(err.Error())
	}

	payloadInsert := &models.Account{
		ID:            req.ID,
		Name:          req.Name,
		NIK:           req.NIK,
		PhoneNumber:   req.PhoneNumber,
		Pin:           encryptedPin,
		AccountNumber: resGenNomorRekening,
		Balance:       0,
		CreatedAt:     time.Now(),
	}
	err = a.repo.InsertNewAccount(payloadInsert)
	if err != nil {
		err = fmt.Errorf("failed to create account")
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": req,
		}).Warn(err.Error())
	}

	response = resGenNomorRekening
	a.log.WithFields(logrus.Fields{
		"name":         req.Name,
		"nik":          req.NIK,
		"phone_number": req.PhoneNumber,
		"pin":          req.Pin,
	}).Info("create account success")
	return
}

func (a *AccountApp) GetAccountNumber(req models.ReqGetAccountNumber) (response string, err error) {
	res, err := a.repo.GetAccountNumber(req)
	if err != nil {
		err = fmt.Errorf("failed to create account")
		a.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"payload": req,
		}).Warn(err.Error())
	}

	response = res.AccountNumber
	a.log.WithFields(logrus.Fields{"account_number": res}).Info("get account number success")
	return
}
