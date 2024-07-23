package logging

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type logQueue struct {
	wg     sync.WaitGroup
	m      sync.Mutex
	events []*zerolog.Event
}

var q = logQueue{
	events: []*zerolog.Event{},
}

func WaitAsyncLogs(ctx context.Context, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	waitForErrorsCh := make(chan int)
	go func() {
		q.wg.Wait()
		close(waitForErrorsCh)
	}()

	select {
	case <-ctx.Done():
		return
	case <-waitForErrorsCh:
		return
	}
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

	q.add(e)
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()

		qe := q.fetch()
		if qe != nil {
			qe.Send()
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
		Error().Err(err),
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
