package log

import (
	"context"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
)

type LogFields map[string]interface{}

type Message struct {
	Level domain.LogLevel

	Message  string
	File     string
	Line     int
	Function string
	Module   string
	Time     int64

	Fields LogFields
}

func (m *Message) WithLevel(Level domain.LogLevel) interactor.LogMessage {
	m.Level = Level

	return m
}

func (m *Message) At(time int64) interactor.LogMessage {
	m.Time = time
	return m
}

func (m *Message) WithContext(ctx context.Context) interactor.LogMessage {
	return m
}

func (m *Message) WithField(key string, value interface{}) interactor.LogMessage {
	if m.Fields == nil {
		m.Fields = make(LogFields)
	}

	m.Fields[key] = value

	return m
}

func (m *Message) WithFields(fields map[string]interface{}) interactor.LogMessage {
	if m.Fields == nil {
		m.Fields = make(LogFields)
	}

	for key, value := range fields {
		m.Fields[key] = value
	}

	return m
}

func (m *Message) Msg(message string) {
	m.Message = message

	for _, middleware := range logger.middleware {
		middleware.Log(*m)
	}
}
