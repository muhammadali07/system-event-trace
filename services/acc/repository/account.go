package repository

import (
	"time"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/sirupsen/logrus"
)

func (r *Accountepository) InsertNewAccount(req models.Account) (err error) {
	err = r.db.Create(req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error":      err.Error(),
			"id":         req.ID,
			"nama":       req.Nama,
			"nik":        req.Nik,
			"no_hp":      req.NoHp,
			"pin":        req.Pin,
			"created_at": time.Time.GoString(time.Now()),
		}).Error("create new account data failed")
	}
	return
}
