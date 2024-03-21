package repository

import (
	"time"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

func (r *Accountepository) InsertNewAccount(req *models.Account) (err error) {
	err = r.db.Create(req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error":        err.Error(),
			"id":           req.ID,
			"name":         req.Name,
			"nik":          req.NIK,
			"phone_number": req.PhoneNumber,
			"pin":          req.Pin,
			"created_at":   time.Time.GoString(time.Now()),
		}).Error("create new account data failed")
	}
	return
}

func (r *Accountepository) GetAccountNumber(req models.ReqGetAccountNumber) (response models.Account, err error) {
	// err = r.db.Where("nik = ?", req.NIK).Or("no_hp = ?", req.PhoneNumber).First(&account).Error
	err = r.db.Preload(clause.Associations).First(&response, "phone_number = ?", req.PhoneNumber).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"payload": req,
			"err":     err.Error(),
		})
		return response, err
	}
	return response, nil
}
