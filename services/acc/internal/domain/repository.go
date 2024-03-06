package domain

import "context"

// Repository defines the interface for interacting with the data store
type Repository interface {
	// CreateAccount creates a new account in the data store
	CreateAccount(ctx context.Context, account *Account) error

	// GetAccount retrieves an account by its noRekening
	GetAccount(ctx context.Context, noRekening string) (*Account, error)

	// UpdateAccount updates an existing account in the data store
	UpdateAccount(ctx context.Context, account *Account) error
}
