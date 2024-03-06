package sqlc

import (
	"context"
	"domain"
)

func (q *Queries) CreateAccount(ctx context.Context, req *domain.Account) (*domain.Account, error) {
	arg := CreateAccountParams{
		Nama:  req.Nama,
		NIK:   req.NIK,
		NoHP:  req.NoHP,
		PIN:   req.PIN,
		Saldo: int64(req.Saldo),
	}
	account, err := q.Queries.CreateAccount(ctx, &arg)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (q *Queries) GetAccountByNoRekening(ctx context.Context, noRekening string) (*domain.Account, error) {
	account, err := q.Queries.GetAccountByNoRekening(ctx, noRekening)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (q *Queries) UpdateAccountSaldo(ctx context.Context, noRekening string, saldo int) error {
	arg := UpdateAccountSaldoParams{
		Saldo:      int64(saldo),
		NoRekening: noRekening,
	}
	return q.Queries.UpdateAccountSaldo(ctx, &arg)
}
