package app

import "gorm.io/gorm"

type JournalServicePort interface {
	HandleCashDeposito(payload any) (err error)
	HandleCashWithDraw(payload any) (err error)
	// HandleTransferKliring(payload any) (err error)
	// HandleGetListTransaction(payload any) (err error)
}

type JournalDatastorePort interface {
	Begin() (tx *gorm.DB, err error)
	Rollback(tx *gorm.DB)
	Commit(tx *gorm.DB)
}
