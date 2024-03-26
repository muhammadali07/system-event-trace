// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package datastore

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Journal struct {
	ID                  int32
	TranscationDate     pgtype.Timestamp
	AccountNumberCredit pgtype.Text
	AccountNumberDebit  pgtype.Text
	AmountCredit        pgtype.Numeric
	AmountDebit         pgtype.Numeric
	TypeTransaction     pgtype.Text
	CreatedAt           pgtype.Timestamp
}
