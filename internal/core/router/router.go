package router

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/Masedko/go-backend/internal/core/logger"
	"github.com/Masedko/go-backend/internal/core/middleware"
	"github.com/Masedko/go-backend/internal/core/middleware/validator"
)

type Router struct {
	*echo.Echo
}

func New(log *logger.Logger) *Router {
	e := echo.New()
	e.Pre(echomiddleware.RemoveTrailingSlash())
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Validator = validator.New()
	e.Use(middleware.NewLogger(log))
	e.Use(middleware.NewRecovery())
	return &Router{e}
}
