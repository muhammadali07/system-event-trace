package domain

import (
	"context"
	"errors"
	"fmt"
)

type ContextKey string

const (
	ContextKeyUser ContextKey = "user"
)

type Account struct {
	NoRekening string `json:"noRekening"`
	Nama       string `json:"nama"`
	NIK        string `json:"nik"`
	NoHP       string `json:"noHp"`
	PIN        string `json:"-"` // PIN tidak diekspos
	Saldo      int    `json:"saldo"`
}

type registerRequest struct {
	Nama string `json:"nama" binding:"required"`
	NIK  string `json:"nik" binding:"required"`
	NoHP string `json:"noHp" binding:"required"`
	PIN  string `json:"pin" binding:"required"`
}

type depositRequest struct {
	NoRekening string `json:"noRekening" binding:"required"`
	Nominal    int    `json:"nominal" binding:"required,gt=0"`
}

type withdrawRequest struct {
	NoRekening string `json:"noRekening" binding:"required"`
	Nominal    int    `json:"nominal" binding:"required,gt=0"`
}

type transferRequest struct {
	NoRekeningAsal   string `json:"noRekeningAsal" binding:"required"`
	NoRekeningTujuan string `json:"noRekeningTujuan" binding:"required"`
	Nominal          int    `json:"nominal" binding:"required,gt=0"`
}

type registerResponse struct {
	NoRekening string `json:"noRekening"`
}

type depositResponse struct {
	Saldo int `json:"saldo"`
}

type errorResponse struct {
	Remark string `json:"remark"`
}

var (
	ErrDuplicateNIK      = errors.New("NIK sudah terdaftar")
	ErrDuplicateNoHP     = errors.New("Nomor HP sudah terdaftar")
	ErrAccountNotFound   = errors.New("Nomor rekening tidak ditemukan")
	ErrInsufficientSaldo = errors.New("Saldo tidak mencukupi")
)

// UseCase defines the interface for account operations
type UseCase interface {
	CreateAccount(ctx context.Context, nama, nik, noHp, pin string) (*Account, error)
	Deposit(ctx context.Context, noRekening string, nominal int) (int, error)
	Withdraw(ctx context.Context, noRekening string, nominal int) (int, error)
	Transfer(ctx context.Context, req transferRequest) (int, error)
	GetSaldo(ctx context.Context, noRekening string) (int, error)
	GetMutasi(ctx context.Context, noRekening string) ([]map[string]interface{}, error)
}

// ValidateNoRekening checks if the provided noRekening is valid (e.g., not empty)
func ValidateNoRekening(noRekening string) error {
	if noRekening == "" {
		return fmt.Errorf("noRekening is required")
	}
	return nil
}
