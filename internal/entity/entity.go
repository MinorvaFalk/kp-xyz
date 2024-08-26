package entity

import "time"

type Customers struct {
	ID          int64      `json:"id"`
	Nik         string     `json:"nik"`
	FullName    string     `json:"full_name"`
	LegalName   string     `json:"legal_name"`
	BirthPlace  string     `json:"birth_place"`
	BirthDate   string     `json:"birth_date"`
	Salary      float64    `json:"salary"`
	KtpPhoto    string     `json:"ktp_photo"`
	SelfiePhoto string     `json:"selfie_photo"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type Accounts struct {
	ID         int64      `json:"id"`
	CustomerID int64      `json:"customer_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type Limits struct {
	ID            int64      `json:"id"`
	AccountID     int64      `json:"account_id"`
	Duration      int        `json:"duration"`
	InitialAmount float64    `json:"initial_amount"`
	CurrentAmount float64    `json:"current_amount"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

type Transactions struct {
	ID             int64      `json:"id"`
	AccountID      int64      `json:"account_id"`
	ContractNumber string     `json:"contract_number"`
	AssetName      string     `json:"asset_name"`
	Otr            float64    `json:"otr"`
	TotalPayment   float64    `json:"total_payment"`
	AdminFee       float64    `json:"admin_fee"`
	Installment    float64    `json:"installment"`
	Interest       float64    `json:"interest"`
	Duration       int        `json:"duration"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

type Fees struct {
	ID          int64      `json:"id"`
	Type        int64      `json:"type"`
	Description string     `json:"description"`
	Amount      float64    `json:"amount"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
