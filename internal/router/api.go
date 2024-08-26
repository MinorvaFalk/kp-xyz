package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) MapHandler() {
	v1 := r.e.Group("/api/v1")

	v1.GET("/status", func(c echo.Context) error { return c.String(http.StatusOK, "OK") })

	auth := v1.Group("/auth")
	auth.POST("/register", r.h.CreateAccount)

	account := v1.Group("/account")
	account.GET("/profile", r.h.GetAccountProfile)
	account.GET("/limits", r.h.GetAccountLimits)
	account.GET("/transaction", r.h.GetAccountTransactions)
	account.POST("/transaction", r.h.CreateAccountTransaction, r.h.ValidateCreateTransactionRequest)
}
