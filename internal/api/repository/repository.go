package apirepository

import (
	"context"
	"kp/internal/api"
	"kp/internal/entity"
	"kp/internal/model"
	"kp/pkg/exception"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) api.Repository {
	return &Repository{db: db}
}

func (r *Repository) withTrx(tx *gorm.DB) *Repository {
	return &Repository{
		db: tx,
	}
}

func (r *Repository) Trx(ctx context.Context, fn func(repo api.Repository) error) error {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return exception.NewDBQueryError(err)
	}

	repo := r.withTrx(tx)
	err := fn(repo)
	if err != nil {
		tx.Rollback()
		return exception.NewDBQueryError(err)
	}

	return tx.Commit().Error
}

func (r *Repository) GetListCustomers(ctx context.Context) ([]*entity.Customers, error) {
	var data []*entity.Customers

	res := r.db.Find(&data)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return data, nil
}

func (r *Repository) GetCustomer(ctx context.Context, id int64) (*entity.Customers, error) {
	var data entity.Customers

	res := r.db.First(&data, id)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return &data, nil
}

func (r *Repository) CreateCustomer(ctx context.Context, data *entity.Customers) error {
	res := r.db.Create(data)
	if err := res.Error; err != nil {
		return err
	}

	if res.RowsAffected == 0 {
		return exception.NewDBQueryError(nil, "failed to insert data")
	}

	return nil
}

func (r *Repository) GetAccountCustomer(ctx context.Context, accountID int64) (*entity.Customers, error) {
	var data *entity.Customers

	res := r.db.
		Joins("join accounts on accounts.customer_id = customers.id").
		Where("accounts.id = ?", accountID).
		First(&data)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	if data == nil {
		return nil, exception.NewNotFoundError("repository", nil)
	}

	return data, nil
}

func (r *Repository) GetListAccounts(ctx context.Context) ([]*entity.Accounts, error) {
	var data []*entity.Accounts

	res := r.db.Find(&data)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return data, nil
}

func (r *Repository) GetAccount(ctx context.Context, id int64) (*entity.Accounts, error) {
	var data entity.Accounts

	res := r.db.First(&data, id)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return &data, nil
}

func (r *Repository) CreateAccount(ctx context.Context, data *entity.Accounts) error {
	res := r.db.Create(data)
	if err := res.Error; err != nil {
		return exception.NewDBQueryError(err)
	}

	if res.RowsAffected == 0 {
		return exception.NewDBQueryError(nil, "failed to insert data")
	}

	return nil
}

func (r *Repository) GetListLimits(ctx context.Context) ([]*entity.Limits, error) {
	var data []*entity.Limits

	res := r.db.Find(&data)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return data, nil
}

func (r *Repository) GetLimit(ctx context.Context, id int64) (*entity.Limits, error) {
	var data entity.Limits

	res := r.db.First(&data, id)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return &data, nil
}

func (r *Repository) CreateLimits(ctx context.Context, data ...*entity.Limits) error {
	res := r.db.Create(data)
	if err := res.Error; err != nil {
		return exception.NewDBQueryError(err)
	}

	if res.RowsAffected == 0 {
		return exception.NewDBQueryError(nil, "failed to insert data")
	}

	return nil
}

func (r *Repository) UpdateLimit(ctx context.Context, limits *entity.Limits) error {
	res := r.db.Save(limits)
	if err := res.Error; err != nil {
		return exception.NewDBQueryError(err)
	}

	if res.RowsAffected == 0 {
		return exception.NewDBQueryError(nil, "failed to update data")
	}

	return nil
}

func (r *Repository) GetAccountLimits(ctx context.Context, accountID int64) ([]*model.AccountLimits, error) {
	var data []*model.AccountLimits

	if err := r.db.Model(&entity.Limits{}).
		Select("limits.id, limits.duration, limits.initial_amount, limits.current_amount, fees.amount interest").
		Joins("join fees on fees.type = limits.duration").
		Where("limits.account_id = ?", accountID).Scan(&data).Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return data, nil
}

func (r *Repository) GetListTransactions(ctx context.Context) ([]*entity.Transactions, error) {
	var data []*entity.Transactions

	res := r.db.Find(&data)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return data, nil
}

func (r *Repository) GetAccountTransactions(ctx context.Context, accountID int64) ([]*entity.Transactions, error) {
	var data []*entity.Transactions

	res := r.db.Find(&data, accountID)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return data, nil
}

func (r *Repository) CreateTransactions(ctx context.Context, data *entity.Transactions) error {
	res := r.db.Save(data)
	if err := res.Error; err != nil {
		return exception.NewDBQueryError(err)
	}

	if res.RowsAffected == 0 {
		return exception.NewDBQueryError(nil, "failed to insert data")
	}

	return nil
}

func (r *Repository) GetListFees(ctx context.Context) ([]*entity.Fees, error) {
	var data []*entity.Fees

	res := r.db.Find(&data)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return data, nil
}

func (r *Repository) GetAdminFee(ctx context.Context) (float64, error) {
	var amount float64

	if err := r.db.Model(&entity.Fees{}).
		Where("type = ?", 0).
		Select("amount").
		Scan(&amount).Error; err != nil {
		return amount, exception.NewDBQueryError(err)
	}

	return amount, nil
}
