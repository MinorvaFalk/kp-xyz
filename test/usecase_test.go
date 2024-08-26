package test

import (
	"context"
	"kp/config"
	apiusecase "kp/internal/api/usecase"
	"kp/internal/entity"
	"kp/internal/model"
	"kp/mocks/internal_/api"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	now = time.Now()

	cust1 = &entity.Customers{
		ID:          1,
		Nik:         "9191280503011873",
		FullName:    "Catarina Dziwisz",
		LegalName:   "Catarina Dziwisz",
		BirthPlace:  "China",
		BirthDate:   "1990-07-31",
		Salary:      6219052,
		KtpPhoto:    "http://dummyimage.com/136x100.png/cc0000/ffffff",
		SelfiePhoto: "http://dummyimage.com/250x100.png/cc0000/ffffff",
	}

	accCust1 = &model.AccountProfile{
		ID:          cust1.ID,
		Nik:         cust1.Nik,
		FullName:    cust1.FullName,
		LegalName:   cust1.LegalName,
		BirthPlace:  cust1.BirthPlace,
		BirthDate:   cust1.BirthDate,
		Salary:      cust1.Salary,
		KtpPhoto:    cust1.KtpPhoto,
		SelfiePhoto: cust1.SelfiePhoto,
	}

	accLimit1 = []*model.AccountLimits{
		{
			ID:            1,
			Duration:      1,
			InitialAmount: 1000000,
			CurrentAmount: 1000000,
			Interest:      1.25,
		},
		{
			ID:            2,
			Duration:      2,
			InitialAmount: 2000000,
			CurrentAmount: 2000000,
			Interest:      1.75,
		},
		{
			ID:            3,
			Duration:      3,
			InitialAmount: 3000000,
			CurrentAmount: 3000000,
			Interest:      2.25,
		},
		{
			ID:            4,
			Duration:      4,
			InitialAmount: 4000000,
			CurrentAmount: 4000000,
			Interest:      2.5,
		},
	}

	trAcc1 = []*entity.Transactions{
		{
			ID:             1,
			AccountID:      1,
			ContractNumber: "996f086c-4ad1-4935-b93d-332a6611d14c",
			AssetName:      "Testing",
			Otr:            100000,
			AdminFee:       0,
			TotalPayment:   110000,
			Installment:    27500,
			Interest:       10000,
			CreatedAt:      &now,
		},
	}

	accTr1 = []*model.AccountTransaction{
		{
			ContractNumber: "996f086c-4ad1-4935-b93d-332a6611d14c",
			AssetName:      "Testing",
			Otr:            100000,
			TotalPayment:   110000,
			Interest:       10000,
			Installment:    27500,
			AdminFee:       0,
		},
	}
)

func init() {
	config.InitConfig("../.env")
}

func TestUsecase_GetAccountProfile(t *testing.T) {

	repo := new(api.Repository)
	repo.On("GetAccountCustomer", context.Background(), int64(1)).Return(cust1, nil)
	repo.On("GetAccountCustomer", context.Background(), int64(2)).Return(nil, nil)

	type args struct {
		ctx       context.Context
		accountID int64
	}
	tests := []struct {
		name     string
		args     args
		expected *model.AccountProfile
		wantErr  bool
	}{
		{
			name:     "get customer with account_id 1",
			args:     args{ctx: context.Background(), accountID: 1},
			expected: accCust1,
			wantErr:  false,
		},
		{
			name:     "get customer with account_id 2",
			args:     args{ctx: context.Background(), accountID: 2},
			expected: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := apiusecase.New(repo)
			res, err := u.GetAccountProfile(tt.args.ctx, tt.args.accountID)

			assert.Equal(t, tt.expected, res)

			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestUsecase_GetAccountLimits(t *testing.T) {
	repo := new(api.Repository)
	repo.On("GetAccountLimits", context.Background(), int64(1)).Return(accLimit1, nil)
	repo.On("GetAccountLimits", context.Background(), int64(2)).Return(nil, nil)

	type args struct {
		ctx       context.Context
		accountID int64
	}
	tests := []struct {
		name     string
		args     args
		expected []*model.AccountLimits
		wantErr  bool
	}{
		{
			name:     "get list limits for account_id 1",
			args:     args{ctx: context.Background(), accountID: 1},
			expected: accLimit1,
		},
		{
			name:     "get list limits for account_id 2",
			args:     args{ctx: context.Background(), accountID: 2},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := apiusecase.New(repo)
			res, err := u.GetAccountLimits(tt.args.ctx, tt.args.accountID)

			assert.Equal(t, tt.expected, res)

			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestUsecase_GetAccountTransactions(t *testing.T) {
	repo := new(api.Repository)
	repo.On("GetAccountTransactions", context.Background(), int64(1)).Return(trAcc1, nil)
	repo.On("GetAccountTransactions", context.Background(), int64(2)).Return(nil, nil)

	type args struct {
		ctx       context.Context
		accountID int64
	}
	tests := []struct {
		name     string
		args     args
		expected []*model.AccountTransaction
		wantErr  bool
	}{
		{
			name:     "get list transactions for account_id 1",
			args:     args{ctx: context.Background(), accountID: 1},
			expected: accTr1,
		},
		{
			name:     "get list transactions for account_id 2",
			args:     args{ctx: context.Background(), accountID: 2},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := apiusecase.New(repo)
			res, err := u.GetAccountTransactions(tt.args.ctx, tt.args.accountID)

			assert.Equal(t, tt.expected, res)

			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
