package router

import (
	"context"
	"kp/config"
	"kp/internal/api/handler"
	"kp/pkg/exception"
	"kp/pkg/logger"
	"kp/pkg/validation"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type Router struct {
	e *echo.Echo
	h *handler.Handler
}

func NewRouter(h *handler.Handler) *Router {
	e := echo.New()
	e.Validator = validation.NewValidator()
	e.HTTPErrorHandler = exception.EchoErrorHandler

	// TODO: re-configure request logger
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogRemoteIP: true,
		LogMethod:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Zap().Info("http_request",
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.String("method", v.Method),
				zap.String("remote_ip", v.RemoteIP),
			)

			return nil
		},
	}))
	e.Use(middleware.Recover())
	// e.Use(middleware.CSRF())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// TODO: configure origin based on env
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
	}))

	return &Router{
		e: e,
		h: h,
	}
}

func (r *Router) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := r.e.Start(":" + config.ReadConfig().Port); err != nil && err != http.ErrServerClosed {
			r.e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.e.Shutdown(ctx); err != nil {
		r.e.Logger.Fatal(err)
	}
}
