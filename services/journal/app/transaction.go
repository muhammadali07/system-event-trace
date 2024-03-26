package app

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (h *JournalApplication) HandleCashDeposito(payload any) (err error) {
	tx, err := h.datastore.Begin()
	if err != nil {
		h.log.Error(logrus.Fields{"err": err}, nil, err.Error())
		err = fmt.Errorf(err.Error())
		return
	}
	fmt.Println(tx)
	return

}
