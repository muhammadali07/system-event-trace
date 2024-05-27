package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/muhammadali07/system-event-trace/services/acc/models"
	"github.com/muhammadali07/system-event-trace/services/acc/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (a *AccountApp) CreateAccount(ctx context.Context, req *models.Account) (response string, err error) {
	_, span := a.tracer.Start(ctx, fmt.Sprintf("createAccount %s ", "start"))
	defer span.End()
	// validation account number exist

	valAccountNumber, err := a.GetAccountNumber(models.ReqGetAccountNumber{
		NIK:         req.NIK,
		PhoneNumber: req.PhoneNumber,
	})

	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}

	if valAccountNumber.AccountNumber != "" {
		remark := fmt.Sprintf("nik or phone_number has already exist with account number %v", valAccountNumber.AccountNumber)
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
		Name:          req.Name,
		NIK:           req.NIK,
		PhoneNumber:   req.PhoneNumber,
		Pin:           encryptedPin,
		AccountNumber: resGenNomorRekening,
		Balance:       0,
		Status:        "A",
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

	_, span = a.tracer.Start(ctx, fmt.Sprintf("createAccount %s ", "finish"))
	defer span.End()

	utils.LongProcess(a.tracer, ctx)

	return
}

func (a *AccountApp) GetAccountNumber(req models.ReqGetAccountNumber) (response models.Account, err error) {
	response, err = a.repo.GetAccountNumber(req)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("get account number data does not exist")
		a.log.WithFields(logrus.Fields{
			"req": req,
		}).Warn(err.Error())

		return
	} else if err != nil {
		err = fmt.Errorf("failed to get account number")
		a.log.WithFields(logrus.Fields{
			"req": req,
		}).Warn(err.Error())

		return
	}
	a.log.WithFields(logrus.Fields{"data account number": response}).Info("get data account number success")
	return
}
