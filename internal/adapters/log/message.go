package log

import (
	"context"

	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/AliceDiNunno/yeencloud/internal/core/interactor"
)

type Message struct {
	Level domain.LogLevel

	Message  string
	File     string
	Line     int
	Function string
	Module   string
	Time     int64

	Fields domain.LogFields
}

func (m *Message) WithLevel(level domain.LogLevel) interactor.LogMessage {
	m.Level = level

	return m
}

func (m *Message) At(time int64) interactor.LogMessage {
	m.Time = time
	return m
}

func (m *Message) WithContext(ctx context.Context) interactor.LogMessage {
	return m
}

func (m *Message) WithField(key domain.LogField, value interface{}) interactor.LogMessage {
	if m.Fields == nil {
		m.Fields = make(domain.LogFields)
	}

	m.Fields[key] = value

	return m
}

func (m *Message) WithFields(fields domain.LogFields) interactor.LogMessage {
	if m.Fields == nil {
		m.Fields = make(domain.LogFields)
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
