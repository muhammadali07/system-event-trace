package repository

import "github.com/muhammadali07/system-event-trace/services/acc/models"

func (r *Accountepository) VerifyAccount(account_no string, pin string) (err error) {
	err = r.db.Where("account_no = ? AND pin = ?", account_no, pin).First(&models.Account{}).Error
	if err != nil {
		return err
	}
	return
}
