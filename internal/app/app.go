package app

import (
	"fmt"

	"github.com/Masedko/go-backend/config"
	pkgerrors "github.com/Masedko/go-backend/internal/core/errors"
	"github.com/Masedko/go-backend/internal/core/handler"
	"github.com/Masedko/go-backend/internal/core/logger"
	"github.com/Masedko/go-backend/internal/core/router"
	"github.com/Masedko/go-backend/internal/core/storage"
	"github.com/Masedko/go-backend/internal/data/database"
)

type App struct {
	Logger *logger.Logger
	Config *config.Config
	Router *router.Router
}

func New() *App {
	log := logger.NewLogger()
	cfg, err := config.LoadConfig(config.Load{
		Path: "./config",
		Name: "config",
		Type: "yaml",
	})
	if err != nil {
		log.FatalWithDesc(err)
	}

	r := router.New(log)
	v1 := r.Group("/api/v1")

	db, err := database.NewDB(database.Config{
		Host:     cfg.Database.Host,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	})
	if err != nil {
		log.FatalWithDesc(err)
	}
	log.Info().Msg("Database connected")

	s, err := storage.New(cfg.Storage.BucketNames)
	if err != nil {
		log.FatalWithDesc(err)
	}
	log.Info().Msg("Storage connected")

	h := handler.New(db, s)
	h.Register(v1)

	return &App{
		Logger: log,
		Config: cfg,
		Router: r,
	}
}

func (app *App) Run() {
	err := app.Router.Start(fmt.Sprintf("%s:%d", app.Config.Server.Host, app.Config.Server.Port))
	if err != nil {
		app.Logger.ErrorWithDesc(pkgerrors.Error{
			Err:  err,
			Desc: "Cannot start web app",
		})
	}
}
