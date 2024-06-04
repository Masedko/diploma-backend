package app

import (
	"go-backend/internal/core/router"
	"go-backend/internal/data/database"
)

type App struct {
	Config *config.Config
	Router *router.Router
	DB     *database.DB
}
