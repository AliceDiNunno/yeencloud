package log

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	"github.com/rogpeppe/go-internal/modfile"
	"io/ioutil"
)

type LoggerMiddleware interface {
	Log(message Message)
}

type Log struct {
	middleware []LoggerMiddleware
}

func GetModuleName() string {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "unknown"
	}

	modName := modfile.ModulePath(goModBytes)

	return modName
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
