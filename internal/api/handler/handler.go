package handler

import (
	"kp/internal/api"
	"kp/internal/model"
	"kp/pkg/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	uc api.Usecase
}

func New(uc api.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) GetAccountProfile(c echo.Context) error {
	req := new(model.RequestAccountID)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	res, err := h.uc.GetAccountProfile(c.Request().Context(), req.AccountID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    res,
	})
}

func (h *Handler) GetAccountLimits(c echo.Context) error {
	req := new(model.RequestAccountID)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	res, err := h.uc.GetAccountLimits(c.Request().Context(), req.AccountID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    res,
	})
}

func (h *Handler) GetAccountTransactions(c echo.Context) error {
	req := new(model.RequestAccountID)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	res, err := h.uc.GetAccountTransactions(c.Request().Context(), req.AccountID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    res,
	})
}

func (h *Handler) CreateAccountTransaction(c echo.Context) error {
	req := c.Get(constant.ContextKeyRequestCreateTransaction).(*model.RequestCreateTransaction)

	res, err := h.uc.CreateTransaction(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    res,
	})
}

func (h *Handler) CreateAccount(c echo.Context) error {
	req := new(model.RequestCreateAccount)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	if req.KtpPhoto == nil {
		req.KtpPhoto = &constant.PlaceHolderPhoto
	}

	if req.SelfiePhoto == nil {
		req.SelfiePhoto = &constant.PlaceHolderPhoto
	}

	res, err := h.uc.CreateAccount(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    res,
	})
}
