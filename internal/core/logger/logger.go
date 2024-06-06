package logger

import (
	"errors"
	"os"

	"github.com/rs/zerolog"

	pkgerrors "github.com/Masedko/go-backend/internal/core/errors"
)

type Logger struct {
	zerolog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

func (l *Logger) ErrorWithDesc(err error) {
	var e pkgerrors.Error
	if errors.As(err, &e) {
		l.Error().Err(e.Err).Msg(e.Desc)
	} else {
		l.Error().Err(err).Msg("Unknown error")
	}
}

func (l *Logger) FatalWithDesc(err error) {
	var e pkgerrors.Error
	if errors.As(err, &e) {
		l.Fatal().Err(e.Err).Msg(e.Desc)
	} else {
		l.Fatal().Err(err).Msg("Unknown error")
	}
}
