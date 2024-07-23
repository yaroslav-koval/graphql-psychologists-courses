package logging

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger = zerolog.New(nil)

func SetLogger(l zerolog.Logger) {
	logger = applyDefaultLoggerContext(l)
}

func SetDefaultLogger() {
	logger = applyDefaultLoggerContext(zerolog.New(os.Stdout))
}

func applyDefaultLoggerContext(l zerolog.Logger) zerolog.Logger {
	return l.With().Timestamp().Logger()
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

func Panic() *zerolog.Event {
	return logger.Panic()
}
