package logging

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

type loggingSuite struct {
	suite.Suite
	result *logStrings
}

func (s *loggingSuite) SetupTest() {
	s.result = &logStrings{}
	zl := zerolog.New(s.result)
	SetLogger(zl)
}

func TestLoggingSuite(t *testing.T) {
	suite.Run(t, new(loggingSuite))
}

func (s *loggingSuite) TestSetLogger() {
	zl := zerolog.New(s.result)
	zl = zl.Level(zerolog.ErrorLevel)
	s.Equal(zerolog.ErrorLevel, zl.GetLevel())

	SetLogger(zl)
	s.Equal(zerolog.ErrorLevel, logger.GetLevel())
}

func (s *loggingSuite) TestApplyDefaultLoggerContext() {
	zl := zerolog.New(s.result)
	l := applyDefaultLoggerContext(zl)

	l.Info().Str("key", "value").Send()
	s.Equal("value", s.result.data["key"])

	_, ok := s.result.data["time"]
	s.True(ok)
}

func (s *loggingSuite) TestTrace() {
	Trace().Send()
	s.Equal(zerolog.LevelTraceValue, s.result.data["level"])
}

func (s *loggingSuite) TestDebug() {
	Debug().Send()
	s.Equal(zerolog.LevelDebugValue, s.result.data["level"])
}

func (s *loggingSuite) TestInfo() {
	Info().Send()
	s.Equal(zerolog.LevelInfoValue, s.result.data["level"])
}

func (s *loggingSuite) TestWarn() {
	Warn().Send()
	s.Equal(zerolog.LevelWarnValue, s.result.data["level"])
}

func (s *loggingSuite) TestError() {
	Error().Send()
	s.Equal(zerolog.LevelErrorValue, s.result.data["level"])
}

func (s *loggingSuite) TestPanic() {
	defer func() {
		r := recover()
		s.NotNil(r)
		s.Equal(zerolog.LevelPanicValue, s.result.data["level"])
	}()

	Panic().Send()
}
