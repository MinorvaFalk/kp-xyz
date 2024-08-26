package api

import "github.com/labstack/echo/v4"

type Handler interface {
	GetAccountProfile(c echo.Context) error
	GetAccountLimits(c echo.Context) error
	GetAccountTransactions(c echo.Context) error

	CreateAccountTransaction(c echo.Context) error

	ValidateCreateTransactionRequest(next echo.HandlerFunc) echo.HandlerFunc
}
