package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewLoggerMiddleware(logger *logger.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")
			return nil
		},
	})
}
