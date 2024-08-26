package api

import (
	"context"
	"kp/internal/model"
)

type Usecase interface {
	GetAccountProfile(ctx context.Context, accountID int64) (*model.AccountProfile, error)
	GetAccountLimits(ctx context.Context, accountID int64) ([]*model.AccountLimits, error)
	GetAccountTransactions(ctx context.Context, accountID int64) ([]*model.AccountTransaction, error)

	CreateTransaction(ctx context.Context, req *model.RequestCreateTransaction) (*model.AccountTransaction, error)

	ValidateAccountLimit(ctx context.Context, req *model.RequestCreateTransaction) error
}
