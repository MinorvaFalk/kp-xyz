package api

import (
	"context"
	"kp/internal/entity"
	"kp/internal/model"
)

type Repository interface {
	Trx(ctx context.Context, fn func(repo Repository) error) error

	GetListCustomers(ctx context.Context) ([]*entity.Customers, error)
	GetCustomer(ctx context.Context, id int64) (*entity.Customers, error)
	GetAccountCustomer(ctx context.Context, accountID int64) (*entity.Customers, error)
	CreateCustomer(ctx context.Context, data *entity.Customers) error

	GetListAccounts(ctx context.Context) ([]*entity.Accounts, error)
	GetAccount(ctx context.Context, id int64) (*entity.Accounts, error)

	GetListLimits(ctx context.Context) ([]*entity.Limits, error)
	GetLimit(ctx context.Context, id int64) (*entity.Limits, error)
	GetAccountLimits(ctx context.Context, accountID int64) ([]*model.AccountLimits, error)
	UpdateLimit(ctx context.Context, limits *entity.Limits) error

	GetListTransactions(ctx context.Context) ([]*entity.Transactions, error)
	GetAccountTransactions(ctx context.Context, accountID int64) ([]*entity.Transactions, error)
	CreateTransactions(ctx context.Context, data *entity.Transactions) error

	GetListFees(ctx context.Context) ([]*entity.Fees, error)
	GetAdminFee(ctx context.Context) (float64, error)
}
