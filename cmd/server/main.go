package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-backend/internal/core/router"
)

func main() {
	e := router.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
