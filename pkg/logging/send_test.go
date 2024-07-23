package logging

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

type sendSuite struct {
	suite.Suite
	result *logResult
}

func (s *sendSuite) SetupSuite() {
	res := new(logResult)
	zl := zerolog.New(res)
	SetLogger(zl)
	s.result = res
}

func (s *sendSuite) SetupTest() {
	s.result.data = make(map[string]any)
}

func TestSendSuite(t *testing.T) {
	suite.Run(t, new(sendSuite))
}

func (s *sendSuite) TestAdd() {
	e := Info()
	s.Equal(0, len(q.events))
	q.add(e)
	s.Equal(1, len(q.events))

	q.events = q.events[0 : len(q.events)-1]
}

func (s *sendSuite) TestFetch() {
	s.Equal(0, len(q.events))

	e := Info().Str("key1", "value1")
	q.events = []*zerolog.Event{e}
	fe := q.fetch()
	fe.Str("key2", "value2").Send()

	s.Equal(0, len(q.events))
	s.Equal("value1", s.result.data["key1"])
	s.Equal("value2", s.result.data["key2"])
}

func (s *sendSuite) TestSendAsync() {
	e := Info().Str("key", "value")
	SendAsync(e)
	q.wg.Wait()

	s.Equal("value", s.result.data["key"])
	_, ok := s.result.data["stack"]
	s.True(ok)
}

func (s *sendSuite) TestSend() {
	e := Info().Str("key", "value")
	Send(e)

	s.Equal("value", s.result.data["key"])
	_, ok := s.result.data["stack"]
	s.True(ok)
}

func (s *sendSuite) TestSendSimpleError() {
	errText := "error_text"
	err := fmt.Errorf(errText)
	SendSimpleError(err)

	s.Equal(zerolog.LevelErrorValue, s.result.data["level"])
	s.Equal(errText, s.result.data[zerolog.ErrorFieldName])
	_, ok := s.result.data["stack"]
	s.True(ok)
}

func (s *sendSuite) TestSendSimpleErrorAsync() {
	errText := "error_text"
	err := fmt.Errorf(errText)
	SendSimpleErrorAsync(err)
	q.wg.Wait()

	s.Equal(zerolog.LevelErrorValue, s.result.data["level"])
	s.Equal(errText, s.result.data[zerolog.ErrorFieldName])
	_, ok := s.result.data["stack"]
	s.True(ok)
}

func (s *sendSuite) TestGetStack() {
	func() {
		func() {
			stack := getStack(0)
			s.Equal(2, countStackEntries(*stack, "send_test"))
		}()
	}()

	func() {
		func() {
			stack := getStack(1)
			s.Equal(1, countStackEntries(*stack, "send_test"))
		}()
	}()
}

func countStackEntries(arr []string, value string) int {
	res := 0

	for _, v := range arr {
		if strings.Contains(v, value) {
			res++
		}
	}

	return res
}
