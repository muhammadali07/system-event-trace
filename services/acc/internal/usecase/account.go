package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	model "github.com/muhammadali07/service-grap-go-api/services/acc/internal/domain"
)

// AccountUseCase represents the use case for account management.
type AccountUseCase struct {
	db *sqlx.DB
}

// NewAccountUseCase creates a new instance of AccountUseCase.
func NewAccountUseCase(db *sqlx.DB) *AccountUseCase {
	return &AccountUseCase{
		db: db,
	}
}

// CreateAccount creates a new account.
func (uc *AccountUseCase) CreateAccount(ctx context.Context, req model.RegisterRequest) (res model.Account, err error) {
	// Validate request.

	// Check if account with NIK already exists.
	existingAccount, err := uc.GetAccountByNIK(ctx, req.NIK)
	if err != nil {
		return res, err
	}
	if existingAccount == res {
		return res, errors.New("account with NIK already exists")
	}

	// Create new account in database.
	account, err := uc.createAccountInDB(ctx, req)
	if err != nil {
		return res, err
	}

	return account, nil
}

// GetAccountByNIK gets an account by NIK.
func (uc *AccountUseCase) GetAccountByNIK(ctx context.Context, nik string) (model.Account, error) {
	// Get account from database.
	account, err := uc.getAccountByNIKFromDB(ctx, nik)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (uc *AccountUseCase) createAccountInDB(ctx context.Context, req model.RegisterRequest) (res model.Account, err error) {
	// Prepare SQL statement.
	query := `INSERT INTO accounts (nama, nik, no_hp, pin, saldo) VALUES ($1, $2, $3, $4, $5) RETURNING no_rekening, nama, nik, no_hp, pin, 0`

	// Execute query and get result.
	row := uc.db.QueryRowContext(ctx, query, req.Nama, req.NIK, req.NoHP, req.PIN, 0)

	// Scan the result into an Account struct.
	var acc model.Account
	err = row.Scan(&acc.NoRekening, &acc.Nama, &acc.NIK, &acc.NoHP, &acc.PIN, &acc.Saldo)
	if err != nil {
		return res, fmt.Errorf("error scanning result: %w", err)
	}

	return res, nil
}

func (uc *AccountUseCase) getAccountByNIKFromDB(ctx context.Context, nik string) (res model.Account, err error) {
	// Prepare SQL statement.
	query := `SELECT no_rekening, nama, nik, no_hp, pin, saldo FROM accounts WHERE nik = $1`

	// Execute query and get result.
	row := uc.db.QueryRowContext(ctx, query, nik)

	// Scan the result into an Account struct.
	var acc model.Account
	err = row.Scan(&acc.NoRekening, &acc.Nama, &acc.NIK, &acc.NoHP, &acc.PIN, &acc.Saldo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, fmt.Errorf("error scanning result: %w", err)
	}

	return res, nil
}
