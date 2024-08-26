package constant

import "kp/internal/entity"

type ContextKey string

var (
	ContextKeyRequestCreateTransaction = "request_create_transaction"

	DefaultLimits = func(accountID int64) []*entity.Limits {
		return []*entity.Limits{
			{
				AccountID:     accountID,
				Duration:      1,
				InitialAmount: 1000000,
				CurrentAmount: 1000000,
			},
			{
				AccountID:     accountID,
				Duration:      2,
				InitialAmount: 2000000,
				CurrentAmount: 2000000,
			},
			{
				AccountID:     accountID,
				Duration:      3,
				InitialAmount: 3000000,
				CurrentAmount: 3000000,
			},
			{
				AccountID:     accountID,
				Duration:      4,
				InitialAmount: 4000000,
				CurrentAmount: 4000000,
			},
		}
	}

	PlaceHolderPhoto = "http://dummyimage.com/250x100.png/cc0000/ffffff"
)
