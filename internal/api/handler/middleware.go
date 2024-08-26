package handler

import (
	"kp/internal/model"
	"kp/pkg/constant"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ValidateCreateTransactionRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(model.RequestCreateTransaction)
		if err := c.Bind(req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}

		if err := h.uc.ValidateAccountLimit(c.Request().Context(), req); err != nil {
			return err
		}

		c.Set(constant.ContextKeyRequestCreateTransaction, req)

		return next(c)
	}
}
