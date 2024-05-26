package app

import (
	"github.com/muhammadali07/system-event-trace/services/journal/models"
	"gorm.io/gorm"
)

type JournalServicePort interface {
	HandleCashDeposito(payload map[string]interface{}) (err error)
	HandleCashWithDraw(payload map[string]interface{}) (err error)
	// HandleTransferKliring(payload map[string]interface{}) (err error)
	// HandleGetListTransaction(payload map[string]interface{}) (err error)
}

type JournalDatastorePort interface {
	Begin() (tx *gorm.DB, err error)
	Rollback(tx *gorm.DB)
	Commit(tx *gorm.DB)
	HandleCashDeposito(data *models.Journal) (err error)
	HandleCashWithDraw(data *models.Journal) (err error)
}
