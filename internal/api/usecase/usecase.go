package apiusecase

import (
	"context"
	"kp/internal/api"
	"kp/internal/entity"
	"kp/internal/model"
	"kp/pkg/exception"
	"kp/pkg/logger"
	"slices"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type Usecase struct {
	repo api.Repository
}

func New(repo api.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) GetAccountProfile(ctx context.Context, accountID int64) (*model.AccountProfile, error) {
	cust, err := u.repo.GetAccountCustomer(ctx, accountID)
	if err != nil {
		return nil, err
	}

	data := model.AccountProfile{
		ID:          cust.ID,
		Nik:         cust.Nik,
		FullName:    cust.FullName,
		LegalName:   cust.LegalName,
		BirthPlace:  cust.BirthPlace,
		BirthDate:   cust.BirthDate,
		Salary:      cust.Salary,
		KtpPhoto:    cust.KtpPhoto,
		SelfiePhoto: cust.SelfiePhoto,
	}

	return &data, nil
}

func (u *Usecase) GetAccountLimits(ctx context.Context, accountID int64) ([]*model.AccountLimits, error) {
	data, err := u.repo.GetAccountLimits(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *Usecase) GetAccountTransactions(ctx context.Context, accountID int64) ([]*model.AccountTransaction, error) {
	var data []*model.AccountTransaction

	tr, err := u.repo.GetAccountTransactions(ctx, accountID)
	if err != nil {
		return nil, err
	}

	for _, t := range tr {
		data = append(data, &model.AccountTransaction{
			ContractNumber: t.ContractNumber,
			AssetName:      t.AssetName,
			Otr:            t.Otr,
			TotalPayment:   t.TotalPayment,
			Interest:       t.Interest,
			Installment:    t.Installment,
			AdminFee:       t.AdminFee,
			Duration:       t.Duration,
		})
	}

	return data, nil
}

func (u *Usecase) ValidateAccountLimit(ctx context.Context, req *model.RequestCreateTransaction) error {
	var (
		limit *model.AccountLimits
		admin float64
	)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		limits, err := u.repo.GetAccountLimits(ctx, req.AccountID)
		if err != nil {
			return err
		}

		idx := slices.IndexFunc(limits, func(al *model.AccountLimits) bool { return al.ID == req.LimitID })
		if idx == -1 {
			return exception.NewValidatonError("invalid limit_id", nil)
		}

		limit = limits[idx]

		return nil
	})

	g.Go(func() error {
		res, err := u.repo.GetAdminFee(ctx)
		if err != nil {
			return err
		}

		admin = res

		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	adminFee := req.Otr * admin
	payment := req.Otr + adminFee

	if payment > limit.CurrentAmount {
		return exception.NewValidatonError("insufficient limit", nil)
	}

	req.Admin = adminFee
	req.Interest = payment * limit.Interest / 100 * float64(limit.Duration)
	req.Installment = (payment + req.Interest) / float64(limit.Duration)

	req.TotalPayment = payment + req.Interest
	req.Limit = &entity.Limits{
		ID:            limit.ID,
		AccountID:     req.AccountID,
		Duration:      limit.Duration,
		CurrentAmount: limit.CurrentAmount - req.TotalPayment,
		InitialAmount: limit.InitialAmount,
	}

	return nil
}

func (u *Usecase) CreateTransaction(ctx context.Context, req *model.RequestCreateTransaction) (*model.AccountTransaction, error) {
	tr := entity.Transactions{
		AccountID:      req.AccountID,
		ContractNumber: uuid.NewString(),
		AssetName:      req.AssetName,
		Otr:            req.Otr,
		TotalPayment:   req.TotalPayment,
		AdminFee:       req.Admin,
		Installment:    req.Installment,
		Interest:       req.Interest,
		Duration:       req.Limit.Duration,
	}

	logger.Zap().Sugar().Info(tr)

	if err := u.repo.Trx(ctx, func(repo api.Repository) error {
		if err := repo.CreateTransactions(ctx, &tr); err != nil {
			return err
		}

		return repo.UpdateLimit(ctx, req.Limit)
	}); err != nil {
		return nil, err
	}

	data := model.AccountTransaction{
		ContractNumber: tr.ContractNumber,
		AssetName:      tr.AssetName,
		Otr:            tr.Otr,
		TotalPayment:   tr.TotalPayment,
		Interest:       tr.Interest,
		Installment:    tr.Installment,
		AdminFee:       tr.AdminFee,
		Duration:       tr.Duration,
		CurrentLimit:   &req.Limit.CurrentAmount,
	}

	return &data, nil
}
