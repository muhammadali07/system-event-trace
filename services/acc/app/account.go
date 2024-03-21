package app

import (
	"fmt"
	"time"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/muhammadali07/service-grap-go-api/services/acc/pkg/utils"
	"github.com/sirupsen/logrus"
)

func (a *AccountApp) CreateAccount(req *models.Account) (response string, err error) {
	resGenNomorRekening := utils.GenerateAccountNumber()
	encryptedPin, err := utils.EncryptPin(req.Pin)
	if err != nil {
		err = fmt.Errorf("failed to encrypt pin")
		a.log.WithFields(logrus.Fields{
			"error": err.Error(),
			"pin":   req.Pin,
		}).Warn(err.Error())
	}

	payloadInsert := models.Account{
		ID:            req.ID,
		Nama:          req.Nama,
		Nik:           req.Nik,
		NoHp:          req.NoHp,
		Pin:           encryptedPin,
		NomorRekening: resGenNomorRekening,
		Saldo:         0,
		CreatedAt:     time.Time{},
	}
	err = a.repo.InsertNewAccount(payloadInsert)
	if err != nil {
		err = fmt.Errorf("failed to create account")
		a.log.WithFields(logrus.Fields{
			"error": err.Error(),
			"nama":  req.Nama,
			"nik":   req.Nik,
			"no_hp": req.NoHp,
			"pin":   req.Pin,
		}).Warn(err.Error())
	}

	// payload := models.ReqSendingKafka{}
	// resKafka, err := a.SendMessageToKafka(payload)
	// if err != nil {
	// 	a.log.WithFields(logrus.Fields{
	// 		"error":   err.Error(),
	// 		"payload": req,
	// 	}).Warn(err.Error())
	// }

	// a.log.WithFields(logrus.Fields{
	// 	"payload":  payload,
	// 	"response": resKafka,
	// })

	response = resGenNomorRekening
	a.log.WithFields(logrus.Fields{
		"nama":  req.Nama,
		"nik":   req.Nik,
		"no_hp": req.NoHp,
		"pin":   req.Pin,
	}).Info("create account success")
	return
}
