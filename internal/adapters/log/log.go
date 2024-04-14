package log

import (
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/AliceDiNunno/yeencloud/internal/core/interactor"
)

type LoggerMiddleware interface {
	Log(message Message)
}

type Log struct {
	middleware []LoggerMiddleware
}

func (l *Log) Log(level domain.LogLevel) interactor.LogMessage {
	return (&Message{}).WithLevel(level)
}

var logger = NewLog()

func (l *Log) AddMiddleware(middleware LoggerMiddleware) {
	if l.middleware == nil {
		l.middleware = make([]LoggerMiddleware, 0)
	}

	l.middleware = append(l.middleware, middleware)
}

func NewLog() *Log {
	return &Log{}
}

func Logger() *Log {
	return logger
}
