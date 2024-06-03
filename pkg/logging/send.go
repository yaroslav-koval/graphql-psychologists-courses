package logging

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/rs/zerolog"
)

type logQueue struct {
	wg     sync.WaitGroup
	m      sync.Mutex
	events []*zerolog.Event
}

var q = logQueue{
	m:      sync.Mutex{},
	events: []*zerolog.Event{},
}

func WaitAsyncLogs() {
	q.wg.Wait()
}

func (q *logQueue) add(e *zerolog.Event) {
	q.m.Lock()
	q.events = append(q.events, e)
	q.m.Unlock()
}

func (q *logQueue) fetch() *zerolog.Event {
	q.m.Lock()

	var res *zerolog.Event

	if len(q.events) != 0 {
		res = q.events[0]
		q.events = q.events[1:len(q.events)]
	}

	q.m.Unlock()

	return res
}

func SendAsync(e *zerolog.Event, skipOptional ...int) {
	e = e.Array("stack", getStack(skipOptional...))

	logger.GetLevel()

	q.add(e)
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()

		le := q.fetch()
		if le != nil {
			le.Send()
		}
	}()
}

func Send(e *zerolog.Event, skipOptional ...int) {
	e = e.Array("stack", getStack(skipOptional...))

	e.Send()
}

func SendSimpleError(err error, skipOptional ...int) {
	skip := 1
	if len(skipOptional) != 0 {
		skip = skip + skipOptional[0]
	}

	Send(
		Error().Str("message", err.Error()),
		skip,
	)
}

func SendSimpleErrorAsync(err error, skipOptional ...int) {
	skip := 1
	if len(skipOptional) != 0 {
		skip = skip + skipOptional[0]
	}

	SendAsync(
		Error().Err(err),
		skip,
	)
}

func getStack(skip ...int) *StackTrace {
	i := 2

	if len(skip) != 0 {
		i = i + skip[0]
	}

	stack := StackTrace{}

	for true {
		_, file, line, ok := runtime.Caller(i)
		i++

		if !ok {
			break
		}

		stack = append(stack, fmt.Sprintf("%s:%v", file, line))
	}

	return &stack
}
