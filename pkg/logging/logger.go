package logging

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger = getDefaultLoggerContext(zerolog.New(os.Stdout))

func Logger() zerolog.Logger {
	return logger
}

func SetLogger(l zerolog.Logger) {
	logger = getDefaultLoggerContext(l)
}

func getDefaultLoggerContext(l zerolog.Logger) zerolog.Logger {
	resL := l.With().Timestamp().Logger()
	return resL
}

func Trace() *zerolog.Event {
	return logger.Trace()
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

func Panic() *zerolog.Event {
	return logger.Panic()
}

func WithLevel(lvl zerolog.Level) *zerolog.Event {
	return logger.WithLevel(lvl)
}
