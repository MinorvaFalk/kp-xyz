package model

import "kp/internal/entity"

type RequestAccountID struct {
	AccountID int64 `json:"account_id" validate:"required"`
}

type RequestCreateAccount struct {
	Nik         string  `json:"nik" validate:"nik"`
	FullName    string  `json:"full_name" validate:"required"`
	LegalName   string  `json:"legal_name" validate:"required"`
	BirthPlace  string  `json:"birth_place" validate:"required"`
	BirthDate   string  `json:"birth_date" validate:"date"`
	Salary      float64 `json:"salary" validate:"required"`
	KtpPhoto    *string `json:"ktp_photo"`
	SelfiePhoto *string `json:"selfie_photo"`
}

type RequestCreateTransaction struct {
	AccountID int64   `json:"account_id" validate:"required"`
	LimitID   int64   `json:"limit_id" validate:"required"`
	AssetName string  `json:"asset_name" validate:"required"`
	Otr       float64 `json:"otr" validate:"required"`

	Interest     float64
	Installment  float64
	Admin        float64
	TotalPayment float64

	Limit *entity.Limits
}

type AccountTransaction struct {
	ContractNumber string  `json:"contract_number"`
	AssetName      string  `json:"asset_name"`
	Otr            float64 `json:"otr"`
	AdminFee       float64 `json:"admin_fee"`
	TotalPayment   float64 `json:"total_payment"`
	Interest       float64 `json:"interest"`
	Installment    float64 `json:"installment"`
	Duration       int     `json:"duration"`

	CurrentLimit *float64 `json:"current_limit,omitempty"`
}

type AccountLimits struct {
	ID            int64   `json:"id"`
	Duration      int     `json:"duration"`
	InitialAmount float64 `json:"initial_amount"`
	CurrentAmount float64 `json:"current_amount"`
	Interest      float64 `json:"interest"`
}

type AccountProfile struct {
	ID          int64   `json:"id"`
	Nik         string  `json:"nik"`
	FullName    string  `json:"full_name"`
	LegalName   string  `json:"legal_name"`
	BirthPlace  string  `json:"birth_place"`
	BirthDate   string  `json:"birth_date"`
	Salary      float64 `json:"salary"`
	KtpPhoto    string  `json:"ktp_photo"`
	SelfiePhoto string  `json:"selfie_photo"`
}

type AccountCreated struct {
	AccountID   int64                   `json:"account_id"`
	CustomerID  int64                   `json:"customer_id"`
	Nik         string                  `json:"nik"`
	FullName    string                  `json:"full_name"`
	LegalName   string                  `json:"legal_name"`
	BirthPlace  string                  `json:"birth_place"`
	BirthDate   string                  `json:"birth_date"`
	Salary      float64                 `json:"salary"`
	KtpPhoto    string                  `json:"ktp_photo"`
	SelfiePhoto string                  `json:"selfie_photo"`
	Limits      []*AccountCreatedLimits `json:"limits"`
}

type AccountCreatedLimits struct {
	ID            int64   `json:"id"`
	Duration      int     `json:"duration"`
	InitialAmount float64 `json:"initial_amount"`
	CurrentAmount float64 `json:"current_amount"`
}
