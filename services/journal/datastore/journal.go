package datastore

import (
	"github.com/muhammadali07/system-event-trace/services/journal/models"
	"github.com/sirupsen/logrus"
)

func (h *JournalDatabase) HandleCashDeposito(data *models.Journal) (err error) {
	err = h.db.Create(data).Error
	if err != nil {
		h.log.Error(logrus.Fields{"err": err.Error()}, nil, "error when handle cash deposito")
	}
	return
}

func (h *JournalDatabase) HandleCashWithDraw(data *models.Journal) (err error) {
	err = h.db.Create(data).Error
	if err != nil {
		h.log.Error(logrus.Fields{"err": err.Error()}, nil, "error when handle cash deposito")
	}
	return
}
