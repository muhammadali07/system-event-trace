package app

import (
	"fmt"
	"time"

	"github.com/muhammadali07/system-event-trace/services/journal/models"
	"github.com/sirupsen/logrus"
)

func (h *JournalApplication) HandleCashDeposito(payload map[string]interface{}) (err error) {
	_, err = h.datastore.Begin()
	if err != nil {
		h.log.Error(logrus.Fields{"err": err}, nil, err.Error())
		err = fmt.Errorf(err.Error())
		return
	}
	// id := utils.GetNextNumber()
	paramInsert := models.Journal{
		// ID:                  int64(id),
		AccountNumberCredit: payload["account_number_credit"].(string),
		AmountCredit:        payload["amount_credit"].(float64),
		TransactionDate:     time.Now().UTC(),
		TypeTransaction:     payload["transaction_type"].(string),
	}
	err = h.datastore.HandleCashDeposito(&paramInsert)
	if err != nil {
		h.log.Error(logrus.Fields{"err": err}, nil, err.Error())
		err = fmt.Errorf(err.Error())
		return
	}
	return

}

func (h *JournalApplication) HandleCashWithDraw(payload map[string]interface{}) (err error) {
	_, err = h.datastore.Begin()
	if err != nil {
		h.log.Error(logrus.Fields{"err": err}, nil, err.Error())
		err = fmt.Errorf(err.Error())
		return
	}
	// id := utils.GetNextNumber()
	paramInsert := models.Journal{
		// ID:                  int64(id),
		AccountNumberCredit: payload["account_number_debit"].(string),
		AmountCredit:        payload["amount_debit"].(float64),
		TransactionDate:     time.Now().UTC(),
		TypeTransaction:     payload["transaction_type"].(string),
	}
	err = h.datastore.HandleCashDeposito(&paramInsert)
	if err != nil {
		h.log.Error(logrus.Fields{"err": err}, nil, err.Error())
		err = fmt.Errorf(err.Error())
		return
	}
	return

}
