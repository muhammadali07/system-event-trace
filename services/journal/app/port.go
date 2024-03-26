package app

import (
	"github.com/jackc/pgx/v5"
)

type JournalServicePort interface {
	HandleCashDeposito(payload any) (err error)
	// HandleCashWithDraw(payload any) (err error)
	// HandleTransferKliring(payload any) (err error)
	// HandleGetListTransaction(payload any) (err error)
}

type JournalDatastorePort interface {
	Begin() (tx *pgx.Conn, err error)
	// Rollback(tx *pgx.Conn)
	// Commit(tx *pgx.Conn)
}
